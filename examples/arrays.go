package main

import "fmt"

func main() {
	// Declare array of length 5
	var xs [5]int
	fmt.Println(xs) // => [0 0 0 0 0]

	// Addressing is pretty standard
	xs[2] = 3
	fmt.Println(xs[2]) // => 3

	// Can declare and initialise
	ys := [5]int{1, 2, 3, 4, 5}
	fmt.Println(ys) // => [1 2 3 4 5]

	// Can declare multi-dimensional
	var zs [2][2]int
	fmt.Println(zs) // => [[0 0] [0 0]]
}
