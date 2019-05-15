package main

import (
	"time"
)

// func main() {
// 	var wait sync.WaitGroup
// 	wait.Add(3)

// 	go func() {
// 		defer wait.Done()
// 		for i := 0; i < 10; i++ {
// 			println(i)
// 			time.Sleep(time.Second * 1)
// 		}
// 	}()

// 	go func() {
// 		defer wait.Done()
// 		for i := 0; i < 10; i++ {
// 			println(i)
// 			time.Sleep(time.Second * 1)
// 		}
// 	}()

// 	go func() {
// 		defer wait.Done()
// 		say2("said")
// 	}()
// 	for i := 0; i < 10; i++ {
// 		println(i)
// 		time.Sleep(time.Second * 1)
// 	}
// 	wait.Wait()
// 	for i := 0; i < 10; i++ {
// 		println("after", i)
// 		time.Sleep(time.Second * 1)
// 	}
// }

func say2(str string) {
	for i := 0; i < 10; i++ {
		println(i)
		time.Sleep(time.Second * 1)
	}
}
