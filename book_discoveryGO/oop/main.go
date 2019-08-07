package oop

import (
	"fmt"
	"reflect"
)

type Shape interface {
	Area() float32
}

type Rectangle struct {
	Width, Height float32
}

// IsA 상속이 되려면 이 메서드가 Rectangle이 들어가는 곳에 들어갈 수 있어야한다.
// Golang 에서는 이를 interface 구현을 이용한다.
func (r Rectangle) Area() float32 {
	return r.Width * r.Height
}

func TotalArea(s []Shape) float32 {
	var totalArea float32
	for _, sh := range s {
		totalArea += sh.Area()
	}
	return totalArea
}

// HasA
type RectangleInherit struct {
	Rectangle
}

func (r RectangleInherit) Circum() float32 {
	return 2 * (r.Width * r.Height)
}

//// Overriding, 주석 풀면 결과는 420 + 210 = 630이 됨
//func (r RectangleInherit) Area() float32 {
//	return r.Width * r.Height * 2
//}

// Constructor
func NewRectangleInherit(w, h float32) *RectangleInherit {
	return &RectangleInherit{Rectangle{
		Width:  w,
		Height: h,
	}}
}

func main() {
	res := TotalArea(
		[]Shape{
			//SubType, RectangleInherit이 Area()를 가진 것은 아니지만 Shape구현체로 사용가능
			NewRectangleInherit(14, 15),
			Rectangle{
				Width:  14,
				Height: 15,
			},
		},
	)
	fmt.Println(res)
	// 210 + 210 = 420


	//reflect implements confirm
	impl := reflect.TypeOf(NewRectangleInherit(1,2)).Implements(reflect.TypeOf((*Shape)(nil)).Elem())
	fmt.Println(impl)
	//true

	field, ok := reflect.TypeOf(RectangleInherit{}).FieldByName("Rectangle")
	emb := ok && field.Anonymous && field.Type == reflect.TypeOf(Rectangle{})
	fmt.Println(emb)
	//true
}
