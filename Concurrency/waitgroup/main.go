package main

import(
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func foo()  {
	for i:=0;i<10;i++{
		fmt.Println(i)
		time.Sleep(2*time.Millisecond)
	}
	wg.Done()
}

func boo()  {
	for i:=0;i<10;i++{
		fmt.Println(i)
		time.Sleep(3*time.Millisecond)
	}
	wg.Done()
}

// we set 2 to wg.add and whenever wg.done is executed then 1 is taken away from wg.add
// wg.wait waits for wg.add to become zero

func main()  {
	wg.Add(2)
	go foo()
	go boo()
	wg.Wait()
}