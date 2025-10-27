package main

import (
	"fmt"
	"sync"
	"time"
)

//No Global variables shared between functions --A good IDEA

func main() {
	var wg sync.WaitGroup
	// Unbuffered channel for synchronization
	barrier := make(chan bool)

	doStuffOne := func() {
		defer wg.Done()
		fmt.Println("StuffOne - Part A")
		//wait here (Send signal - will block until doStuffTwo receives)
		barrier <- true
		fmt.Println("StuffOne - PartB")

	}
	doStuffTwo := func() {
		defer wg.Done()

		// Simulate some work
		time.Sleep(time.Second * 5)
		fmt.Println("StuffTwo - Part A")

		//wait here (Receive signal - unblocks doStuffOne_
		<-barrier
		fmt.Println("StuffTwo - PartB")

	}
	wg.Add(2)
	go doStuffOne()
	go doStuffTwo()
	wg.Wait()      //wait here until everyone (10 go routines) is done
	close(barrier) // Clean up the channel
	fmt.Println("All done!")

}
