package main

import "fmt"

func calculate(a,b int) <-chan int {
	ch := make(chan int, 4)
	go func(){
		ch <- a+b
		ch <- a-b
		ch <- a*b
		ch <- a/b
		close(ch)
	}()
	//ch<0
	return ch
}

func main(){
	ch := calculate(10,2)
	for c := range ch {
		fmt.Println(c)
	}
}

