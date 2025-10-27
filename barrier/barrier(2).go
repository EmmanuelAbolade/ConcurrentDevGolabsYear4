//Barrier.go Template Code
//Copyright (C) 2025 Emmanuel Abolade

// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

//--------------------------------------------
// Author: Emmanuel Abolade (c00288657@setu.ie)
// Created on 19/10/2024
// Modified by:
// Issues:
//--------------------------------------------

package main

import (
	"context"
	"fmt"
	"sync"
	"time"

	"golang.org/x/sync/semaphore"
)

// Barrier implementation using Mutex and Semaphores
func doSomeStuff(goNum int, wg *sync.WaitGroup, mutex *sync.Mutex, counter *int,
	sem *semaphore.Weighted, ctx context.Context, totalRoutines int) bool {
	defer wg.Done()

	time.Sleep(time.Second)
	fmt.Println("Part A", goNum)

	//Barrier: we wait here until everyone has completed part A
	mutex.Lock()
	*counter++
	if *counter == totalRoutines {
		// Last one to arrive - release everyone
		sem.Release(int64(totalRoutines))
	}
	mutex.Unlock()

	// Wait at barrier
	if err := sem.Acquire(ctx, 1); err != nil {
		fmt.Printf("Error acquiring semaphore: %v\n", err)
		return false
	}
	fmt.Println("Part B", goNum)

	return true
}

func main() {
	totalRoutines := 10
	var wg sync.WaitGroup
	var mutex sync.Mutex
	var counter = 0

	//we will need some of these
	ctx := context.TODO()
	//var theLock sync.Mutex
	sem := semaphore.NewWeighted(int64(totalRoutines))

	// Acquire all semaphore slots initially (barrier is closed)
	err := sem.Acquire(ctx, int64(totalRoutines))
	if err != nil {
		return
	}

	//theLock.Lock()
	//sem.Acquire(ctx, 1)
	wg.Add(totalRoutines)

	for i := range totalRoutines { //create the go Routines here
		go doSomeStuff(i, &wg, &mutex, &counter, sem, ctx, totalRoutines)
	}
	//sem.Release(1)
	//theLock.Unlock()

	wg.Wait() //wait for everyone to finish before exiting
}
