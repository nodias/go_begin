package main

import (
	"fmt"
)

func ExampleMin() {
	a := ArrayGenerator(1000)
	minNum := Min(a)
	fmt.Println(minNum)

	//Output:
	//1
}

func ExampleParallelMin() {
	a := ArrayGenerator(1000)
	fmt.Println(a)
	minNum := ParallelMin(a, 5)
	fmt.Println(minNum)

	//Output:
	//1
}

func ExampleArrayGenerator() {
	a := ArrayGenerator(1000)
	fmt.Println(len(a))

	//Output:
	//10000
}

// TotalComplexity is a function. To calculate total complexity
// n is number of transaction
// th is number of thread
// return is TimeComplexity
func TotalComplexity(th, n int) float32 {
	fmt.Printf("@@ 쓰레드 개수 : %d, 트랜잭션 개수 : %d\n", th, n)
	if th > n/2 {
		fmt.Println("number of thread must be less than half the number of n")
	}
	var complexityA float32 = Complexity(ThreadNum(th, n))
	var complexityB float32 = Complexity(th)
	fmt.Println("복잡도 A :", complexityA)
	fmt.Println("복잡도 B :", complexityB)

	return complexityA + complexityB
}

func AllComplexity(n int) {
	for i := 1; i <= n/2; i++ {
			res := TotalComplexity(i, n)
			fmt.Printf("시간복잡도 : %f\n", res)
	}
}

func Example_ViewEfficient(){
	AllComplexity(111)
	//Output:
	//
}

var n= 100
var arr = ArrayGenerator(n)

func Example_EfficientlyParallelMin(){
	ParallelMin(arr, EfficientThreadNumber(n))
	//Output:
	//
}

func Example_Un_efficientlyParallelMin(){
	ParallelMin(arr, n/2)
	//Output:
	//
}

func Example_JustMin(){
	Min(arr)
	//Output:
	//
}


