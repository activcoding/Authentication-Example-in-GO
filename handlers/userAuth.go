package handlers

import (
	"auth_example/config"
	"auth_example/models"
	"auth_example/utils"
	"encoding/json"
	"go.mongodb.org/mongo-driver/bson"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

type UserAuth struct {
	Config *config.DatabaseConfig
}

// SignIn ensures that the user is authenticated and returns a JWT token
func (userAuth *UserAuth) SignIn(w http.ResponseWriter, r *http.Request) {
	var signInModel models.SignInModel
	err := json.NewDecoder(r.Body).Decode(&signInModel)
	if err != nil {
		http.Error(w, "Failed to parse request body", http.StatusBadRequest)
		return
	}

	var user models.User

	// find the user by email
	filter := bson.M{"email": signInModel.Email}
	err = userAuth.Config.UserCollection.FindOne(r.Context(), filter).Decode(&user)
	if err != nil {
		http.Error(w, "Failed to find user by email", http.StatusNotFound)
		return
	}

	// check if the password is correct
	if user.Password != signInModel.Password {
		http.Error(w, "Incorrect password", http.StatusUnauthorized)
		return
	}

	// Return a JWT
	// Generate a JWT token
	token, err := utils.GenerateJWT(user.Email, user.Roles)
	if err != nil {
		http.Error(w, "Failed to generate JWT", http.StatusInternalServerError)
		return
	}

	// Return the token as JSON
	jsonObj := map[string]string{
		"token": token,
	}

	// Marshal the JSON object into a byte array
	jsonData, err := json.Marshal(jsonObj)
	if err != nil {
		http.Error(w, "Failed to marshal JSON", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(jsonData)
	if err != nil {
		http.Error(w, "Failed to write response", http.StatusInternalServerError)
		return
	}
}

func (userAuth *UserAuth) SignUp(w http.ResponseWriter, r *http.Request) {
	// decode the request body into struct and failed if any error occurs
	var user models.User
	// TODO: Don't save the password in plaintext
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Failed to parse request body", http.StatusBadRequest)
		return
	}

	// insert the user into the database
	_, err = userAuth.Config.UserCollection.InsertOne(r.Context(), user)
	if err != nil {
		http.Error(w, "Failed to insert user into database", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (userAuth *UserAuth) SendEmailVerification(w http.ResponseWriter, r *http.Request) {
	var user models.SignInModel
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Failed to parse request body", http.StatusBadRequest)
		return
	}

	verificationCode := generateVerificationCode()
	userAuth.Config.UserCollection.FindOneAndUpdate(r.Context(), bson.M{"email": user.Email, "password": user.Password},
		bson.M{"$set": bson.M{"activationCode": verificationCode}})
	//utils.SendEmail(user.Email, "Activation code for your account", "Your activation code is: "+verificationCode)
}

func generateVerificationCode() string {
	src := rand.NewSource(time.Now().UnixNano())
	r := rand.New(src)

	return strconv.Itoa(r.Intn(9000) + 1000)
}

func (userAuth *UserAuth) UpdateAccount(w http.ResponseWriter, r *http.Request) {

}

func (userAuth *UserAuth) DeleteAccount(w http.ResponseWriter, r *http.Request) {

}
