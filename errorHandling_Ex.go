package main

import (
	"errors"
	"fmt"
	"log"
)

// Error처리 example

// type error interface {
// 	Error()
// }

const (
	SUCCESS  = 1
	FAIL     = 2
	UNSTABLE = 3
)

func main() {
	err := myFunc(UNSTABLE)
	if err != nil {
		errorHandler(err)
	}
}

func errorHandler(err error) {
	switch err.(type) {
	default:
		println("ok")
	case *myErr:
		log.Println(err)
	case error:
		log.Fatal(err)
	}
}

type myErr struct {
	error
}

func (m *myErr) Error() string {
	return "myError : a is not 1 or 2"
}

func myFunc(a int) error {
	if a == SUCCESS {
		fmt.Println("Success")
		return nil
	} else if a == FAIL {
		return errors.New("Error : a is 2")
	} else {
		return new(myErr)
	}
}
