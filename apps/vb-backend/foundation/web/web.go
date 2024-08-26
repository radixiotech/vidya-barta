// Package web contains a small web framework extension.
package web

import (
	"net/http"
	"os"
	"syscall"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

// App is the entrypoint into our application and what configures our context
// object for each of our http handlers.
type App struct {
	Mux      *chi.Mux
	shutdown chan os.Signal
}

type AppConfig struct {
	Cors     *cors.Options
	Shutdown chan os.Signal
}

var defaultCors cors.Options = cors.Options{
	MaxAge:           300,
	AllowCredentials: true,
	AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
	AllowedMethods: []string{
		http.MethodGet,
		http.MethodPut,
		http.MethodPost,
		http.MethodHead,
		http.MethodPatch,
		http.MethodDelete,
		http.MethodOptions,
	},
}

// NewApp creates an App value that handle a set of routes for the application.
func NewApp(cfg AppConfig) *App {
	mux := chi.NewRouter()

	if cfg.Cors == nil {
		cfg.Cors = &defaultCors
	}

	mux.Use(cors.Handler(*cfg.Cors))
	mux.Use(middleware.Logger)
	mux.Use(middleware.RealIP)
	mux.Use(middleware.RequestID)
	mux.Use(middleware.Recoverer)

	return &App{
		Mux:      mux,
		shutdown: cfg.Shutdown,
	}
}

// SignalShutdown is used to gracefully shut down the app when an integrity issue is identified.
func (a *App) SignalShutdown() {
	a.shutdown <- syscall.SIGTERM
}
