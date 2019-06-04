package main

func main() {
	ArrayEx()
	SliceEx()
}

func ArrayEx() {
	arr := [3]string{"one", "two", "three"}
	// println(arr)
	println("&arr : ")
	println(&arr)
	println("&arr[0] : ")
	println(&arr[0])
	println("&arr[1] : ")
	println(&arr[1])

	passByValueArray(arr)
	passByReferenceArray(&arr)
}

func SliceEx() {
	sli := []string{"one", "two", "three"}
	println("sli : ")
	println(sli)
	println("&sli : ")
	println(&sli)
	println("&sli[0] : ")
	println(&sli[0])
	println("&sli[1] : ")
	println(&sli[1])

	passByValueSlice(sli)
	passByReferenceSlice(&sli)
}

func passByValueArray(arr [3]string) {
	println("passByValueArray : ")
	println(&arr[0])
}

func passByReferenceArray(arr *[3]string) {
	println("passByReferenceArray : ")
	println(&arr[0])
}

func passByValueSlice(sli []string) {
	println("passByValueSlice : ")
	println(&sli[0])
}

func passByReferenceSlice(sli *[]string) {
	println("passByReferenceSlice : ")
	// println(&sli[0])
}
