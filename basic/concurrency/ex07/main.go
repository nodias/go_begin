package main

import (
	"fmt"
	"time"
)

func main(){
	c1 := make(chan int)
	c2 := make(chan int)

	go func(){
		c1 <-1
	}()

	for i:=0; i<5; i++{
		go func(){
			c2 <-2
		}()
	}
	for {
		select {
		case i := <-c1:
			fmt.Println("c1 : ", i)
			case i:= <- c2:
				fmt.Println("c2 : ", i)
		case <- time.After(time.Second*3):
			fmt.Println("timeout")
		default:
			time.Sleep(time.Second*2)
			fmt.Println("def@")
		}
	}
}