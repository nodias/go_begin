package main

import (
	"math/rand"
	"sync"
)

func main() {

}

func ArrayGenerator() (a []int) {
	for i := 0; i < 10; i++ {
		a = append(a, rand.Intn(10000))
	}
	return a
}

func Min(a []int) int {
	if len(a) == 0 {
		return 0
	}
	min := a[0]
	a = a[1:]
	for _, i := range a {
		if min > i {
			min = i
		}
	}
	return min
}

func ParallelMin(a []int, n int) int {
	if len(a) < n {
		return Min(a)
	}
	mins := make([]int, n)
	size := (len(a) + n-1) / n
	var wg sync.WaitGroup
	for i := 0; i < n; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			begin, end := i*size, (i+1)*size
			if end > len(a) {
				end = len(a)
			}
			mins[i] = Min(a[begin:end])
		}(i)
	}
	wg.Wait()
	return Min(mins)
}
