package main

import (
	"log"
	"net"
)

func main() {
	ln, err := net.Listen("tcp", ":5000")
	if err != nil {
		log.Println(err)
		return
	}
	defer ln.Close()
	for {
		conn, err := ln.Accept()
		if err != nil {
			return
		}
		defer conn.Close()
		log.Println("Connected")
	}

}
