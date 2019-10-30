package main

import "fmt"

var globalVariable = 1

func makeClosure(freeVariable int)func(boundVariable int) int{
	return func(boundVariable int)int {
		return globalVariable + freeVariable + boundVariable
	}
}
func main(){
	closure := makeClosure(2)
	res := closure(3)
	fmt.Println(res)
	globalVariable = 5
	res = closure(3)
	fmt.Println(res)
}



