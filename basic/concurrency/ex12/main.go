package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	var data int = 0
	sw := sync.WaitGroup{}
	sw.Add(3)

	rwMutex := new(sync.RWMutex)

	//mutex := sync.Mutex{}
	go func() {
		for i := 0; i < 10000; i++ {
			rwMutex.Lock()
			data += 1
			runtime.Gosched()
			fmt.Println("write : ", data)
			rwMutex.Unlock()
			time.Sleep(10 * time.Millisecond)
		}
		sw.Done()
	}()

	go func() {
		for i := 0; i < 10000; i++ {
			rwMutex.RLock()
			fmt.Println("read : ", data)
			runtime.Gosched()
			rwMutex.RUnlock()
			time.Sleep(1 * time.Second)
		}
		sw.Done()
	}()
	go func() {
		for i := 0; i < 10000; i++ {
			rwMutex.RLock()
			fmt.Println("read : ", data)
			runtime.Gosched()
			rwMutex.RUnlock()
			time.Sleep(2 * time.Second)
		}
		sw.Done()
	}()
	sw.Wait()
	fmt.Println(data)
}
