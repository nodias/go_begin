package main

// test 1
func StringEx(str string) {
	str = "4,5,6"
}

func ArrayEx(arr [3]int) {
	arr = [3]int{4, 5, 6}
}

func SliceEx(sli []int) {
	sli = []int{4, 5, 6}
}

func MapEx(m map[int]int) {
	m = map[int]int{1: 4, 2: 5, 3: 6}
}

// test 2
func StringEx2(str string) {
	str = "4,5,6"
}

func ArrayEx2(arr [3]int) {
	arr[0] = 4
	arr[1] = 5
	arr[2] = 6
}

func SliceEx2(sli []int) {
	sli[0] = 4
	sli[1] = 5
	sli[2] = 6
}

func MapEx2(m map[int]int) {
	m[1] = 4
	m[2] = 5
	m[3] = 6
}

// test 3
func StringEx3(str string) {
	str = "4,5,6"
}

func ArrayEx3(arr [3]int) {
	arr = [3]int{4, 5, 6}
	arr[0] = 4
	arr[1] = 5
	arr[2] = 6
}

func SliceEx3(sli []int) {
	sli = []int{4, 5, 6}
	sli[0] = 4
	sli[1] = 5
	sli[2] = 6
}

func MapEx3(m map[int]int) {
	m = map[int]int{1: 4, 2: 5, 3: 6}
	m[1] = 4
	m[2] = 5
	m[3] = 6
}

// test 4
func StringEx4(str *string) {
	*str = "4,5,6"
}

func ArrayEx4(arr *[3]int) {
	*arr = [3]int{4, 5, 6}
}

func SliceEx4(sli *[]int) {
	*sli = []int{4, 5, 6}
}

func MapEx4(m *map[int]int) {
	*m = map[int]int{1: 4, 2: 5, 3: 6}
}

// test 5
func StringEx5(str *string) {
	*str = "4,5,6"
}

func ArrayEx5(arr *[3]int) {
	arr[2] = 4
	*arr[2] = 4
	*arr = [3]int{4, 5, 6}
	arr[0] = 7
	arr[1] = 8
	arr[2] = 9
}

func SliceEx5(sli *[]int) {
	*sli = []int{4, 5, 6}

}

func MapEx5(m *map[int]int) {
	*m = map[int]int{1: 4, 2: 5, 3: 6}
}
