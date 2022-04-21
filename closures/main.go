package main

import(
	"fmt"
)

// func main()  {
// 	v := func() string{
// 		return "Hello World"
// 		// fmt.Println("Hello World")
// 	}
// 	// v()
// 	fmt.Println(v())
// 	fmt.Printf("%T\n",v)	
// }

// func getHello() func() string {
// 	return func() string{
// 		return "Hello World"
// 	}
// }

// func main()  {
// 	v := getHello()
// 	fmt.Println(v())
// 	fmt.Println(v())

// }

// func getSno() func() int {
// 	var s int
// 	return func() int {
// 		s+=1
// 		return s
// 	}
// }

// func main()  {
// 	sno := getSno()
// 	fmt.Println(sno())
// 	fmt.Println(sno())
// }

func printNum(x int,getNegative func(x int) int)  {
	fmt.Println(getNegative(x))
}

func getNegative(x int) int {
	return x*-1
}

func main()  {
	printNum(45,getNegative)
}