package array

import "fmt"

func Example_array() {

	fruits := []string{"apple", "banana", "kiwi"}
	fmt.Println(fruits)
	fmt.Println(fruits[1:2])

	fruits2 := fruits
	fruits2[2] = "fish"
	fmt.Println(fruits2)
	fmt.Println(fruits)

	fruits3 := make([]string, len(fruits))
	for i, r := range fruits {
		fruits3[i] = r
	}
	fmt.Println(fruits3)
	fruits3[2] = "elephant"
	fmt.Println(fruits3)
	fmt.Println(fruits)
}

func Example_copySlice() {
	fruits := []string{"apple", "banana", "kiwi"}

	dummy := make([]string, len(fruits))
	copy(dummy, fruits)

	if n := copy(dummy, fruits); n != len(fruits) {
		fmt.Println("fail copy")
	}
	fmt.Println(dummy)

	src := []int{30, 20, 50, 10, 40}
	dest := append([]int(nil), src...)
	wow := []int(nil)
	fmt.Println("22", wow)
	println(dest)
	println(src)
	//Output:
}
