package main

import (
	"fmt"
	"strings"
)

func main() {
	s1 := []string{"Hello,", "World"}
	fmt.Println(strings.Join(s1, " "))
	s2 := strings.Split("Hello, World!", " ")
	fmt.Println(s2[1])
	s3 := strings.Fields("Hello, World!")
	fmt.Println(s3[0])

}
