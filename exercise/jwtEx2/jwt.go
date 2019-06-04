package main

import (
	"log"

	"github.com/dgrijalva/jwt-go"
)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	jwt.StandardClaims
}

func CreateTokenString(u *User) string {
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), u)
	tokenstring, err := token.SignedString([]byte("foobar"))
	if err != nil {
		log.Fatalln(err)
	}
	return tokenstring
}
