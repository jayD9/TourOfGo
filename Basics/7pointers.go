package main

import "fmt"

func main() {
	fmt.Println("pointers")

	//declaring pointers
	var p *int;

	//default value of pointer
	fmt.Println("default value of pointer: ",p)

	i := 33
	fmt.Println("i: ",i)

	//assigning value to the pointer
	p = &i

	//afeter assigning value of the pointer
	fmt.Println("after assigning value of pinter: ",p)	

	//fetching value through the pointer 
	fmt.Println(*p)

	//changing value through the pointer 
	*p = 22
	fmt.Println("i: ",i)

}
