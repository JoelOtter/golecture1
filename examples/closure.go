package main

import "fmt"

func sequence() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}

func main() {
	seq := sequence()
	fmt.Println(seq()) // => 1
	fmt.Println(seq()) // => 2

	seq = sequence()
	fmt.Println(seq()) // => 1
}
