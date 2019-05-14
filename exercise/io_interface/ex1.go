package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	ex2()
}


func ex2(){
	file, err := os.OpenFile("hello.txt")
}

func ex() {
	file, err := os.OpenFile(
		"hello.txt", os.O_CREATE|os.O_RDWR|os.O_TRUNC, os.FileMode(0644),
	)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	w := bufio.NewWriter(file)
	w.WriteString("Hello, world!")
	w.Flush()

	r := bufio.NewReader(file)
	fi, _ := file.Stat()
	b := make([]byte, fi.Size())
	n, err := r.Read(b)
	if err != nil {
		fmt.Println(err)
		return
	}
	file.Seek(0, os.SEEK_SET)
	fmt.Println(n, "buffer를 읽었습니다.")
	r.Read(b)

	fmt.Println(string(b))
}
