package main

import (
	"fmt"
	"math"
)

// define interface
type Shape interface {
	// define methods to implement
	area() float64
}

// define structs to work with interface
type Circle struct {
	r float64
}

type Rectangle struct {
	x, y float64
}

// implement interface methods for all structs
func (c Circle) area() float64 {
	return math.Pi * c.r * c.r
}

func (r Rectangle) area() float64 {
	return r.x * r.y
}

// implement interface
func getArea(s Shape) float64 {
	return s.area()
}

func main() {
	circle := Circle{5}
	rectangle := Rectangle{10, 5}

	fmt.Println(getArea(circle))
	fmt.Println(getArea(rectangle))
}
