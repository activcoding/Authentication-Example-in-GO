package routes

import (
	"auth_example/config"
	"auth_example/handlers"
	"auth_example/middleware"
	"github.com/gorilla/mux"
)

func SetupUserAuthRoutes(mainRouter *mux.Router, config *config.DatabaseConfig) {
	subRouter := mainRouter.PathPrefix("/auth").Subrouter()
	subRouter.Use(middleware.Logger)

	// Use JWT middleware for routes that require a signed-in user
	// Use the API key for routes that are for returning a JWT to the user
	// subRouter.Use(middleware.JWTUserAuthMiddleware)
	subRouter.Use(middleware.APIKeyValidation)

	userAuthHandler := &handlers.UserAuth{Config: config}

	subRouter.HandleFunc("/signin", userAuthHandler.SignIn).Methods("POST")
	subRouter.HandleFunc("/signup", userAuthHandler.SignUp).Methods("POST")
	subRouter.HandleFunc("/sendActivationEmail", userAuthHandler.SendActivationEmail).Methods("Post")
	subRouter.HandleFunc("/activateAccount", userAuthHandler.ActivateAccount).Methods("Post")
}
