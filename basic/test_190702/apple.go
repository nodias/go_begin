package main

type Apple struct {
	Color string
	Sugar int
}

var apple = Apple{
	"red",
	1,
}

func NewApple() *Apple {
	println(&apple)
	return &apple
}
