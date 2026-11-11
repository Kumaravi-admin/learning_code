/*
              Need for WaitGroup


• The main goroutine.
• Additional goroutines are started from the main goroutine and then merged back.
• How do we prevent goroutine leak?



                  WaitGroup


• Additional concurrency primitives.
• The WaitGroup struct.
• WaitGroup wait for goroutine completion.


                  Package Docs

A WaitGroup waits for a collection of goroutines to finish. The
main goroutine calls Add to set the number of goroutines to wait
for. Then each of the goroutines runs and calls Done when
finished. At the same time, Wait can be used to block until all
goroutines have finished.

A WaitGroup must not be copied after first use.


                  Using WaitGroup


• WaitGroup is a memory safe counter.
• Methods Add and Wait should be called outside of goroutines.

*/

package main

import (
	"fmt"
	"sync"
)

// Step 3: Each goroutine gets a POINTER to the same WaitGroup.
// (Pointer, so all goroutines share ONE counter. A WaitGroup must not be copied.)
func sayHello(wg *sync.WaitGroup) {

	// Done() reduces the counter by 1 when this function ends: "I am finished".
	defer wg.Done()

	fmt.Println("Hello")
}

func main() {

	// Step 1: Create a WaitGroup (a counter that starts at 0).
	var wg sync.WaitGroup

	// Step 2: Add(2) = "I am going to wait for 2 goroutines" (counter = 2).
	// Called once, BEFORE the goroutines start, and outside of them (slide: Using WaitGroup).
	wg.Add(2)

	// Start 2 goroutines. Each one will call Done() once, so the counter goes 2 -> 1 -> 0.
	go sayHello(&wg)
	go sayHello(&wg)

	// Step 4: Wait() blocks main here until the counter reaches 0.
	wg.Wait()

	// Step 5: Runs only after both goroutines are finished.
	fmt.Println("Both goroutines completed")
}

/*  Terminal output:

Hello
Hello
Both goroutines completed

*/