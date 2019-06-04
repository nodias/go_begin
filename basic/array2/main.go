package main

import "fmt"

func main() {
	var arr [3]string
	fmt.Println(arr[0] == "")

	var sli []string
	fmt.Println(sli == nil)
	sli = make([]string, 3)
	fmt.Println(sli[2])
}
