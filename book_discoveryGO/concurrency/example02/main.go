package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

func main() {

}

func ArrayGenerator(n int) (a []int) {
	for i := 0; i < n; i++ {
		a = append(a, rand.Intn(10000000))
	}
	return a
}

func Min(a []int) int {
	fmt.Println("sleep$")
	if len(a) == 0 {
		return 0
	}
	min := a[0]
	a = a[1:]
	for _, i := range a {
		time.Sleep(time.Second/10)
		if min > i {
			min = i
		}
	}
	return min
}

func ParallelMin(a []int, th int) int {
	runtime.GOMAXPROCS(4)
	fmt.Println("@")
	if len(a) < th {
		return Min(a)
	}
	mins := make([]int, th)
	size := (len(a) + th-1) / th
	var wg sync.WaitGroup
	for i := 0; i < th; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			begin, end := i*size, (i+1)*size
			if end > len(a) {
				end = len(a)
			}
			fmt.Println("#", i)
			mins[i] = Min(a[begin:end])
		}(i)
	}
	wg.Wait()
	return Min(mins)
}
