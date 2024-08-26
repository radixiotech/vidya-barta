package v1

import (
	v1_users_handler "github.com/radixiotech/vidya-barta/apps/vb-api/handlers/v1/users"
	"github.com/radixiotech/vidya-barta/business/config"
	"github.com/radixiotech/vidya-barta/foundation/web"
	"go.uber.org/zap"
)

const version = "/api/v1"
const userRoute = version + "/users"

type Config struct {
	Log    *zap.SugaredLogger
	Config *config.VBApiConfig
}

func SetupCRoutes(app *web.App, cfg Config) {
	usersHandler := v1_users_handler.NewHandler()

	// Users Handlers
	app.Mux.Get(userRoute, usersHandler.Query)
}
