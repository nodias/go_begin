package main

import (
	"fmt"
	. "time"
)

func sum(a, b int, c chan int){
	Sleep(Second *3)
	c <- a+b
}


func main(){
	c:= make(chan int)
	go sum(1,3, c)
	res := <-c
	fmt.Println(res)
}