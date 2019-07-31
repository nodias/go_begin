package structTest

import (
	"fmt"
	"time"
)

type A struct {
	*time.Time
}

type B struct {
	time.Time
}

func a(t *time.Time) *A {
	return &A{t}
}

func b(t time.Time) *B {
	return &B{t}
}

func ExampleAB() {
	c := time.Now()
	println(&c)
	bb := b(c)
	println(&bb.Time)
	aa := a(&c)
	println(aa.Time)

	fmt.Println(c)
	fmt.Println(bb.Time)
	fmt.Println(aa.Time)

	c = time.Now().Add(time.Hour)

	fmt.Println(c)
	fmt.Println(bb.Time)
	fmt.Println(aa.Time)
	//Output:
}
