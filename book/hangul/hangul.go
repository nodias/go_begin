package hangul

import "fmt"

var (
	start = rune(44032) //가
	end   = rune(55204) //힣
)

func HasConsonantSuffix(s string) bool {
	numEnds := 28
	result := false
	for _, r := range s {
		if start <= r && r < end { //if string s is between start && end
			index := int(r - start)
			result = index%numEnds != 0
		}
	}
	return result
}

func PrintBytes() bool {
	fmt.Println("print")
	return true
}
