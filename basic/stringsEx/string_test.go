package stringsEx

import (
	"fmt"
	"math/rand"
	"strings"
)

func Example_strings() {
	s := []string{"a", "b", "c"}
	sj := strings.Join(s, ",")
	fmt.Println(sj)
	//Output:
}

func Example_question() {
	arr := [10]int{}
	for i := 0; i < 100; i++ {
		arr[rand.Intn(10)+1]++
	}
	for i := range arr {
		fmt.Printf("%s : %s", i+1, arr[i])
	}
}
