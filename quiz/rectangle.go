package main

import (
	"fmt"
)

type Rect interface{}

type rect struct {
	width, height int
}

func (r *rect) NewRect(w, h int) Rect {
	return rect{width: w, height: h}
}

func (r *rect) area() int {
	return r.width * r.height
}

func (r *rect) perim() int {
	return r.width + r.height
}

func main() {
	fmt.Println("Welcome to rectngle OOP")

	r := rect{width: 10, height: 5}
	fmt.Printf("Area: %d  Perimeter: %d\n", r.area(), r.perim())

	s := NewRect(1, 2)
	fmt.Printf("Area: %d  Perimeter: %d\n", s.area(), s.perim())

}
