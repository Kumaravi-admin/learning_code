/*

                  Cond


• Transmit signals between two goroutines
• Cond implements a condition variable, a rendezvous point for goroutines waiting for or announcing the occurrence of an event.

                  Why?


• Naive approach to use infinite loop.
• Better is the waiting goroutine that can sleep.


                  Structure


    • Locker
    • Wait()
    • Signal()
    • Broadcast()

*/

package main

import (
	"fmt"
	"sync"
	"time"
)

// Step 1: Create a Cond. It needs a lock (Locker), here a Mutex.
// cond.L is that lock. The Cond is the "signal" between goroutines.
var cond = sync.NewCond(&sync.Mutex{})

// Shared data: a slice that can hold at most 2 numbers in this program.
var integers = make([]int, 0, 10)

// Step 4: The "remover". After a delay it removes one item and SIGNALS the waiting goroutine.
func remove(delay time.Duration) {
	time.Sleep(delay)
	fmt.Println("Len before removing: ", len(integers))
	integers = integers[1:] // drop the first item
	fmt.Println("Len after removing: ", len(integers))

	// Signal() wakes up ONE goroutine that is sleeping in cond.Wait().
	// (Broadcast() would wake up ALL waiting goroutines.)
	cond.Signal()
}

// Step 2: The "adder". It adds 5 numbers, but only 2 can be in the slice at a time.
func add() {
	for i := 0; i < 5; i++ {

		// Take the lock before checking or changing shared data.
		cond.L.Lock()

		// Step 3: Slice is full (2 items), so we cannot add. Wait for a signal.
		// Wait() does 3 things: unlocks cond.L, puts this goroutine to SLEEP
		// (no busy loop, no CPU wasted), and locks cond.L again after it wakes up.
		// It is "for" and not "if": after waking up, re-check the condition.
		for len(integers) == 2 {
			cond.Wait()
			fmt.Println("Len before adding: ", len(integers))
		}

		// Space is available, so add the number.
		integers = append(integers, i)
		fmt.Println("Len after adding: ", len(integers))

		// Start a remover that frees one space after 1 second.
		go remove(1 * time.Second)

		cond.L.Unlock()
	}
}

func main() {
	go add()

	// Give the goroutines time to finish before main exits.
	// (In real code you would use sync.WaitGroup instead of Sleep.)
	time.Sleep(5 * time.Second)
}

/*  Terminal output (approximate, order can vary):

Len after adding:  1
Len after adding:  2
Len before removing:  2
Len after removing:  1
Len before removing:  1
Len after removing:  0
Len before adding:  0
Len after adding:  1
Len after adding:  2
...

*/