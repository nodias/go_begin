package main

import (
	"fmt"
	"math/rand"
)

func main() {
	// for {
	// 	o := rand.Intn(100) + 1
	// 	if o < 1 || o > 100 {
	// 		fmt.Println(o)
	// 		return
	// 	}
	// 	fmt.Println(o)
	// }

	oranges := rand.Intn(100) + 1
	boxs := BoxsCountIs(oranges)
	fmt.Printf("오렌지의 개수가 : %d개 일때 박스 개수는 %d개 \n", oranges, boxs)
}

func BoxsCountIs(oranges int) (boxes int) {
	boxes = (oranges-1)/7 + 1
	return
}
