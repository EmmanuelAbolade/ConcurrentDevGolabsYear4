package main

import (
	"fmt"
	"sync"
	"time"
)

// Semaphore struct with channel - methods for acquire and release
type Semaphore struct {
	permits chan struct{}
}

// NewSemaphore initializes a new semaphore with given capacity
func NewSemaphore(maxConcurrent int) *Semaphore {
	return &Semaphore{
		permits: make(chan struct{}, maxConcurrent),
	}
}

// Acquire gets a permit - blocks if none available
func (s *Semaphore) Acquire() {
	s.permits <- struct{}{}
}

// Release returns a permit
func (s *Semaphore) Release() {
	<-s.permits
}

func main() {
	maxGoroutines := 5
	//semaphore := make(chan struct{}, maxGoroutines)
	sem := NewSemaphore(maxGoroutines)
	var wg sync.WaitGroup
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func(taskID int) {
			//go func(i int) {
			defer wg.Done()

			// Acquire semaphore permit
			sem.Acquire()
			// Always release when done
			defer sem.Release()

			// Simulate a task
			fmt.Printf("Running task %d\n", taskID)
			time.Sleep(2 * time.Second)
			fmt.Printf("Task %d completed\n", taskID)
		}(i)
	}
	wg.Wait()
	fmt.Println("All tasks completed!")
}
