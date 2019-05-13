package main

func main() {
	a()
}

func a() {
	var arr = [1]string{"t"}
	var sli = []string{"t"}
	println("# array")
	println(&arr)
	println(&arr[0])
	println("# slice")
	println(&sli)
	println(&sli[0])
	b(arr, sli)
}

func b(arr [1]string, sli []string) {
	println("# B")
	println("# array")
	println(&arr)
	println(&arr[0])
	println("# slice")
	println(&sli)
	println(&sli[0])
}
