package polymorphism

import "fmt"

type Shape interface {
	Render()
}

type Circle struct {
}

type Square struct {
}

func (c Circle) Render() {
	fmt.Println("This is a circle! ")

}

func (s Square) Render() {
	fmt.Println("This is a circle!")

}
