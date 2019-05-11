package main

// func main() {
// 	// slice()
// 	// array()
// 	// slice3()
// 	slice4()
// }

func array() {
	var multiArray [5][5]int
	// println(multiArray)
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			multiArray[i][j] = j + 1 + i*5
		}
	}
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			print(multiArray[i][j], "\t")
		}
		print("\n")
	}
}

func slice() {
	var sliceArray [][]int = make([][]int, 5)
	for i := 0; i < 5; i++ {
		sliceArray[i] = make([]int, 5)
		for j := 0; j < 5; j++ {
			sliceArray[i][j] = i*5 + j + 1
		}
	}
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			print(sliceArray[i][j], "\t")
		}
		println()
	}
}

func slice2() {
	var sliceArray []int
	for i := 0; i < 25; i++ {
		sliceArray = append(sliceArray, i)
		print(sliceArray[i]+1, "\t")
		if (i+1)%5 == 0 {
			print("\n")
		}
	}
}

func slice3() {
	var sliceArray = []int{}
	for i := 0; i < 100; i++ {
		sliceArray = append(sliceArray, i)
		println(sliceArray)
	}
}

//24 ~ 32
func slice4() {
	var sliceArray = []int{}
	for i := 0; i < 100; i++ {
		sliceArray = append(sliceArray, i)
	}
	partialArray := sliceArray[24:33]

	for _, i := range partialArray {
		println(i)
	}
}
