// Write a function with one variadic parameter that finds the greatest
// number in a list of numbers.

package main

import "fmt"

func greatest(x ...int) int {
	g := x[0]
	for _,v := range x{
		if v>g{
			g=v
		}
	}
	return g
}

func main() {
	fmt.Println(greatest(91,543,678,6,79,46,468,474,6,7,78,843,3311,466,777,273))
}

// What's the value of this expression: (true && false) || (false && true) ||
// !(false && false)?

//True

