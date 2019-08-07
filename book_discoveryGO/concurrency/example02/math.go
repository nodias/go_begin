package main

func Complexity(n int) float32 {
	f := float32(n)
	return f * (f - 1) / 2
}

func ThreadNum(th, n int) int {
	return (n-1)/th + 1
}

func EfficientThreadNumber(n int) int {
	for th:=1;th<=n/2;th++{
		Complexity(ThreadNum(th, n))
		Complexity(th)
			if Complexity(ThreadNum(th, n)) - Complexity(ThreadNum(th+1, n)) < Complexity(th+1)-Complexity(th) {
				return th
			}
	}
	return n/2
}