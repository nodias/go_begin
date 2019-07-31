package main

import "log"

func recoverPanic (){
	if r := recover(); r != nil {
		log.Fatal(r)
	}
}

func Must(v interface{}, err error){
	if err != nil {
		panic(err)
	}
}
