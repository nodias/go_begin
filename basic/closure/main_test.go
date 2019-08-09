package closure_test

import (
	"fmt"
	"github.com/nodias/go_begin/basic/closure"
)

//func Example_A() {
//	a := closure.GetA()
//	fmt.Println(a(4))
//	fmt.Println(a(11))
//
//	//Output:
//	//
//}
//
////javascript와 달리 람다의 뒤에 (i)처럼 실행 인자를 미리 넣어주면,
////그것은 반환되기 전 처리되기 때문에 이미 int type이다.
//func Example_B() {
//	var a []int
//	for i := 0; i < 10; i++ {
//		a = append(a, func(id int) int {
//			return id
//		}(i))
//
//	}
//	for _, f := range a {
//		fmt.Println(f)
//	}
//
//	//Output:
//	//
//}

//func Example_C() {
//	var a []func() int
//	for i := 0; i < 10; i++ {
//		a = append(a, func() int {
//			return i
//		})
//
//	}
//	for _, f := range a {
//		fmt.Println(f())
//	}
//
//	//Output:
//	//
//}

func Example_D() {
	d := closure.B(true)
	arr1 := d("apple")
	for _, apple := range arr1{
		fmt.Println(apple)
	}
	arr2 := d("banana")
	for _, banana:= range arr2{
		fmt.Println(banana)
	}

	//Expected Output:
	//apple : 0
	//apple : 1
	//apple : 2
	//apple : 3
	//apple : 4
	//banana : 0
	//banana : 1
	//banana : 2
	//banana : 3
	//banana : 4

	//Output:
	//apple : 0
	//apple : 1
	//apple : 2
	//apple : 3
	//apple : 4
	//apple : 0
	//apple : 1
	//apple : 2
	//apple : 3
	//apple : 4
	//banana : 0
	//banana : 1
	//banana : 2
	//banana : 3
}


