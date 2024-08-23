package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	if err := run(); err != nil {
		fmt.Printf("Error :%v", err)
		os.Exit(1)
	}
}

func run() error {
	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, syscall.SIGINT, syscall.SIGTERM)

	api := http.Server{}
	serverErrors := make(chan error, 1)

	go func() {
		fmt.Printf("Server listening on %s", api.Addr)
		serverErrors <- api.ListenAndServe()
	}()

	// Shutdown Server
	select {
	case err := <-serverErrors:
		return fmt.Errorf("server error: %w", err)

	case sig := <-shutdown:
		fmt.Printf("Shutting down server with signal : %+v", sig)
		defer fmt.Printf("Shutdown complete with signal : %+v", sig)

		ctx, cancel := context.WithTimeout(context.Background(), time.Second*20)
		defer cancel()

		if err := api.Shutdown(ctx); err != nil {
			api.Close()
			return fmt.Errorf("could not stop server gracefully: %w", err)
		}
	}

	return nil
}
