package main

import (
	"fmt"
	"math/rand/v2"
	"sync"
	"time"
)

//No Global variables shared between functions --A good IDEA

func WorkWithRendezvous(wg *sync.WaitGroup, arrivalChan chan bool, departureChan chan bool, Num int) bool {
	defer wg.Done()
	// Random sleep (0-4 seconds)
	var X time.Duration
	X = time.Duration(rand.IntN(5))
	time.Sleep(X * time.Second) //wait random time amount
	fmt.Println("Part A", Num)
	//Rendezvous here Two-phase barrier
	arrivalChan <- true // Signal: "I've arrived"
	<-departureChan     // Wait: "Everyone's here, proceed!"

	fmt.Println("Part B", Num)

	return true
}

// Coordinator manages the rendezvous barrier
func coordinator(arrivalChan chan bool, departureChan chan bool, threadCount int) {
	// Wait for all to arrive
	for i := 0; i < threadCount; i++ {
		<-arrivalChan
	}

	// Release all to proceed
	for i := 0; i < threadCount; i++ {
		departureChan <- true
	}

	close(arrivalChan)
	close(departureChan)
}

func main() {
	var wg sync.WaitGroup
	threadCount := 5
	//using barrier := make(chan bool)
	arrivalChan := make(chan bool)
	departureChan := make(chan bool)

	// Start coordinator goroutine
	go coordinator(arrivalChan, departureChan, threadCount)
	wg.Add(threadCount)
	for N := range threadCount {
		go WorkWithRendezvous(&wg, arrivalChan, departureChan, N)
	}
	wg.Wait() //wait here until everyone (10 go routines) is done

}
