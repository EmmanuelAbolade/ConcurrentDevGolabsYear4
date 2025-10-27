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
// Author: Emmanuel abolade (c00288657@setu.ie)
// Created on 20/10/2025
// Modified by:
// Description:
// A simple barrier implemented using mutex and unbuffered channel
// Issues:
// None I hope
//1. Change mutex to atomic variable
//2. Make it a reusable barrier
//--------------------------------------------

package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

// Place a barrier in this function --use Mutex's and Semaphores
func doSomething(goNum int, arrived *int32, max int, wg *sync.WaitGroup, theChan chan bool) {
	defer wg.Done()
	time.Sleep(time.Second)
	fmt.Println("Part A", goNum)
	// BARRIER: Wait here until everyone has completed Part A
	// Use atomic instead of mutex for counting
	count := atomic.AddInt32(arrived, 1)

	if count == int32(max) {
		// Last to arrive - signal others to go
		theChan <- true
		<-theChan
	} else {
		// Not all here yet - wait for signal
		<-theChan
		theChan <- true // Pass signal to next goroutine
	}

	// Decrement counter
	atomic.AddInt32(arrived, -1)

	fmt.Println("Part B", goNum)

} //end-doSomething

func main() {
	totalRoutines := 10
	var arrived int32 = 0
	var wg sync.WaitGroup
	wg.Add(totalRoutines)
	//we will need some of these

	//var theLock sync.Mutex
	theChan := make(chan bool)     //use unbuffered channel in place of semaphore
	for i := range totalRoutines { //create the go Routines here
		go doSomething(i, &arrived, totalRoutines, &wg, theChan)
	}
	wg.Wait() //wait for everyone to finish before exiting
} //end-main
