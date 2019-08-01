package main

import (
	"fmt"
	"log"
)

func recoverPanic (){
	if r := recover(); r != nil {
		fmt.Println("panic occured")
		log.Fatal(r)
	}
}

func Must(v interface{}, err error){
	if err != nil {
		panic(err)
	}
}
