package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	iu()
}

func iu() {
	s := "Hello, World!"

	err := ioutil.WriteFile("hello2.txt", []byte(s), os.FileMode(644))

	if err != nil {
		fmt.Println(err)
		return
	}
	data, err := ioutil.ReadFile("hello2.txt")
	fmt.Println(string(data))

}

func writeNRead() {
	file, err := os.OpenFile("hello.txt", os.O_CREATE|os.O_APPEND, os.FileMode(0644))
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	n := 0
	s := "안녕하세요"
	n, err = file.Write([]byte(s))
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(n, "바이트 저장 완료")

	fileinfo, err := file.Stat()
	if err != nil {
		fmt.Println(err)
		return
	}

	var data = make([]byte, fileinfo.Size())
	file.Seek(0, os.SEEK_SET)

	n, err = file.Read(data)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(n, "바이트 읽기 완료")
	fmt.Println(string(data))
}

func read() {
	file, err := os.Open("hello.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	fi, err := file.Stat()
	if err != nil {
		log.Fatal(err)
	}

	var data = make([]byte, fi.Size())
	n, err := file.Read(data)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(n, "바이트 읽기 완료")
	fmt.Println(string(data))
}

func write() {
	file, err := os.Create("hello.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	s := "Hello, world!"

	n, err := file.Write([]byte(s))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(n, "만큼의 데이터 저장")
}
