package main

import "fmt"

type Coordintae struct {
	X int
	Y int
}

func main(){
	fmt.Println("basic struct")

	//defining struct
	var co1 = Coordintae {5,3}

	//defining struct. "," is ncessary after the last parameter
	var co2 = Coordintae{
		5,
		9,
	}


	fmt.Println(co1)
	fmt.Println(co2.Y)
	fmt.Println(co1.X)

}