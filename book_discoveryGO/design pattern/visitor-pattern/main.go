package main

import "fmt"

func main(){
	car := Car {
		Wheel("front left"),
		Wheel("front right"),
		Wheel("back left"),
		Wheel("back left"),
		Body{},
		Engine{},
	}

	car.Accept(CarElementPrintVisitor{})
	car.Accept(CarElementDoVisitor{})
}

//
type CarElementVisitor interface {
	VisitWheel(wheel Wheel)
	VisitEngine(engine Engine)
	VisitBody(body Body)
	VisitCar(car Car)
}

type Acceptor interface {
	Accept(visitor CarElementVisitor)
}

// CarElement들을 모두 Acceptor로 만들어 준다.
type Wheel string

func (w Wheel) Name() string{
	return string(w)
}

func (w Wheel) Accept(visitor CarElementVisitor){
	visitor.VisitWheel(w)
}

type Engine struct{}

func (e Engine) Accept(visitor CarElementVisitor) {
	visitor.VisitEngine(e)
}

type Body struct{}

func (b Body) Accept(visitor CarElementVisitor) {
	visitor.VisitBody(b)
}

type Car []Acceptor

func (c Car) Accept(visitor CarElementVisitor) {
	for _, e := range c {
		e.Accept(visitor)
	}
	visitor.VisitCar(c)
}

type CarElementPrintVisitor struct {}

func (c CarElementPrintVisitor) VisitWheel(wheel Wheel) {
	fmt.Println("Visiting" + wheel.Name() + " wheel")
}

func (c CarElementPrintVisitor) VisitEngine(engine Engine) {
	panic("implement me")
}

func (c CarElementPrintVisitor) VisitBody(body Body) {
	panic("implement me")
}

func (c CarElementPrintVisitor) VisitCar(car Car) {
	panic("implement me")
}

type CarElementDoVisitor struct {}

func (c CarElementDoVisitor) VisitWheel(wheel Wheel) {
	fmt.Println("Kicking my " + wheel.Name() + "wheel")
}

func (c CarElementDoVisitor) VisitEngine(engine Engine) {
	panic("implement me")
}

func (c CarElementDoVisitor) VisitBody(body Body) {
	panic("implement me")
}

func (c CarElementDoVisitor) VisitCar(car Car) {
	panic("implement me")
}




