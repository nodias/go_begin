package main

import (
	"fmt"
	"sync"
	"time"
)

func main(){
	data := 10
	mutex := sync.RWMutex{}

	go func (){
		mutex.RLock()
		for i := 0; i < 100; i++ {
			fmt.Println(data)
			time.Sleep(time.Microsecond*10000)
		}
		mutex.RUnlock()
	}()

	go func(){
		//time.Sleep(time.Millisecond*1)
		mutex.Lock()
		data += 10
		mutex.Unlock()
	}()


	go func (){
		mutex.RLock()
		for i := 0; i < 100; i++ {
			fmt.Println(data)
			time.Sleep(time.Microsecond*10000)
		}
		mutex.RUnlock()
	}()
	fmt.Scanln()


}