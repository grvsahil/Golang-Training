package main

import "fmt"

func main()  {
	fmt.Println("Even numbers b/w 0 to 100:")
	for i:=2 ; i<=100 ; i++{
		if i%2==0{
			fmt.Println(i)
		}
	}
}