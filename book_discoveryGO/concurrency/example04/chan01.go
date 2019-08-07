package example04

import (
	"fmt"
	"time"
)

var c = make(chan int)
var r int

func A(){
	c <- 1
}

func B(){
	r = <- c
	fmt.Println(r)
}

func C()int{
	return 12
}

func MyWay(){
	fmt.Println("i don't care")
}

func sendChan(ch chan string, done chan bool){
	ch<-"I'm a send Channel"
	//done <- true
	//fmt.Println(<-ch)
	//<-done
}

func receiveChan(ch chan string, done chan bool){
	fmt.Println(<-ch)
	//<-done
	//ch<- "I'm a send Channer 2"
	//done <- true
}

func run1(done chan bool){
	time.Sleep(1*time.Second)
	done<-true
}

func run2(done chan bool){
	time.Sleep(2*time.Second)
	done<-true
}

