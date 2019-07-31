package main_test

import (
	"fmt"
	"log"
	"testing"

	m "github.com/nodias/go_begin/exercise/jwtEx3"
)

func Test_GenerateToken(t *testing.T) {
	user := m.User{
		"nodias",
		"password123",
		"19910404",
		"nodias@naver.com",
	}
	tokenString, err := m.GenerateToken(user)
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println(tokenString)
	return
}
