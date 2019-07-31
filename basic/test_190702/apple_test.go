package main

import "fmt"

func Example_Apple() {
	a := NewApple()
	println(a)
	a.Color = "blue"
	fmt.Println(a)
	b := NewApple()
	println(b)
	fmt.Println(b)
	//Output:
}

func Example_Apple2() {
	a := NewApple()
	println(a)
	fmt.Println(a)
	//Output:
}
