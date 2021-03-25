package main

import (
	"fmt"
)

func main() {
	a()
}

func a() {
	fmt.Println("a begins")
	defer b()
	fmt.Println("a ends")
}

func b() {
	fmt.Println("In b")
}
