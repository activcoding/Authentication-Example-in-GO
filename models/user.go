package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID               primitive.ObjectID `bson:"_id,omitempty"`
	Username         string             `bson:"username"`
	Password         string             `bson:"password"`
	Email            string             `bson:"email"`
	AccountActivated bool               `bson:"accountActivated"`
	ActivationCode   string             `bson:"activationCode"`
	Roles            []string           `bson:"roles"`
}
