package routes

import (
	"auth_example/config"
	"auth_example/handlers"
	"auth_example/middleware"
	"github.com/gorilla/mux"
)

func SetupUserAuthRoutes(config *config.DatabaseConfig) *mux.Router {
	router := mux.NewRouter()
	router.Use(middleware.Logger)
	//router.Use(middleware.JWTUserAuthMiddleware)

	userAuthHandler := &handlers.UserAuth{Config: config}

	router.HandleFunc("/signin", userAuthHandler.SignIn).Methods("POST")
	router.HandleFunc("/signup", userAuthHandler.SignUp).Methods("POST")
	router.HandleFunc("/deleteAccount", userAuthHandler.DeleteAccount).Methods("DELETE")
	return router
}
