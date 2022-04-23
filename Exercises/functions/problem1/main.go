// Write a function which takes an integer. The function will have two
// returns. The first return should be the argument divided by 2. The
// second return should be a bool that let’s us know whether or not the
// argument was even. For example:
// a. half(1) returns (0, false)
// b. half(2) returns (1, true)

package main

import "fmt"

func half(x int) (y float64,z bool) {
	y = float64(x)/2.0
	z=false
	if x%2 == 0{
		z = true
	}
	return
}

func main() {
	var num int
	fmt.Printf("Enter number: ")
	fmt.Scanln(&num)
	fmt.Println(half(num));
}