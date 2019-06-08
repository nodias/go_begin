package main

import "fmt"

var num = 5

func a() {
	num = 6
}

func b() {
	fmt.Println(num)
}

func main() {
	a()
	fmt.Println(num)
	b()
	c()
	fmt.Println(num)
}
