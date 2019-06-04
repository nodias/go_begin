package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtKey = []byte("my_secret_key")

// DB
var users = map[string]string{
	"kss": "kss123",
	"jeh": "jeh123",
}

type Credentials struct {
	Password string `json: "password"`
	Username string `json: "username"`
}

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func Signin(w http.ResponseWriter, r *http.Request) {
	expirationTime := time.Now().Add(5 * time.Minute)
	user := User{
		Name: "kss123",
		Age:  29,
	}
	tokenString := CreateTokenString(&user)

	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   tokenString,
		Expires: expirationTime,
	})
}

func Welcome(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("token")
	if err != nil {
		if err == http.ErrNoCookie {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	tknStr := cookie.Value

	claims := &Claims{}
	tkn, err := jwt.ParseWithClaims(tknStr, claims, func(jwt *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	fmt.Println(tkn)

}

func Refresh(w http.ResponseWriter, r *http.Request) {
}
