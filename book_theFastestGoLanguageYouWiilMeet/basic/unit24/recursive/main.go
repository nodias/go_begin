package main

import "fmt"

func main() {
	res := factorial(30)
	fmt.Println(res)
}

func factorial(n uint64) uint64 {
	if n == 1 {
		return 1
	}
	return n * factorial(n-1)
}
