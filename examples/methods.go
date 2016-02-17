package main

import "fmt"

type Rectangle struct {
	width, height int
}

func (r *Rectangle) Area() int {
	return r.width * r.height
}

func (r *Rectangle) Perimeter() int {
	return (2 * r.width) + (2 * r.height)
}

func main() {
	r := Rectangle{width: 2, height: 3}
	fmt.Println(r.Area())      // => 6
	fmt.Println(r.Perimeter()) // => 10
}
