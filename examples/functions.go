package main

import "fmt"

func sum1(a int, b int) int {
	return a + b
}

func sum2(a, b int) int {
	return a + b
}

func sum(nums ...int) int {
	result := 0
	for _, num := range nums {
		result += num
	}
	return result
}

func divideWithRemainder(a, b int) (int, int) {
	return a / b, a % b
}

func main() {
	div, rem := divideWithRemainder(10, 3)
	fmt.Println(div, rem) // => 3 1
}
