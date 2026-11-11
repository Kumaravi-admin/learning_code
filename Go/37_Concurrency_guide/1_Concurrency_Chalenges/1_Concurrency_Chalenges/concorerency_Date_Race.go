/*

              What is concurrency?


• Dividing and executing programs in parts.
• Concurrency in not parallelism.
• Although concurrency may give an illusion of parallelism.


              Need of concurrency?


• If implemented correctly it can increase the performance of the program.
• Concurrency is evolution to extract more performance from computers.
• Meant to exploit the multiple CPU now available.


              The Challenges


• Concurrency comes with inherent challenges like race conditions.
• When two or more concurrent stream of a program try to access a shared resource.
• It cause problems as computer programs need to be deterministic.


                  Atomicity


• Atomicity refers to an operation that is indivisible and irreducible.
• Database transactions are one example of atomic operations.
• By employing atomic operations, we can ensure that critical sections of code are executed as a single unit.


                  Critical Section


• Accessing and modifying shared data is done in critical sections.
• Identify critical sections of our code is very important.
*/

package main

import (
	"fmt"
)

func main() {

	// Step 1: A shared variable. All goroutines below will modify THIS same variable.
	var counter int

	// Step 2: Start 1000 goroutines (small tasks running concurrently).
	// "for range 1000" simply repeats 1000 times (needs Go 1.22 or newer).
	for range 1000 {

		// "go" starts the function in the background and moves on immediately.
		go func() {

			// CRITICAL SECTION: this line reads and modifies shared data.
			// counter++ looks like one step, but it is really THREE steps:
			//   1. read counter   2. add 1   3. write back
			// Two goroutines can read the same old value, both add 1,
			// and write the same result, so one increment is LOST.
			// This is a RACE CONDITION. The operation is NOT atomic.
			counter++
		}()
	}

	// Step 3: Print the result.
	// Problem: main does not wait for the goroutines to finish,
	// so some may not even have run yet when we print.
	fmt.Println("Expected coounter: ", 1000)
	fmt.Println("Actual counter:", counter)
}

/*  Terminal output (different every run, almost never 1000):

Expected coounter:  1000
Actual counter: 995



*/