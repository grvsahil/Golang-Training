package main

import (
	"fmt"
	"time"
)

func server1(ch1 chan string){
	time.Sleep(2*time.Second)
	ch1<-"response from server1"
}

func server2(ch2 chan string){
	time.Sleep(4*time.Second)
	ch2<-"response from server2"
}


func main()  {
	ch1 := make(chan string)
	ch2 := make(chan string)
	go server1(ch1)
	go server2(ch2)

	//select blocks until response from first channel using cases
	select{
		case c1:=<-ch1:
			fmt.Println(c1)

		case c2:=<-ch2:
			fmt.Println(c2)

		// default:
		// 	fmt.Println("No response")
	}
}