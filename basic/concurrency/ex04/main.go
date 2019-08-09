package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	ch := make(chan int, 4)
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		go A(ch)
		B(ch)
		wg.Done()
	}()
	wg.Wait()
}

func A(ch chan int) {
	ch<- 99
	ch<-100
	close(ch)
}

func B(ch chan int) {
	f, ok := <-ch
	fmt.Println(f, ok)
	time.Sleep(time.Second*3)
	for i := range ch {
		fmt.Println(i)
	}
}
