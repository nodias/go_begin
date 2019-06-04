package main_test

import (
	"fmt"
	"testing"

	cont "github.com/nodias/go_begin/exercise/jwtEx2"
)

func Test_createTokenString(t *testing.T) {
	user := cont.User{
		Name: "kss",
		Age:  29,
	}
	tokenstring := cont.CreateTokenString(&user)
	fmt.Println(tokenstring)
}
