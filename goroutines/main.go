package main

import(
	"fmt"
	"time"
)

func insight() {
	fmt.Printf("Inside insight\nCurrent time:%v\n\n",time.Now())
	time.Sleep(2*time.Second)
	fmt.Printf("Inside insight now\nCurrent time:%v\n\n",time.Now())

	
}

func main()  {
	fmt.Printf("Inside main\nCurrent time:%v\n\n",time.Now())
	go insight()
	time.Sleep(1 * time.Second)
	fmt.Printf("Inside main now\nCurrent time:%v\n\n",time.Now())
	time.Sleep(2*time.Second)
	fmt.Printf("Inside main exiting now\nCurrent time:%v\n\n",time.Now())
}