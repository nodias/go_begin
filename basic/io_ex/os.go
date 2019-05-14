package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	os1()
}

func os1() {
	file1, err := os.Create("hello1.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file1.Close()
	fmt.Fprint(file1, 1, 1.1, "Hello, World")

	file2, err := os.Create("hello2.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file2.Close()
	fmt.Fprintln(file2, 1, 1.1, "Hello, world")

	file3, err := os.Create("hello3.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file3.Close()

	fmt.Fprintf(file3, "%d,%f,%s", 1, 1.1, "Hello, World")

}
