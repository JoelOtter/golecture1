package main

import (
	"errors"
	"fmt"
)

func divideTenBy(n int) (int, error) {
	if n == 0 {
		return -1, errors.New("Can't divide by zero")
	}
	return 10 / n, nil
}

func main() {
	if n, err := divideTenBy(0); err != nil {
		fmt.Println("Got error:", err)
	} else {
		fmt.Println("Got result:", n)
	}
}
