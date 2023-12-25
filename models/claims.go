package models

import "github.com/golang-jwt/jwt"

type Claims struct {
	jwt.StandardClaims
	iss   string
	sub   string
	aud   string
	exp   int64
	iat   int64
	roles []string
}
