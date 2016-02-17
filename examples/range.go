package main

import "fmt"

func main() {
	dogs := []string{"Pug", "Corgi"}
	for i, dog := range dogs {
		fmt.Printf("%v: %v\n", i, dog) // %v is default format
	}

	ages := map[string]int{"Alice": 21, "Bob": 22}
	for person, age := range ages {
		fmt.Printf("%v: %v\n", person, age)
	}
}
