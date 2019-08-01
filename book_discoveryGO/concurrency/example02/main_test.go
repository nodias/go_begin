package main

import (
	"fmt"
)

func ExampleMin() {
	a:= ArrayGenerator()
	minNum := Min(a)
	fmt.Println(minNum)

	//Output:
	//1
}

func ExampleParallelMin() {
	a:= ArrayGenerator()
	minNum := ParallelMin(a, 2)
	fmt.Println(minNum)

	//Output:
	//1
}

func ExampleArrayGenerator() {
	a := ArrayGenerator()
	fmt.Println(len(a))

	//Output:
	//10000
}

func ExampleAnything(){
	fmt.Println((97+3-1)/3)

	//Output:
	//
}
