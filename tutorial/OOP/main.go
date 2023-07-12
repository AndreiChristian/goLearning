package main

import (
	"fmt"
	"oop/composition"
	"oop/polymorphism"
)

type Line struct{}

func (l Line) Render() {
	fmt.Println("This is a line")
}

type Triangle struct{}

func (t Triangle) Render() {
	fmt.Println("This is a triangle")
}

func main() {

	var square polymorphism.Shape = polymorphism.Square{}

	var circle polymorphism.Shape = polymorphism.Circle{}

	var line polymorphism.Shape = Line{}

	var triangle polymorphism.Shape = Triangle{}

	square.Render()

	circle.Render()

	line.Render()

	triangle.Render()

	car := composition.NewCar()

	fmt.Println(car)

}
