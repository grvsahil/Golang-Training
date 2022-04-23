// When the same part of code called critical section is accessed by multiple goroutines at the same time
// then we get undesirable values so we want that only one goroutine should access the critical section

// At this time mutex come in role.It is used to lock and unlock code block so that if one goroutine is
// accessing it another one won't be able to access it.

package main

import (
	"fmt"
	"sync"
)

var x int

func increment(wg *sync.WaitGroup,mut *sync.Mutex)  {
	mut.Lock()
	x+=1
	mut.Unlock()
	wg.Done()
}

func main()  {
	var wg sync.WaitGroup
	var mut sync.Mutex
	for i:=0;i<100;i++{
		wg.Add(1)
		go increment(&wg,&mut)
	}
	wg.Wait()
	fmt.Println("final value of x is ",x)
}