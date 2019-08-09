package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan int)
	c2 := make(chan string)

	go func() {
		for i := 0; true; i++ {
			c1 <- i
			time.Sleep(time.Second)
		}
	}()

	go func() {
		for i := 0; true; i++ {
			c2 <- fmt.Sprintf("hey channel!+%d", i)
			time.Sleep(time.Second*3)
		}
	}()

EXIT:
	for {
		select {
		case i := <-c1:
			fmt.Println(i)
		case s := <-c2:
			fmt.Println(s)
		case <-time.After(time.Second * 10):
			break EXIT
		}
	}
}
