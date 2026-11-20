/*

                  Introduction


      • Goroutines are easy to create.
      • Garbage collector does not clean Goroutines.
      • Responsibility of the creator to clean.


                Termination Paths


1. The job assigned to the goroutine is now completed.
2. There was an error when the goroutine was executed.
3. The parent goroutine needs to terminate it.

*/

package main

/*
We will cover in this lesson:

1. Goroutine leak (a goroutine stuck forever)
2. Termination path 3: the parent terminates the goroutine using a "done" channel
*/

import (
	"fmt"
	"sync"
	"time"
)

// Step 1: A function that starts a goroutine and returns a channel.
// done is read-only (<-chan): the parent uses it to tell the goroutine "stop now".
// struct{} is an empty value that takes no memory. It is used when we only need a signal, not data.
func someWork(done <-chan struct{}, wg *sync.WaitGroup) <-chan int {
	ch := make(chan int)

	go func() {
		// Tell the WaitGroup "I am finished" when this goroutine ends.
		defer wg.Done()

		// Step 2: The goroutine waits on TWO things, whichever happens first.
		select {

		// Case 1: Send 1 into ch. But nobody is reading ch (main ignores it),
		// so this case is NEVER ready. Without the done case below,
		// the goroutine would be stuck here forever. That is a GOROUTINE LEAK.
		case ch <- 1:

		// Case 2: done was closed by the parent, so we must stop.
		// Reading from a closed channel is not an error, it returns immediately.
		case <-done:
		}

		fmt.Println("Goroutne finished")
	}()

	return ch
}

func main() {
	wg := &sync.WaitGroup{}

	// Step 3: We will start 1 goroutine, so wait for 1.
	wg.Add(1)

	// The "done" channel is only a signal. Nothing is ever sent into it.
	done := make(chan struct{})

	// Start the goroutine. We ignore the returned channel, so the goroutine's send can never complete.
	someWork(done, wg)

	// Simulate main doing other work for 1 second. The goroutine is stuck in select all this time.
	time.Sleep(1 * time.Second)

	// Step 4: The parent terminates the goroutine (termination path 3).
	// close(done) makes "<-done" ready for everyone listening, so the goroutine wakes up and exits.
	close(done)

	// Step 5: Wait until the goroutine has really finished (wg.Done was called).
	wg.Wait()
}

/*  Terminal output:

Goroutne finished

*/