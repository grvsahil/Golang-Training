package main

import (
	"fmt"
	// "time"
	//"time"
)

// // receiving data from channel

// func add(x int,y int,chnl chan int){
// 	sum := x+y
// 	chnl <- sum

// }

// func main()  {
// 	chnl := make(chan int)
// 	go add(5,6,chnl)
// 	dataReceived := <- chnl
// 	fmt.Println(dataReceived)

// }

// //Receiving struct through channel

// type rec struct{
// 	length int
// 	breadth int
// }

// func area(x chan rec)  {
// 	y := <- x
// 	fmt.Println("Received data from channel now printing result")
// 	fmt.Printf("Area is: %d\n",y.length*y.breadth)
// }

// func main()  {
// 	chnl := make(chan rec)
// 	m := rec{5,4}
// 	fmt.Println("Sending data through channel now")
// 	go area(chnl)
// 	chnl <- m

// }

// //Unidirectional channel

// func main()  {
// 	chnl := make(chan<- bool)
// 	go func(chnl chan<- bool) {
// 		chnl <- false
// 	}(chnl)
// 	fmt.Println("Hello")
// }

// //Bidirectional channel can be converted to unidirectional channel

// // retNeg only sends data and never receives from the channel so channel can be converted to
// // send only channel in retNeg's scope but it is bidirectional in main scope
// func retNeg(chnl chan<- int,x int){
// 	chnl <- x*-1
// }

// func main()  {
// 	chnl := make(chan int)
// 	go retNeg(chnl, 56)

// 	fmt.Printf("Negative value is %d\n",<-chnl)

// }

// //Closing a Channel and using it in for loops

// func sendData(chnl chan<- int)  {
// 	for i:=0;i<10;i++{
// 		chnl <- i+1
// 	}
// 	close(chnl)
// }

// func main()  {
// 	chnl := make(chan int)
// 	go sendData(chnl)
// 	for {
// 		v,ok := <-chnl
// 		if ok==true{
// 			fmt.Println(v)
// 		}else{
// 			fmt.Println("Channel closed")
// 			break
// 		}
// 	}

// }

//Buffered Channels

//The send and receive in Channels having buffer are not blocking until the buffer get full
//The send is blocking only when the buffer is full
//The receive is blocking only when bufferis empty

// func getNeg(buffChnl chan int)  {
// 	for v:= range buffChnl{
// 		fmt.Println(-1*v)
// 	}
// }

// func main()  {
// 	buffChnl := make(chan int,2)
// 	go getNeg(buffChnl)
// 	buffChnl<-5
// 	buffChnl<-6
// 	buffChnl<-7
// 	buffChnl<-8
// 	time.Sleep(time.Second)

// }

func sendData(chnl chan int){
	for i:=0;i<10;i++{
		chnl<-i
	}
}

func main()  {
	chnl := make(chan int,2)
	go sendData(chnl)
	for i:=0;i<10;i++{
		// time.Sleep(time.Second)
		a:=<-chnl
		fmt.Println(a)
	}
}