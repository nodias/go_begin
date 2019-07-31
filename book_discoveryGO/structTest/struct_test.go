package structTest

// import (
// 	"fmt"
// 	"time"
// )

// type D struct {
// 	dd string
// }

// type E struct {
// 	*time.Time
// }

// type F struct {
// 	time.Time
// }

// func ExampleStruct() {
// 	a := "a"            //literal
// 	b := [3]string{"b"} //&
// 	c := []string{"c"}  //p
// 	d := D{a}           //&
// 	t := time.Now()     //&
// 	e := E{&t}          //p
// 	f := F{t}           //&

// 	println("a:", a)
// 	println("b:", "err (b)")
// 	println("c:", c)
// 	println("d:", "err (d)")
// 	println("t:", "err (t)")
// 	println("e:", e.Time)    //e.Time은 구조체의 주소를 가짐
// 	println("f:", "err (f)") //f.Time은 구조체를 복사한값을 가짐

// 	println()
// 	println("&a:", &a)
// 	println("&b:", &b)
// 	println("&c:", &c)
// 	println("&d:", &d)
// 	println("&t:", &t)
// 	println("&e:", &e.Time)
// 	println("&f:", &f.Time)

// 	//Output:
// 	//
// }

// func ExampleA() {
// 	d := D{"dd"}
// 	fmt.Println(d)
// 	B(d)
// 	fmt.Println(d)
// 	C(&d)
// 	fmt.Println(d)
// 	//Output:
// }

// func B(d D) {
// 	d.dd = "dd=B"
// 	fmt.Println("B : ", d.dd)
// }

// func C(d *D) {
// 	d.dd = "dd=C"
// 	fmt.Println("C : ", d.dd)
// }
