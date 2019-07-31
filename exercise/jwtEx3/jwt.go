package main

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtSecret = []byte("jwt_Secret")

type Claims struct {
	Username string `json:"username"`
	Birthday string `json:"birthday"`
	Email    string `json:"email"`
	jwt.StandardClaims
}

func GenerateToken(user User) (string, error) {
	exirationTime := time.Now().Add(5 * time.Minute)
	claims := Claims{
		user.Username,
		user.Birthday,
		user.Email,
		jwt.StandardClaims{
			ExpiresAt: exirationTime.Unix(),
		},
	}
	signingMethod := jwt.GetSigningMethod("HS256")
	token := jwt.NewWithClaims(signingMethod, claims)
	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func AuthenticateToken(tokenString string) error {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if err != nil {
		return err
	}
	if !token.Valid {
		return errors.New("tokenValid")
	}
	return nil
}

func RefreshToken(token string) {

}
