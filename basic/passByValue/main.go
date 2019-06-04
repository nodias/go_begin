package main

import "fmt"

func main() {
	PassString()
	PassArray()
	PassSlice()
	PassMap()
}

func PassString() {
	str := "1,2,3"
	fmt.Println("string before : ", str)
	// StringEx(str)
	// StringEx2(str)
	// StringEx3(str)
	// StringEx4(&str)
	StringEx5(&str)
	fmt.Println("string after : ", str)
}

func PassArray() {
	arr := [3]int{1, 2, 3}
	fmt.Println("array before : ", arr)
	// ArrayEx(arr)
	// ArrayEx2(arr)
	// ArrayEx3(arr)
	// ArrayEx4(&arr)
	ArrayEx5(&arr)
	fmt.Println("array after : ", arr)
}

func PassSlice() {
	sli := []int{1, 2, 3}
	fmt.Println("slice before : ", sli)
	// SliceEx(sli)
	// SliceEx2(sli)
	// SliceEx3(sli)
	// SliceEx4(&sli)
	SliceEx5(&sli)
	fmt.Println("slice after : ", sli)
}

func PassMap() {
	m := map[int]int{1: 1, 2: 2, 3: 3}
	fmt.Println("map before : ", m)
	// MapEx(m)
	// MapEx2(m)
	// MapEx3(m)
	// MapEx4(&m)
	MapEx5(&m)
	fmt.Println("map after : ", m)
}
