package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float32
	Perimeter() float32
}

type Rectangle struct {
	width, height float32
}

func (r *Rectangle) Area() float32 {
	return r.width * r.height
}

func (r *Rectangle) Perimeter() float32 {
	return (2 * r.width) + (2 * r.height)
}

type Circle struct {
	radius float32
}

func (c *Circle) Area() float32 {
	return math.Pi * c.radius * c.radius
}

func (c *Circle) Perimeter() float32 {
	return 2 * math.Pi * c.radius
}

func main() {
	shapes := []Shape{&Rectangle{2, 3}, &Circle{5}}
	for _, shape := range shapes {
		fmt.Println(shape.Area())
	}
}
