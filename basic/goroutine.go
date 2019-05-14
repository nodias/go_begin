package main

import (
	"fmt"
	"time"
)

// func main() {
// 	// say("Sync")
// 	println("main")
// 	go say("Async1")
// 	go say("Async2")
// 	go say("Async3")
// 	time.Sleep(time.Second * 3)

// }

func say(s string) {
	for i := 0; i < 10; i++ {
		fmt.Println(s, "***", i)
		time.Sleep(time.Second * 1)
	}
}
