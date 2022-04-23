package main

import (
	"fmt"
	"sync"
	// "time"
)

func goroutine(i int,wg *sync.WaitGroup)  {
	fmt.Println("started goroutine ",i)
	// time.Sleep(time.Millisecond)
	fmt.Println("ended goroutine ",i)
	wg.Done()
}

func main()  {
	var wg sync.WaitGroup
	for i:=0 ; i<3 ;i++{
		wg.Add(1)
		go goroutine(i,&wg)
	}
	wg.Wait()
}