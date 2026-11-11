package main

import (
	"fmt"
	"sync"
	"time"
)

// Step 3: Each goroutine receives a POINTER to the WaitGroup (*sync.WaitGroup).
// It must be a pointer so all goroutines share the SAME counter.
// (A WaitGroup must not be copied, see the "Package Docs" slide.)
func printNumber(wg *sync.WaitGroup, number int) {

	// defer runs when this function ends. Done() reduces the counter by 1,
	// which tells the WaitGroup: "I am finished".
	defer wg.Done()

	fmt.Println("Printing number: ", number)
	time.Sleep(time.Second * 1) // pretend the work takes 1 second
}

func main() {

	// Step 1: Create a WaitGroup. Think of it as a counter that starts at 0.
	var wg sync.WaitGroup

	for i := 1; i <= 3; i++ {

		// Step 2: Add(1) = "one more goroutine to wait for" (counter + 1).
		// Call Add BEFORE starting the goroutine, and outside of it (slide: Using WaitGroup).
		wg.Add(1)

		// Start the goroutine in the background
		go printNumber(&wg, i)
	}

	// Step 4: Wait() blocks main here until the counter becomes 0
	// (all 3 goroutines have called Done).
	wg.Wait()

	// Step 5: This line runs only after every goroutine is finished.
	fmt.Println("All numbers printed.")
}

/*  Terminal output (order of the numbers can change):

Printing number:  3
Printing number:  1
Printing number:  2
All numbers printed.

*/