package closure

import "fmt"

//var str int = 15

// Closed expression
//func a() func(i int) int {
//	return func(i int) int {
//		return i + str
//	}
//}
//
//func GetA() func(i int) int {
//	return a()
//}

// free a, b
// bound c

func B(b bool) func(fr string) []string {
	var arr []string
	if b {
		return func(fr string) []string {
			for i := 0; i < 5; i++ {
				arr = append(arr, fmt.Sprintf("%s : %d", fr, i))
			}
			return arr
		}
	}
	return func(fr string) []string {
		arr = append(arr, fmt.Sprintf("%s : 0", fr))
		return arr
	}
}
