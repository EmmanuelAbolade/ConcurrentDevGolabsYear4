package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

//No Global variables shared between functions --A Good IDEA

func addsAtomic(n int, total *atomic.Int64, wg *sync.WaitGroup) {

	defer wg.Done() //let waitgroup know we have finished
	for i := 0; i < n; i++ {
		total.Add(1)
	}

}

func main() {

	var total atomic.Int64
	var wg sync.WaitGroup

	// Initializing waitgroup BEFORE starting goroutines to prevents race conditions
	wg.Add(10)

	//for loop using range option (launch 10 goroutines)
	for i := range 10 {
		//the waitgroup is used as a barrier
		// init it to number of go routines

		fmt.Println("go Routine ", i)
		go addsAtomic(1000, &total, &wg)
	}
	wg.Wait() //wait here until everyone (10 go routines) is done
	fmt.Println("Final total:", total.Load())

}
