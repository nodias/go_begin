package main

import (
	"fmt"
	"net"
)

func main() {
	ln, err := net.Listen("tcp", ":5000")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer ln.Close()
	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}
		defer conn.Close()
		go requestHandler(conn)
	}
}

func requestHandler(c net.Conn) {
	data := make([]byte, 4096) //4096 크기의 데이터를 생성
	for {
		n, err := c.Read(data)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(string(data[:n]))
		_, err = c.Write(data[:n])
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}
