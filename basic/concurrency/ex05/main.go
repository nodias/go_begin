package main

import (
	"fmt"
	"sync"
)

func main(){
	ch := make(chan int)
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		go producer(ch)
		i:=consumer(ch)
		po(i)
		wg.Done()
	}()
	wg.Wait()
}

func consumer(ch <-chan int) int{
	return <-ch
}

func producer(ch chan<- int) {
	ch<-99
}

func po(i int){
	fmt.Println(i)
}
