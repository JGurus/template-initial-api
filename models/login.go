package models

import "github.com/dgrijalva/jwt-go"

//Login estructura
type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

//Claim struct
type Claim struct {
	Email string `json:"email"`
	jwt.StandardClaims
}
