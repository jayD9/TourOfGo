package main

import (
	"fmt"
)

func main() {
	fmt.Println("main begins")

	for i := 1; i <= 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("main ends")
}
