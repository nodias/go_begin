package main

import "fmt"

var num = 5
var str = "str!"

func d(n *int) {
	fmt.Println("in d, &int :", n)
	fmt.Println("in d, *&int :", *n)
	*n = 22
	fmt.Println("in d, &int :", n)
}

func e(n *string) {
	fmt.Println("in d, &int :", n)
	fmt.Println("in d, *&int :", *n)
	*n = "not int!"
	fmt.Println("in d, &int :", n)
}

func main() {
	//num
	w := &num
	fmt.Println("&w :", w)
	fmt.Println("*&w :", *w)
	*w = 15
	fmt.Println("after w : ", num)

	d(&num)
	fmt.Println("final num : ", num)
	fmt.Println()

	//string
	h := &str
	fmt.Println("&h :", h)
	fmt.Println("*&h :", *h)
	*h = "not enough mana"
	fmt.Println("after h : ", str)

	e(&str)
	fmt.Println("final str : ", str)

	//immutable
	fmt.Println("index : ", str[0])
	bstr := []byte(str)
	fmt.Println(bstr)
	bstr[0] = 116
	fmt.Println(string(bstr))
}
