package main

import (
	"io"
	"io/ioutil"
	"os"
)

func util_ioutil() {
	bytes, err := ioutil.ReadFile("README.md")
	if err != nil {
		panic(err)
	}

	println(string(bytes))

	err = ioutil.WriteFile("README_OUT2.txt", bytes, 0)
	if err != nil {
		panic(err)
	}
}

func util_os() {

	//입력파일 열기
	fi, err := os.Open("README.md")
	if err != nil {
		panic(err)
	}
	defer fi.Close()

	//출력파일 생성
	fo, err := os.Create("README_OUT.txt")
	if err != nil {
		panic(err)
	}
	defer fo.Close()

	buff := make([]byte, 1024)

	for {
		cnt, err := fi.Read(buff)
		if err != nil && err != io.EOF {
			panic(err)
		}
		if cnt == 0 {
			break
		}
		_, err = fo.Write(buff[:cnt])
		if err != nil {
			panic(err)
		}
	}
}
