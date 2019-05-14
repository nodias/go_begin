package main

import (
	"fmt"
	"os"
)

func main() {
	f2()
}

func f2() {
	var num1 int
	var num2 float32
	var s string

	file1, _ := os.Open("hello1.txt")
	defer file1.Close()
	fmt.Fscan(file1, &num1, &num2, &s)
	println(num1, num2, s, "123", "123")
}

func f1() {
	file1, _ := os.Create("hello1.txt")
	defer file1.Close()
	fmt.Fprint(file1, 1, 1.1, "Hello, World!")

	file2, _ := os.Create("hello3.txt")
	defer file2.Close()
	fmt.Fprint(file2, 1, 1.1, "Hello, World!")
}
