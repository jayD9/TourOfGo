package main

import (
	"fmt"
	"time"
)

func main() {

	// switch statement
	v := "five"
	switch v {
	case "one":
		fmt.Println(1)
	case "two":
		fmt.Println(2)
	case "three":
		fmt.Println(3)
	default:
		fmt.Println("> 3")
	}

	// switch wioth short statement
	switch v := getValue(); v {
	case "one":
		fmt.Println(1)
	case "two":
		fmt.Println(2)
	case "three":
		fmt.Println(3)
	default:
		fmt.Println("> 3")
	}

	t := time.Now()

	// switch with condition
	switch {
	case t.Hour() < 12:
		fmt.Println("morning")
	case t.Hour() < 17:
		fmt.Println("afternoon")
	default:
		fmt.Println("evening")
	}

}

func getValue() string {
	return "three"
}
