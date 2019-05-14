package main

import "fmt"

func main() {
	var num1 int = 10
	// var num2 float32 = 3.2
	// var num3 complex64 = 2.5 + 8.1i
	// var s string = "Hello, World"
	// var b bool = true
	// var a []int = []int{1, 2, 3}
	// var m map[string]int = map[string]int{"Hello": 1}
	// var p *int = new(int)
	// type Data struct{ a, b int }
	// var data = Data{1, 2}
	// var i interface{} = 1

	// fmt.Printf("%d\n", num1)
	// fmt.Printf("%f\n", num2)
	// fmt.Printf("%f\n", num3)
	// fmt.Printf("%s\n", s)
	// fmt.Printf("%t\n", b)
	// fmt.Printf("%v\n", a)
	// fmt.Printf("%v\n", m)
	// fmt.Printf("%p\n", p)
	fmt.Printf("%#v\n", &num1)
	fmt.Printf("%#v\n", num1)
	// fmt.Printf("%v\n", data)
	// fmt.Printf("%v\n", i)
	func(i *int) {
		fmt.Printf("%#v\n", i)
	}(&num1)
}
