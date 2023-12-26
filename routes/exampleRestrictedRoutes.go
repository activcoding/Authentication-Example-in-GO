package routes

import (
	"auth_example/middleware"
	"github.com/gorilla/mux"
)

func SetupExampleRestrictedRoutes(mainRouter *mux.Router) *mux.Router {
	subRouter := mainRouter.PathPrefix("/exampleRestricted").Subrouter()

	subRouter.Use(middleware.Logger)
	subRouter.Use(middleware.JWTUserAuthMiddleware)

	subRouter.HandleFunc("/exampleRestricted", nil).Methods("GET")

	return subRouter
}
