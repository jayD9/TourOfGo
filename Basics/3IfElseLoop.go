package main

import (
	"fmt"
)

func main() {

	a := 5
	b := 6

	// if loop
	if a < b {
		fmt.Println("a is smaller than b")
	}

	// if-else loop
	if a < b {
		fmt.Println("a is smaller than b")
	} else {
		fmt.Println("b is samller than a")
	}

	a = 10
	b = -100

	// if loop with short statment
	if sum := a + b; sum > 0 {
		fmt.Println("Sum is positive: ", sum)
	}
	// we cant access sum variable outside of if loop

	// if-else loop with short statment
	if sum := a + b; sum > 0 {
		fmt.Println("Sum is positive: ", sum)
	} else {
		fmt.Println("Sum is negative: ", sum)
	}
	// we cant access sum variable outside of if-else loop
}
