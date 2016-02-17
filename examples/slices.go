package main

import "fmt"

func main() {
	// We use 'make' to create slices
	words := make([]string, 3)
	// Initially zero-valued
	fmt.Println(words) // => [  ]

	// Initialise explicitly
	dogs := []string{"Pug", "Corgi"}
	dogs = append(dogs, "The Hound")
	fmt.Println(dogs) // => [Pug Corgi The Hound]

	// Can copy slices too
	alsoDogs := make([]string, len(dogs))
	copy(alsoDogs, dogs)
	fmt.Println(alsoDogs) // => [Pug Corgi The Hound]

	fmt.Println(dogs[1:]) // => [Corgi The Hound]
	fmt.Println(dogs[:1]) // => [Pug]
}
