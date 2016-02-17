package main

import "fmt"

func main() {
	// Uninitialised variables are 'zero-valued'.
	var emptyVariable int
	fmt.Println(emptyVariable) // => 0
}
