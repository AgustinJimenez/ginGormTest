package types

import "github.com/dgrijalva/jwt-go"

type RegisterUserResponseType struct {
	Email string `json:"email"`
	Id uint `json:"id"`
	Username string `json:"username"`
}

type Claims struct {
    UserID uint `json:"user_id"`
    jwt.StandardClaims
}

