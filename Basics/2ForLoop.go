package main

import (
	"fmt"
)

func main() {

	// for loop
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
	fmt.Println("------------------")

	// initailization and increment statment is optional
	var j = 1
	for j <= 128 {
		fmt.Println(j)
		j += j
	}
	fmt.Println("------------------")

	// for can also work as while
	j = 1
	for j <= 128 {
		fmt.Println(j)
		j += j
	}
	fmt.Println("------------------")

	// infinite loop
	//for {
	//}

}
