package main

func main() {
	sli := []string{"one", "two", "three"}
	println("sli : ", sli)
	println("&sli : ", &sli)
	println("&sli[0] : ", &sli[0])

	pbvsli(&sli)

	arr := [3]string{"one", "two", "three"}
	// println("arr : ", arr)
	println("&arr : ", &arr)
	println("&arr[0] : ", &arr[0])

	pbvarr(&arr)
}
func pbvsli(sli *[]string) {
	println("pbv sli : ", sli)
	println("pbv &sli : ", &sli)

	// println("pbv sli : ", sli[0])
	// println("pbv &sli : ", &sli[0])
}

func pbvarr(arr *[3]string) {
	println("pbv arr : ", arr)
	println("pbv &arr : ", &arr)

	println("pbv arr[0] : ", arr[0])
	println("pbv &arr[0] : ", &arr[0])
}
