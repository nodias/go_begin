package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	done := make(chan bool, 3)
	count := 5

	sw := sync.WaitGroup{}
	sw.Add(1)
	go func() {
		go A(done, count)
		B(done, count)
		sw.Done()
	}()
	sw.Wait()
}

func A(done chan bool, count int) {
	func() {
		for i := 0; i < count; i++ {
			done <- true
			fmt.Println("Go routine : ", i)
		}
	}()
}

func B(done chan bool, count int) {
	for i := 0; i < count; i++ {
		<-done
		fmt.Println("Main : ", i)
		time.Sleep(1 * time.Second*2)
	}
}
