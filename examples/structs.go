package main

import "fmt"

type Person struct {
	name string
	age  int
}

func main() {
	bob := Person{"Bob", 22}
	alice := Person{name: "Alice", age: 21}
	fmt.Println(bob.age)    // => 22
	fmt.Println(alice.name) // => Alice
}
