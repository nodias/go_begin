package io

import (
	"bufio"
	"io"
	"log"
	"os"
	"strings"
)

func ReadHtml(htmlAdress string) []byte {
	htmlString := []string{}
	htmlFile, err := OpenFile(htmlAdress)
	if err != nil {
		log.Fatal(err)
	}
	htmlToByte := func(scannerText string) {
		htmlString = append(htmlString, scannerText)
	}
	ReadFrom(htmlFile, htmlToByte)

	htmlByte := strings.Join(htmlString, "\n")
	return []byte(htmlByte)
}

func OpenFile(fileAdress string) (*os.File, error) {
	file, err := os.OpenFile(fileAdress, os.O_CREATE|os.O_RDWR, os.FileMode(0644))
	if err != nil {
		return nil, err
	}
	return file, nil
}

func ReadFrom(file io.Reader, f func(string)) {
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		f(scanner.Text())
	}
}
