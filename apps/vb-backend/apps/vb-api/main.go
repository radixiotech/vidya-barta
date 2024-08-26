package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"runtime"
	"syscall"
	"time"

	"github.com/radixiotech/vidya-barta/apps/vb-api/handlers"
	"github.com/radixiotech/vidya-barta/business/config"
	"github.com/radixiotech/vidya-barta/foundation/logger"
	"go.uber.org/zap"
)

func main() {
	log := logger.New("Vidya Barta Backend")
	defer log.Sync()

	if err := run(log); err != nil {
		log.Error("Startup", "Error", err)
		log.Sync()
		os.Exit(1)
	}
}

func run(log *zap.SugaredLogger) error {
	// GOMAXPROCS
	log.Infow("Startup", "GOMAXPROCS", runtime.GOMAXPROCS(0))

	// Initialize Config
	cfg := config.New()

	// Graceful Shutdown
	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, syscall.SIGINT, syscall.SIGTERM)

	// Initialize API support
	mux := handlers.APIMux(handlers.APIHandlersConfig{Shutdown: shutdown, Log: log, Config: cfg})

	api := http.Server{
		Handler: mux,
		// Addr:         cfg.Web.APIHost,
		Addr:         "localhost:3000",
		ReadTimeout:  cfg.Web.ReadTimeout,
		IdleTimeout:  cfg.Web.IdleTimeout,
		WriteTimeout: cfg.Web.WriteTimeout,
		ErrorLog:     zap.NewStdLog(log.Desugar()),
	}

	serverErrors := make(chan error, 1)

	go func() {
		log.Infow("Server Listening", "address", api.Addr)
		serverErrors <- api.ListenAndServe()
	}()

	// Shutdown Server
	select {
	case err := <-serverErrors:
		return fmt.Errorf("server error: %w", err)

	case sig := <-shutdown:
		log.Infow("Shutting Down Server", "signal", sig)
		defer log.Infow("Shutdown Complete", "signal", sig)

		ctx, cancel := context.WithTimeout(context.Background(), time.Second*20)
		defer cancel()

		if err := api.Shutdown(ctx); err != nil {
			api.Close()
			return fmt.Errorf("could not stop server gracefully: %w", err)
		}
	}

	return nil
}
