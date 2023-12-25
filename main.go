package main

import (
	"auth_example/config"
	"auth_example/routes"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"net/http"
)

func main() {
	port := 8081
	address := fmt.Sprintf(":%d", port)

	databaseConfig := setupDatabase()
	userAuthRouter := routes.SetupUserAuthRoutes(databaseConfig)

	http.Handle("/", userAuthRouter)

	fmt.Printf("Server is listening on port %d", port)
	err := http.ListenAndServe(address, nil)
	if err != nil {
		fmt.Println(err)
	}
}

func setupDatabase() *config.DatabaseConfig {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		panic(err)
	}

	database := client.Database("auth_example")
	userCollection := database.Collection("users")

	databaseConfig := &config.DatabaseConfig{
		Database:       database,
		UserCollection: userCollection,
	}

	return databaseConfig
}
