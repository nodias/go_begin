package main

func main() {

}

func FibonacciGenerator(max int) func() int {
	next, a, b := 0,0,1
	return func() int{
		next, a, b = a,b, a+b
		if next > max {
			return -1
		}
		return next
	}
}
