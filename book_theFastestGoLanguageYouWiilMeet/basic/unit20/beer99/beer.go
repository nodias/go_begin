package main

import "fmt"

var phrase = "%d %s of beer on the wall, %d %s of beer.\n"
var phraseHave = "Take one down, pass it around, %d %s of beer on the wall.\n"
var phraseEmpty = "No more bottles of beer on the wall, No more bottles of beer\n" +
	"Go to the store and buy some more, 99 bottles of beer on the wall\n"

const (
	BOTTLES = "bottles"
	BOTTLE  = "bottle"
)

func main() {
	for j := 2; j > 0; j-- {
		for i := 99; i > 0; i-- { //100 ~ 1
			switch {
			case i == 2:
				fmt.Printf(phrase, i, BOTTLES, i, BOTTLES)
				fmt.Printf(phraseHave, i-1, BOTTLE)
			case i == 1:
				fmt.Printf(phrase, i, BOTTLE, i, BOTTLE)
				fmt.Printf(phraseEmpty)
			default:
				fmt.Printf(phrase, i, BOTTLES, i, BOTTLES)
				fmt.Printf(phraseHave, i-1, BOTTLES)
			}
		}
	}
}

//99 복수 99 복수 -99
//98 복수
//.
//.
//.
//2 복수 2 복수 -2
//1 단수
//1 단수 1 단수 -1
//empty 99 복수
