package main

import (
	"context"
	"fmt"
	"sync"
	"time"

	"golang.org/x/sync/semaphore"
)

// Barrier structure to synchronize goroutines
type Barrier struct {
	mutex      sync.Mutex
	waitCount  int
	totalCount int
	sem        *semaphore.Weighted
}

// NewBarrier creates a new barrier for n goroutines
func NewBarrier(n int) *Barrier {
	return &Barrier{
		totalCount: n,
		sem:        semaphore.NewWeighted(1),
	}
}

// Wait implements the barrier synchronization
func (b *Barrier) Wait() {
	b.mutex.Lock()
	b.waitCount++
	if b.waitCount == b.totalCount {
		b.waitCount = 0
		b.sem.Release(1)
	}
	b.mutex.Unlock()

	_ = b.sem.Acquire(context.Background(), 1)
	b.sem.Release(1)
}

func doStuff(goNum int, wg *sync.WaitGroup, barrier *Barrier) bool {
	time.Sleep(time.Second)
	fmt.Println("Part A", goNum)

	// Wait at barrier until all goroutines complete Part A
	barrier.Wait()

	fmt.Println("Part B", goNum)
	wg.Done()
	return true
}

func main() {
	totalRoutines := 10
	var wg sync.WaitGroup
	wg.Add(totalRoutines)

	// Create barrier for synchronization
	barrier := NewBarrier(totalRoutines)

	// Launch goroutines
	for i := 0; i < totalRoutines; i++ {
		go doStuff(i, &wg, barrier)
	}

	wg.Wait() // wait for everyone to finish before exiting
}
