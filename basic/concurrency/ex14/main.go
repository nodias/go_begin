package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	var mutex = new(sync.Mutex)
	var cond = sync.Cond{
		L: mutex,
	}
	var data = "before"

	c := make(chan bool, 3)
	defer close(c)

	for i := 9; i >= 0; i-- {
		go func(n int) {
			fmt.Println("1 : ", data)
			mutex.Lock()
			c <- true
			fmt.Println("wait begin : ", n)
			data = "middle"
			//cond.Wait()
			fmt.Scanln()
			fmt.Println("1 wait end : ", n)
			fmt.Println("1 in for : ", data)
			mutex.Unlock()
		}(i)
	}

	for i := 0; i < 10; i++ {
		<-c
	}

	//time.Sleep(time.Second)
	go func(){
		fmt.Println("2.1 : ",data)
		time.Sleep(time.Second)
		data = "after"
		fmt.Println("2.2 : ",data)
	}()

	defer fmt.Println("2.defer : " + data)
	//순서대로 깨운다?
	for i := 0; i < 2; i++ {
		mutex.Lock()
		fmt.Println("2.signal : ", i)
		fmt.Println(fmt.Scanln())
		cond.Signal()
		time.Sleep(time.Second)
		mutex.Unlock()
	}
	fmt.Println("3 : ",data)
	fmt.Scanln()
}
