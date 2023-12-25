package config

import "go.mongodb.org/mongo-driver/mongo"

type DatabaseConfig struct {
	Database       *mongo.Database
	UserCollection *mongo.Collection
}
