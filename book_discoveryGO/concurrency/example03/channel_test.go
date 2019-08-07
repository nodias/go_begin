package main

import (
	"fmt"
	"runtime"
)

func Example_simpleChannel(){
	c:= make(chan int)
	go func() {
		c<-1
		c<-2
		c<-3
	}()
	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println(<-c)
	//Output:
	//
}

func Example_simpleChannel2(){
	c:= make(chan int)
	go func() {
		c<-1
		c<-2
		c<-3
		close(c)
	}()
	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println(runtime.NumCPU())
	//Output:
	//

}
