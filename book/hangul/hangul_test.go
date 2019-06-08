package hangul

import "fmt"

func ExampleHasConsonantSuffix() {
	fmt.Println(HasConsonantSuffix("가나다"))
	fmt.Println(HasConsonantSuffix("가나다."))
	fmt.Println(HasConsonantSuffix("간s나단"))
	// Output:
	// false
	// false
	// true
}
