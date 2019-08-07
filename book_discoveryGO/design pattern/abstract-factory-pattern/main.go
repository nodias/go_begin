package main

import "fmt"

type Button interface {
	Paint()
	OnClick()
}

type Label interface {
	Paint()
}

type WinButton struct{}

func (w WinButton)Paint() {
	fmt.Println("win button paint")
}

func (w WinButton)OnClick() {
	fmt.Println("win button click")
}

type MacButton struct{}

func (w MacButton)Paint() {
	fmt.Println("mac button paint")
}

func (w MacButton)OnClick() {
	fmt.Println("mac button click")
}

type WinLabel struct{}

func (w WinLabel)Paint() {
	fmt.Println("win label paint")
}

type MacLabel struct{}

func (w MacLabel)Paint() {
	fmt.Println("mac label paint")
}

type UIFactory interface {
	CreateButton() Button
	CreateLabel() Label
}

type WinFactory struct {

}

func (w WinFactory) CreateButton() Button {
	return WinButton{}
}

func (w WinFactory) CreateLabel() Label {
	return WinLabel{}
}

type MacFactory struct {

}

func (m MacFactory) CreateButton() Button {
	return MacButton{}
}

func (m MacFactory) CreateLabel() Label {
	return MacLabel{}
}

func CreateFactory(os string) UIFactory {
	if os == "window" {
		return WinFactory{}
	} else if os =="mac" {
		return MacFactory{}
	} else {
		panic("Unsupported OS")
	}
}

func Run(f UIFactory){
	button := f.CreateButton()
	button.Paint()
	button.OnClick()
	label := f.CreateLabel()
	label.Paint()
}

func main() {
	winfac := CreateFactory("window")
	Run(winfac)
	macfac := CreateFactory("mac")
	Run(macfac)
}


