package example04

import (
	"fmt"
	"time"
)

func Example_test() {
	var done = make(chan bool)
	var res int
	go func() {
		res = C()
		done <- true
	}()
	MyWay()
	<-done

	fmt.Println(res)
	//Output:
	//
}

func Example_test2() {
	done := make(chan bool)
	ch := make(chan int, 1)
	go func() {
		ch <- 1
	}()
	go func() {
		time.Sleep(time.Second * 2)
		fmt.Println(<-ch)
		done <- true
	}()
	<-done
	//Output:
	//
}

func Example_SendandRecieveChan() {
	ch := make(chan string, 1)
	done := make(chan bool)

	go func() {
		receiveChan(ch, done)
		sendChan(ch, done)
		done <- true
	}()
	<-done
	//Output:
	//
}

func Example_test3() {
	done1 := make(chan bool)
	done2 := make(chan bool)
	go func() {
		go run1(done1)
		go run2(done2)
	}()
EXIT:
	for {
		select {
		case <-done1:
			println("run1 complete")
		case <-done2:
			println("run2 complete")
			break EXIT
		}
	}


	//Output:
	//
}

func Example_test4() {
	ch := make(chan int, 2)
	done := make(chan bool)
	//go routine 속의 <-chan 은 기다려지지 않기에 wait이 필요하다.
	go func() {
		fmt.Println(<-ch)
		fmt.Println(<-ch)
		fmt.Println(<-ch)
		fmt.Println(<-ch)
		fmt.Println(<-ch)
		fmt.Println(<-ch)
		fmt.Println(<-ch)
		fmt.Println(<-ch)
		fmt.Println(<-ch)
		fmt.Println(<-ch)
		done<-true
	}()
	go func() {
		ch <- 1
		ch <- 2
		ch <- 3
		ch <- 4
		ch <- 5
		ch <- 6
		ch <- 7
		ch <- 8
		ch <- 9
	}()
	<-done
	//close(ch)
	//for i:= range ch {
	//	println(i)
	//}

	//Output:
	//
}

func Example_test5() {
	ch := make(chan int, 2)
	done := make(chan bool)
	//go routine 속의 <-chan 은 기다려지지 않기에 wait이 필요하다.
	go func() {
		ch <- 9
	}()
	close(ch)
	go func() {
		fmt.Println(<-ch)
		done<-true
	}()
	<-done

	//Output:
	//
}

func Example_test6() {
	ch := make(chan int, 2)
	ch2 := make(chan string, 2)
	done := make(chan bool)
	//go routine 속의 <-chan 은 기다려지지 않기에 wait이 필요하다.
	//close 후에는 받는 것만 가능, stack소모 후 받게 되면 0, "" 등 빈값 출력
	//close(ch2)
	ch2<-"wow"
	a, b := <-ch2
	fmt.Println(a)
	fmt.Println(b)

	go func() {
		ch <- 9
		close(ch)
	}()

	go func() {
		fmt.Println(<-ch)
		fmt.Println(<-ch)
		//ch<-1
		done<-true
	}()
	<-done

	//Output:
	//
}
