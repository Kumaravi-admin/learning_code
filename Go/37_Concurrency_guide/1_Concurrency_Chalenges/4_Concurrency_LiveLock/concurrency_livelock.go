/*

                  Livelocks


• Similar to deadlocks but processes continuously change state.
• Two processes continuous trying to acquire same resource from each other but timing causes livelocks.
• Stuck in loop of requesting and releasing resource.

                  Prevention

• Acquire mutexes in specific order.

*/

package main

import (
	"fmt"
	"sync"
)

// Step 1: Two locks (keys). Only one goroutine can hold a key at a time.
var (
	mutexA sync.Mutex
	mutexB sync.Mutex
)

// Step 2: routineA takes A first, then tries to take B.
func routineA(name string) {
	mutexA.Lock()
	defer mutexA.Unlock() // runs when this function returns
	fmt.Println(name, "has mutex A")
	fmt.Println(name, "trying to acquire mutex B...")

	// TryLock does NOT wait. It returns false immediately if the lock is taken.
	// (A normal Lock() would wait forever = deadlock. TryLock gives up = backing off.)
	if !mutexB.TryLock() {
		fmt.Println("cannot acquire mutex B. Backing off.")

		// Simulate a backoff (a tiny pause before retrying).
		// Both routines wait the SAME time, so they keep colliding.
		for i := 0; i < 1000; i++ {
		}

		// Retry by calling itself again. The idea: keep trying again and again.
		routineA(name)

	} else {
		// Got both locks, so the work can finally happen.
		defer mutexB.Unlock()
		fmt.Println(name, "acquire mutex b")
		fmt.Println(name, "doing some work....")
	}
}

// Step 3: routineB is the mirror of routineA.
// It takes B first, then tries to take A (opposite order).
func routineB(name string) {
	mutexB.Lock()
	defer mutexB.Unlock()
	fmt.Println(name, "has mutex B")
	fmt.Println(name, "trying to acquire mutex A...")

	if !mutexA.TryLock() {
		fmt.Println(name, "cannot acquire mutex A. Backing off.")

		// Simulate some backoff time
		for i := 0; i < 1000; i++ {
		}

		// Retry
		routineB(name)

	} else {
		defer mutexA.Unlock()
		fmt.Println(name, "acquired mutex A")
		fmt.Println(name, "doing some work...")
	}
}

func main() {

	// Step 4: Start both routines. They want the same two locks in OPPOSITE order.
	go routineA("Routine A")
	go routineB("Routine B")

	fmt.Println("Waiting....")

	// Block main forever so the program does not exit.
	select {}
}