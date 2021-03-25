package main

import "fmt"

type Coord struct {
	X int
	Y int
}
// struct literals
var (
	p1 = Coord{
		1,
		1,
	} 
	p2 = Coord{2,2}
	p3 = Coord{X:3, Y:3}
	p4 = Coord{Y:4, X:4}
	p5 = Coord{}
	p6 = Coord{X: -1}
	p7 = Coord{Y: -1}
	p = &p1
)
func main(){
	fmt.Println("struct with pointers")

	//creating struct coordinate
	a := Coord{
		10,
		10,
	}

	// creating pointer for point a
	pOfA := &a

	// accessing value of through the pointer
	fmt.Println("a (X,Y): (",(*pOfA).X,",",pOfA.Y,")")


	//creating duplicate of struct
	 b := a

	 b.X = 5
	 b.Y = 5

	 fmt.Println("a (X,Y): (",a.X,a.Y,")")
	 fmt.Println("b (X,Y): (",b.X,b.Y,")")

	 var c = a

	 c.X = 100
	 c.Y = 100

	 fmt.Println("a (X,Y): (",a.X,a.Y,")")
	 fmt.Println("c (X,Y): (",c.X,c.Y,")")



	 //printing variables
	 fmt.Println(p1, p2, p3, p4, p5, p6, p7, p)



}