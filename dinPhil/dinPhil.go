// Dining Philosophers Code
// Author: Emmanuel Abolade
// Course: Concurrent Development, Year 4, SETU
// Instructor: Dr. Joseph Kehoe
// Date: October 2025
//
// License: GPL v3 - See LICENSE file for details

// [Brief description of what this file does]
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program. If not, see <https://www.gnu.org/licenses/>.

//PROBLEM BACKGROUND:
// Five philosophers sit at a round table with five forks between them.
// Each philosopher needs TWO forks to eat (left and right).
// If all philosophers pick up their left fork simultaneously, deadlock occurs
// because everyone waits forever for their right fork (circular wait).

// SOLUTION: Asymmetric Fork Pickup
// This solution prevents deadlock by breaking the circular wait condition.
// Philosophers 0-3 pick up LEFT fork first, then RIGHT fork.
// Philosopher 4 picks up RIGHT fork first, then LEFT fork.
// This asymmetry ensures that deadlock cannot occur.

// WHY THIS SOLUTION WORKS:
// By having one philosopher use opposite order, we break the symmetry.
// This prevents the circular chain where everyone waits for the next person.
// Proof: If Phil 0-3 hold left forks and Phil 4 holds its right fork (fork 4),
// then Phil 4 got fork 4 first and will get fork 3 next, so Phil 3 can't have
// fork 3. This contradicts our assumption that everyone is waiting.
// Therefore, deadlock is impossible.

package main

import (
	"fmt"
	"math/rand/v2"
	"sync"
	"time"
)

// Constants define the simulation parameters
const philCount = 5 // Number of philosophers (and forks)
//const iterations = 3 // Number of think-eat cycles each philosopher performs

// think simulates a philosopher thinking for a random duration
// The philosopher's identification number (0-4)
func think(index int) {
	// Generate random sleep time between 0-2 seconds
	X := time.Duration(rand.IntN(3))
	time.Sleep(X * time.Second) //wait random time amount
	fmt.Println("Phil: ", index, "was thinking")
}

// eat simulates a philosopher eating for a random duration
// index: The philosopher's identification number (0-4)
func eat(index int) {
	// Generate random sleep time between 0-2 seconds
	var X time.Duration
	X = time.Duration(rand.IntN(3))
	time.Sleep(X * time.Second) //wait random time amount
	fmt.Println("Phil: ", index, "was eating")
}

// getForks acquires both forks needed for eating using asymmetric pickup order
// -the KEY function that prevents deadlock through asymmetry.
//
//   - index: The philosopher's identification number (0-4)
//   - forks: Map of fork channels where each fork is a buffered channel
//
// Fork Acquisition Strategy:
//   - Philosophers 0-3: Pick up LEFT fork first, then RIGHT fork
//   - Philosopher 4:    Pick up RIGHT fork first, then LEFT fork
//
// ASYMMETRIC: Last philosopher picks up forks in OPPOSITE order preventing a deadlock
func getForks(index int, forks map[int]chan bool) {

	// Calculate which forks this philosopher needs
	left := index                    // Left fork has same index as philosopher
	right := (index + 1) % philCount // Right fork (wraps around for philosopher 4)

	// Check if this is the last philosopher (who breaks symmetry)
	if index == philCount-1 {
		// Last philosopher: Pick up RIGHT first, then LEFT
		<-forks[right] // Receive from channel = acquire fork
		fmt.Printf("Philosopher %d: picked up RIGHT fork %d\n", index, right)
		<-forks[left] // Then acquire left fork
		fmt.Printf("Philosopher %d: picked up LEFT fork %d\n", index, left)
	} else {
		// Other philosophers: Pick up LEFT first, then RIGHT
		<-forks[left] // Receive from channel = acquire fork
		fmt.Printf("Philosopher %d: picked up LEFT fork %d\n", index, left)
		<-forks[right] // Then acquire right fork
		fmt.Printf("Philosopher %d: picked up RIGHT fork %d\n", index, right)
	}
	// Confirmation that both forks have been acquired
	fmt.Printf("Philosopher %d: acquired both forks\n", index)
}

//func putForks(index int, forks map[int]chan bool) {
//	<-forks[index]
//	<-forks[(index+1)%5]
//}

// putForks releases both forks back to the table
func putForks(index int, forks map[int]chan bool) {
	// Calculate which forks this philosopher is holding
	left := index
	right := (index + 1) % philCount

	// Release both forks by sending to their channels
	// Send to channel = release fork, making it available for others
	forks[left] <- true
	forks[right] <- true

	fmt.Printf("Philosopher %d: released both forks\n", index)
}

// think → acquire forks → eat → release forks → repeat
func doPhilStuff(index int, wg *sync.WaitGroup, forks map[int]chan bool) {
	defer wg.Done()
	for {
		// Standard Dining Philosophers cycle
		think(index)
		getForks(index, forks)
		eat(index)
		putForks(index, forks)
	}

}

// the simulation that coordinates all philosopher goroutines
func main() {
	// WaitGroup to wait for all philosopher goroutines to complete
	var wg sync.WaitGroup
	philCount := 5
	wg.Add(philCount)

	//   Sending to channel (fork <- true) = releasing/making available
	//   Receiving from channel (<-fork) = acquiring/taking
	//   Buffered channel (capacity 1) acts like a binary semaphore
	forks := make(map[int]chan bool)
	for k := range philCount {
		forks[k] = make(chan bool, 1)
		forks[k] <- true // Fork available
	} //set up forks (Launch all philosopher goroutines concurrently)
	for N := range philCount {
		go doPhilStuff(N, &wg, forks)
	} //start philosophers
	wg.Wait() //wait here until everyone (10 go routines) is done

} //main
