package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"runtime"
	"syscall"

	"github.com/joho/godotenv"
	"github.com/radixiotech/vidya-barta/apps/vb-api/handlers"
	"github.com/radixiotech/vidya-barta/business/config"
	"github.com/radixiotech/vidya-barta/business/database"
	"github.com/radixiotech/vidya-barta/foundation/logger"
	"go.uber.org/zap"
)

func main() {
	// Load envs.
	godotenv.Load()

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
	cfg := config.NewVBConfig()
	log.Infow("Config", "config", cfg)

	// Initialize Database
	db, err := database.Open(cfg.DB)
	if err == nil {
		log.Infow("DATABASE ERROR", "error", err)
		return err
	}

	err = database.StatusCheck(context.Background(), db)
	if err != nil {
		log.Infow("DATABASE Status Check Error", "error", err)
		return err
	}

	// Shutdown Signals
	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, syscall.SIGINT, syscall.SIGTERM)

	// Initialize API support
	mux := handlers.APIMux(handlers.APIHandlersConfig{Shutdown: shutdown, Log: log, Config: cfg, DB: db})

	api := http.Server{
		Handler:      mux,
		Addr:         cfg.Web.APIHost,
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

	// Graceful Shutdown of Server
	select {
	case err := <-serverErrors:
		return fmt.Errorf("server error: %w", err)

	case sig := <-shutdown:
		log.Infow("Shutting Down Server", "signal", sig)
		defer log.Infow("Shutdown Complete", "signal", sig)

		ctx, cancel := context.WithTimeout(context.Background(), cfg.Web.ShutdownTimeout)
		defer cancel()

		if err := api.Shutdown(ctx); err != nil {
			api.Close()
			return fmt.Errorf("could not stop server gracefully: %w", err)
		}
	}

	return nil
}
