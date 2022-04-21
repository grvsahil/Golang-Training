package main

import(
	"fmt"
)

type shape interface{
	area() int
}

type square struct{
	side int
}

type rectangle struct{
	length int
	breadth int
}

func (s square) area() int {
	fmt.Println("Square area:")
	return (s.side * s.side)
}

func (r rectangle) area() int {
	fmt.Println("Rectangle area:")
	return (r.length*r.breadth)
}

// func printArea(s shape){
// 	fmt.Println(s.area())
// }

func main()  {
	sq := square{25}
	rec := rectangle{25,25}
	shapes := []shape{sq,rec}
	
	for _,val := range shapes{
		fmt.Println(val.area())
	}

	//or

	// for _,val := range shapes{
	// 	printArea(val)
	// }
}