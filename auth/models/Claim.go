package models

import "github.com/dgrijalva/jwt-go"

type Claim struct {
	jwt.StandardClaims
	Username string `json:"username"`
}
