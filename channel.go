package main

import (
	"fmt"
	"reflect"
	"time"
)

func main() {
	ch7()
}

func ch7() {
	done1 := make(chan bool)
	done2 := make(chan bool)

	go run1(done1)
	go run2(done2)
EXIT:
	for {
		select {
		case <-done1:
			println("run1 완료")
		case <-done2:
			println("run2 완료")
			break EXIT
		}
	}
}

func run1(done chan bool) {
	time.Sleep(1 * time.Second)
	done <- true
}
func run2(done chan bool) {
	time.Sleep(2 * time.Second)
	done <- true
}
func ch6() {
	ch := make(chan int, 3)
	ch <- 1
	ch <- 2
	close(ch)
	// for i := range ch {
	// 	println(i)
	// }

	for {
		if i, success := <-ch; success {
			println(i)
		} else {
			break
		}
	}
}
func ch5() {
	ch := make(chan string, 1)
	sendChan(ch)
	receiveChan(ch)
}

func sendChan(ch chan<- string) {
	ch <- "Data"
}
func receiveChan(ch <-chan string) {
	data := <-ch
	fmt.Println(data)
}

func ch4() {
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	close(ch)
	println(<-ch)
	// println(<-ch)
	if i, success := <-ch; success {
		println("더 이상 데이타 없음.")
		println(i)
	}
}

func ch3() {
	c := make(chan int, 2) //2 안주면 deadlock
	c <- 1
	c <- 2
	fmt.Println(<-c)
	fmt.Println(<-c)
}

func ch2() {

	done := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Print(i, "\n")
			time.Sleep(time.Microsecond * 1)
		}
		done <- 1
	}()
	<-done
}

func ch1() {

	ch := make(chan int)
	typeOfch := reflect.TypeOf(ch)
	fmt.Println(typeOfch)

	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
	}()

	ch_result := make([]int, 10)
	for i := 0; i < 10; i++ {
		ch_result[i] = <-ch
		println(ch_result[i])
	}
}
