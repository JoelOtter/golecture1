package main

import "fmt"

func main() {
	// We use 'make' to create a new map
	prices := make(map[string]float32)
	prices["banana"] = 0.48
	fmt.Println(prices) // => map[banana:0.48]

	// Deleting from a map
	prices["apple"] = 0.29
	delete(prices, "banana")
	fmt.Println(prices) // => map[apple:0.29]

	// Explicit initialisation
	ages := map[string]int{"Bob": 22, "Alice": 23}

	// We can check for membership...
	_, bobExists := ages["Bob"]
	fmt.Println(bobExists) // => true
}
