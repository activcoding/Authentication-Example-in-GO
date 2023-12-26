package middleware

import (
	"fmt"
	"net/http"
)
import "auth_example/utils"

func APIKeyValidation(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// The api key that the user sends in the header
		apiKey := r.Header.Get("X-API-KEY")

		// The valid api key that is stored in the environment variables
		validAPIKey := utils.GetVariable("API_KEY")

		if apiKey != validAPIKey {
			fmt.Println("Unauthorized")
			fmt.Println("API Key: " + apiKey)
			fmt.Println("Valid API Key: " + validAPIKey)

			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}
