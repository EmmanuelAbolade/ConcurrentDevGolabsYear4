package main

import (
	"fmt"
	"sync"
)

//Global variables shared between functions --A BAD IDEA

func adds(n int, total *int64, theLock *sync.Mutex, wg *sync.WaitGroup) {

	wg.Done()
	for i := 0; i < n; i++ {
		theLock.Lock()
		*total++
		theLock.Unlock()
	}

}

func main() {

	//theLock will be passed by reference between go routines
	//better than using a global variable
	var theLock sync.Mutex

	//total = 0
	//the waitgroup is used as a barrier
	var wg sync.WaitGroup
	var total int64

	// init it to number of go routines
	wg.Add(10)

	//for loop using range option
	for i := range 10 {
		//starting
		fmt.Println("Starting goroutine", i)
		go adds(1000, &total, &theLock, &wg)
	}
	wg.Wait() //wait here until everyone (10 go routines) is done
	fmt.Println("Final total:", total)
}
