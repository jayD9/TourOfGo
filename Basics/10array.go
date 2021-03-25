package main

import "fmt"

var(
	primes = [6]int{1,3,5,7,11,13}
)

func main(){
	fmt.Printf("Arrays \n")

	// defining array
	var a [10]int
	a[0] = 10
	a[1] = 20

	fmt.Printf("Integer values of an arrays are %d %d \n",a[0],a[1])
	copyOfPrimes := primes
	fmt.Printf("Prime numbers are %v \n",primes)
	fmt.Println("-------------------- ")
	// slicing an array
	var s = primes[1:5]
	fmt.Println("primes: ",primes)
	fmt.Println("slice form 1 to 5 - s: ",s)

	s[3] = 99

	s1 := primes[2:6]
	fmt.Println("primes: ",primes)
	fmt.Println("slice form 1 to 5 - s1: ",s1)

	fmt.Println("copyOfPrimes: ",copyOfPrimes)
}