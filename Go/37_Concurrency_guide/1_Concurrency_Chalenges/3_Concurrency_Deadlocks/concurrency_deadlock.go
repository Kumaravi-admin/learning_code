/*

              Deadlocks


• Processes blocking each other for a resource.



              Coffman Conditions


• Mutual Exclusion: At least one resource that can be held by only one process at a time.
• Hold & Wait: Process should be able to hold while waiting.
• No Preemption: Process holding resource cannot be forced to release it.
• Circular Wait: Chain of waiting processes.


                  Prevention


• Acquire locks in same order.
• At least one of the Coffman conditions is not satisfied.


*/

package main

import (
	"fmt"
	"sync"
)

// Step 1: Two locks. Think of them as two keys. Only one goroutine can hold a key at a time.
// (Coffman condition: Mutual Exclusion)
var (
	mutex1 sync.Mutex
	mutex2 sync.Mutex
)

// Step 2: resource1 takes key 1 first, then asks for key 2 (by calling resource2).
func resource1(name string) {
	mutex1.Lock() // take key 1
	fmt.Println(name, "acquired resource 1")

	// Holding key 1 while waiting for key 2.
	// (Coffman condition: Hold & Wait)
	resource2(name)

	// This line is never reached, because the line above gets stuck.
	mutex1.Unlock()
}

// Step 3: resource2 takes key 2 first, then asks for key 1 (by calling resource1).
// The order is OPPOSITE of resource1, and this is what creates the deadlock.
func resource2(name string) {
	mutex2.Lock() // take key 2
	fmt.Println(name, "acqured resource 2")

	// Holding key 2 while waiting for key 1.
	// If another goroutine holds key 1 and wants key 2, both wait forever.
	// (Coffman condition: Circular Wait)
	resource1(name)

	// Never reached for the same reason.
	mutex2.Unlock()
}

func main() {

	// Step 4: Start two goroutines that grab the keys in opposite order.
	//   Goroutine A: holds key 1, waits for key 2
	//   Goroutine B: holds key 2, waits for key 1
	// Neither can continue, and Go cannot force a lock to be released.
	// (Coffman condition: No Preemption)
	go resource1("Goroutine 1")
	go resource2("Goroutine 1")

	fmt.Println("Waiting...")

	// Step 5: select{} blocks main forever, so the program does not exit.
	// Go sees that every goroutine is stuck and stops with a fatal error.
	select {}
}

/*  Terminal output (order of lines can vary):

Waiting...
Goroutine 1 acquired resource 1
Goroutine 1 acqured resource 2
fatal error: all goroutines are asleep - deadlock!

*/