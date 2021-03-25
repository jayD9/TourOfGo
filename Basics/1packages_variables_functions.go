package main

// main package will always run firsrt
import (
	"fmt"
)

// variables at package level
var i int

// main function
func main() {
	// variables at function level
	var p bool
	var q string

	// varibles with initializer and without type
	var name, age, status = "Jay", 27, false

	fmt.Println("Hello, playground")
	fmt.Println("addition: ", add(3, 5))
	fmt.Printf("Subtraction %d \\n", sub(5, 3))
	a, b := swap("Hello", "world")
	fmt.Println(a, b)
	x, y := split(99)
	fmt.Println(x, y)

	// default values of variables
	fmt.Println("i = ", i)
	fmt.Println("p = ", p)
	fmt.Println("q = ", q)

	// variables defined without data type
	fmt.Printf("Type of name is %T \\n", name)
	fmt.Printf("Type of age is %T \\n", age)
	fmt.Printf("Type of status is %T \\n", status)

	// Type convesion
	fmt.Printf("convert type form int to %T", convertIntToFloat(5))

	// declaring constants
	const newConstant = 8
	fmt.Println("new constant is ", newConstant)
}

// function with two parameters
func add(x int, y int) int {
	return x + y
}

// function with multiple consicutive same type paramtere
func sub(x, y int) int {
	return x - y
}

// function returning multiple values
func swap(x, y string) (string, string) {
	return y, x
}

// function with named return value akso called as naked return
func split(sum int) (x, y int) {
	x = sum % 10
	y = sum - x
	return
}

func convertIntToFloat(x int) float32 {
	return float32(x)
}
