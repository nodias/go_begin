package randTest

import (
	"fmt"
	"math/rand"
)

type Dice struct {
	MaxNum int
}

func NewDice(n int) *Dice {
	return &Dice{n}
}

func (d Dice) throw() int {
	return rand.Intn(d.MaxNum) + 1
}

func ExampleRandom() {
	diceMaxNum := 8
	expectDiceNum := 8

	sum6 := 0
	for i := 1; i <= diceMaxNum; i++ {
		for j := 1; j <= expectDiceNum/2; j++ {
			if i+j == expectDiceNum {
				fmt.Println(i, j)
				sum6++
			}
		}
	}
	fmt.Printf("result is %d", sum6)
	//Output:
}
