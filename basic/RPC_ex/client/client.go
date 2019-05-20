package main

import (
	"log"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:5000")
	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Close()
}
