package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	runtime.GOMAXPROCS(4)

	var data = []int{}
	var mutex = new(sync.Mutex)

	go func() {
		for i := 0; i < 1000000; i++ {
			mutex.Lock()
			data = append(data, i)
			mutex.Unlock()
			runtime.Gosched()
		}
	}()
	go func() {
		for i := 0; i < 1000000; i++ {
			mutex.Lock()
			data = append(data, i)
			mutex.Unlock()
			runtime.Gosched()
		}
	}()

	time.Sleep(time.Second * 10)

	fmt.Println(len(data))
}
