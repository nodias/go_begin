package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	data := FileRead("helloJson.json")
	jsonData := map[string]interface{}{}
	jsonByte := []byte{}
	json.Unmarshal(data, &jsonData)
	fmt.Println(jsonData["you"])
	fmt.Println(jsonData["me"])
	jsonByte, _ = json.Marshal(jsonData)
	fmt.Println(string(jsonByte))
}

// FileWrite is write to file
func FileWrite(fileAddr string, wordsToWrite string) {
	file, _, err := Open(fileAddr, os.O_APPEND)
	if err != nil {
		log.Fatal(err)
	}

	err = WriteTo(file, wordsToWrite)
	if err != nil {
		log.Fatal(err)
	}
}

// FileRead is return []byte
func FileRead(fileAddr string) []byte {
	file, fileInfo, err := Open(fileAddr, os.O_APPEND)
	if err != nil {
		log.Fatal(err)
	}

	fileSize := fileInfo.Size()
	b := make([]byte, fileSize)

	err = ReadFrom(file, b)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("content : ", string(b))
	return b
}

// Open is open file
func Open(fileAddr string, liveOrDead int) (*os.File, os.FileInfo, error) {
	f, err := os.OpenFile(fileAddr, os.O_CREATE|os.O_RDWR|liveOrDead, os.FileMode(0644))
	if err != nil {
		log.Println("failure to open")
		return nil, nil, err
	}
	fstat, err := f.Stat()
	if err != nil {
		log.Println("failure to load fileInfo")
		return nil, nil, err
	}
	return f, fstat, err
}

// WriteTo is write
func WriteTo(iw io.Writer, s string) error {
	if s == "" || iw == nil {
		log.Fatal("nil")
	}
	w := bufio.NewWriter(iw)
	n, _ := w.Write([]byte(s))
	err := w.Flush()
	if err != nil {
		log.Println("failure to write")
		return err
	}
	log.Println(n, "success to write")
	return nil
}

// ReadFrom is read
func ReadFrom(ir io.Reader, b []byte) error {
	if b == nil || ir == nil {
		log.Fatal("nil")
	}
	w := bufio.NewReader(ir)
	n, err := w.Read(b)
	if err != nil {
		log.Println("failure to read")
		return err
	}
	log.Println(n, "success to read")
	return nil
}
