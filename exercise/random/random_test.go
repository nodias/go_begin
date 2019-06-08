package main

import "fmt"

func Example_BoxsCountIs() {
	fmt.Println(BoxsCountIs(7))
	fmt.Println(BoxsCountIs(8))
	fmt.Println(BoxsCountIs(15))
	fmt.Println(BoxsCountIs(100))

	//Output:
	//1
	//2
	//3
	//15
}
