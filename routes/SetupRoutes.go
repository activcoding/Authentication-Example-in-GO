package routes

import (
	"auth_example/config"
	"github.com/gorilla/mux"
)

func SetupRoutes(mainRouter *mux.Router, config *config.DatabaseConfig) {
	SetupUserAuthRoutes(mainRouter, config)
	SetupExampleRestrictedRoutes(mainRouter)
}
