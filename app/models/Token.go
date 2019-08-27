package models

import "github.com/dgrijalva/jwt-go"

type Token struct {
	UserId   int `json:"user_id"`
	jwt.StandardClaims
}

