package main

import (
	"fmt"
	"sync"
)

//TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>

func fib(N int) int {
	if N < 2 {
		return 1
	}
	return fib(N-1) + fib(N-2)

}

// Added threshold parameter to prevent goroutine explosion
func parFib(N int, threshold int) int {
	if N < 2 {
		return N // Fixed: was returning 1 for N=0 (should be 0)
	}

	// NEW: Use sequential version for small values
	if N < threshold {
		return fib(N)
	}
	// Only create goroutines for large values
	var wg sync.WaitGroup
	var A, B int
	wg.Add(2)

	go func(N int, Ans *int) {
		defer wg.Done()
		*Ans = parFib(N-1, threshold) // Pass threshold down
	}(N, &A)

	go func(N int, Ans *int) {
		defer wg.Done()
		*Ans = parFib(N-2, threshold) // Pass threshold down
	}(N, &B)

	wg.Wait()
	return A + B
}

func main() {
	//TIP <p>Press <shortcut actionId="ShowIntentionActions"/> when your caret is at the underlined text
	// to see how GoLand suggests fixing the warning.</p><p>Alternatively, if available, click the lightbulb to view possible fixes.</p>
	threshold := 20 // Only parallelize for N >= 20
	for i := 0; i < 10; i++ {
		n := i * 5
		Seq := fib(n)
		par := parFib(n, threshold)
		fmt.Println(Seq, "---", par)

	}

}
