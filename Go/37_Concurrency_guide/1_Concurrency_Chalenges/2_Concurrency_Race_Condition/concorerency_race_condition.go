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
	"sync" // sync package has Mutex, WaitGroup, etc.
)

func main() {

	// Step 1: The shared variable (all goroutines modify this same variable).
	var counter int

	// Step 2: A Mutex is a lock (mutual exclusion).
	// Only ONE goroutine can hold the lock at a time. Others wait for their turn.
	var lock sync.Mutex

	for range 1000 {
		go func() {

			// Step 3: Wrap the critical section with Lock() and Unlock().
			// Lock()   = "I am entering, everyone else wait"
			// Unlock() = "I am done, next one can enter"
			// Now counter++ runs one goroutine at a time, so no increment is lost.
			lock.Lock()
			counter++
			lock.Unlock()
		}()
	}

	// Step 4: Lock before reading too, so we do not read while a goroutine is writing.
	lock.Lock()
	fmt.Println("Expected coounter: ", 1000)
	fmt.Println("Actual counter:", counter)
	lock.Unlock()
}

/*  Terminal output (can still be less than 1000, see note below):

Expected coounter:  1000
Actual counter: 995

*/