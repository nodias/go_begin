package main

import (
	"log"

	"github.com/dgrijalva/jwt-go"
)

func main() {
	//for example, server receive token string in request header.
	tokenstring := createTokenString()
	//This is that token string.
	log.Println(tokenstring)
	//Let's parse this by the secrete, which only server knows.
	token, err := jwt.Parse(tokenstring, func(token *jwt.Token) (interface{}, error) {
		return []byte("foobar"), nil
	})
	//When using `Parse`, the result `Claims` would be a map.
	log.Println("token claims : ", token.Claims)
	log.Println("token valid : ", token.Valid)
	//in another way, you can decode token to your struct, which needs to satisfy `jwt.StandarsClaims`
	user := User{}
	token, err = jwt.ParseWithClaims(tokenstring, &user, func(token *jwt.Token) (interface{}, error) {
		return []byte("foobar"), nil
	})
	log.Println(token.Valid, user, err)
}

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	jwt.StandardClaims
}

func createTokenString() string {
	//Embed User information to `token`
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), &User{
		Name: "otial10",
		Age:  30,
	})
	//token --> string. Only server knows this secret (foobar)
	tokenstring, err := token.SignedString([]byte("foobar"))
	if err != nil {
		log.Fatalln(err)
	}
	return tokenstring
}
