package auth

import (
	"errors"
	"time"

	"github.com/JGurus/template-initial-api/models"
	"github.com/dgrijalva/jwt-go"
)

//GenerateToken .
func GenerateToken(data *models.Login) (string, error) {
	claim := models.Claim{
		Email: data.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 2).Unix(),
			Issuer:    "jgurus",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claim)
	signedToken, err := token.SignedString(signKey)
	if err != nil {
		return "", err
	}
	return signedToken, nil
}

//ValidatedToken .
func ValidatedToken(t string) (models.Claim, error) {
	token, err := jwt.ParseWithClaims(t, &models.Claim{}, verifyFunction)
	if err != nil {
		return models.Claim{}, err
	}
	if !token.Valid {
		return models.Claim{}, errors.New("token no válido")
	}
	claim, ok := token.Claims.(*models.Claim)
	if !ok {
		return models.Claim{}, errors.New("no se pudo obtener los claims")
	}
	return *claim, nil
}

func verifyFunction(t *jwt.Token) (interface{}, error) {
	return verifyKey, nil
}
