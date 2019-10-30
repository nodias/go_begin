package closure1

func main() {
	calc(func(a, b int) int { return 5 })
}

func calc(a func(a, b int) int) func(x int) int {
	return func(x int) int {
		return 5
	}
}
