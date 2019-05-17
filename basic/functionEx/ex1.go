package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

func main() {
	ExampleNewIntGenerator()
}

func NewIntGenerator() func() int {
	var i int
	return func() int {
		i++
		return i
	}
}

func ExampleNewIntGenerator() {
	gen := NewIntGenerator()
	fmt.Println(gen(), gen(), gen(), gen())
}

func pointerTest() {
	sli := make([]int, 4)
	sli[0] = 1
	sli[1] = 2
	sli[2] = 3
	sli[3] = 4
	// sliceTest(&sli)
	println(&sli)

	updateSlice := func() {
		sli = append(sli, 5)
		println(&sli)
	}
	closureTest(updateSlice)
	println(&sli)
	// fmt.Println(sli[4])
}

func sliceTest(sli *[]int) {
	*sli = append(*sli, 5)
	fmt.Println(sli)
}
func closureTest(f func()) {
	f()
}

func ExampleReadFrom() {
	readThis := "@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@Ldmnlskancwqignpwqnvpwvnkasvlsanvlksanvlksnalkvnlkwnvoqwvnwlnklwnvpwqjpjfpwqnflanvpsajfpwqjflwqflwjqlknflwkqnflkqnf\n@"
	r := strings.NewReader(readThis)
	var readOut = []string{}
	closure := func(scannerText string) {
		readOut = append(readOut, scannerText)
	}
	ReadFrom(r, closure)
	for _, i := range readOut {
		fmt.Println(i)
	}
}

func ReadFrom(r io.Reader, closure func(scannerText string)) error {
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		closure(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return err
	}
	return nil
}
