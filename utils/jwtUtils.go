package utils

import (
	"fmt"
	"github.com/golang-jwt/jwt"
	"time"
)

func GenerateJWT(subject string, roles []string) (string, error) {
	signingKey := []byte(GetVariable("JWT_KEY"))
	iss := []byte(GetVariable("ISS"))
	aud := []byte(GetVariable("AUD"))

	claims := jwt.MapClaims{
		"iss":   string(iss),
		"sub":   subject,
		"aud":   string(aud),
		"exp":   time.Now().Add(time.Hour * 1).Unix(), // Token expires in 1 hour
		"iat":   time.Now().Unix(),
		"roles": roles,
	}

	// generate the token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the token with the secret key
	tokenString, err := token.SignedString(signingKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func VerifyJWT(tokenString string) (jwt.MapClaims, error) {
	signingKey := []byte(GetVariable("JWT_KEY"))

	// Parse the token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Check the signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return signingKey, nil
	})

	// Check for parsing errors
	if err != nil {
		return nil, err
	}

	// Verify the token
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, fmt.Errorf("invalid token")
}
