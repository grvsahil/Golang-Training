package main

import "fmt"

type gaurav int

var x gaurav
var y gaurav

func main() {
	fmt.Println(x)
	fmt.Printf("%T\n",x)
	x = 42
	fmt.Println(x)
	y := gaurav(x)
	fmt.Println(y)
}