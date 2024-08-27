package handlers

import (
	"net/http"
	"os"

	"github.com/go-chi/cors"
	"github.com/jmoiron/sqlx"
	v1 "github.com/radixiotech/vidya-barta/apps/vb-api/handlers/v1"
	"github.com/radixiotech/vidya-barta/business/config"
	"github.com/radixiotech/vidya-barta/foundation/web"
	"go.uber.org/zap"
)

type APIHandlersConfig struct {
	DB       *sqlx.DB
	Shutdown chan os.Signal
	Log      *zap.SugaredLogger
	Config   *config.VBApiConfig
}

func APIMux(cfg APIHandlersConfig) http.Handler {
	app := web.NewApp(web.AppConfig{
		Shutdown: cfg.Shutdown,
		Cors: &cors.Options{
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
		},
	})

	// Register API V1
	v1.SetupCRoutes(app, v1.Config{Log: cfg.Log, Config: cfg.Config})

	return app.Mux
}
