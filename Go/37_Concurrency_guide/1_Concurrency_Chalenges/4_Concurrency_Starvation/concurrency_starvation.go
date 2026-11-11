/*
                  Starvation


• One process denied resource by another.
• Lower priority process keeps waiting for resource.
• Starvation is not limited to concurrent programming.


                  Prevention


• Scheduling algorithms & priority queues.
• Using of aging of process waiting.
• Fair resource allocation algorithms.

*/

package main

import (
	"fmt"
	"sync"
	"time"
)

// Step 1: One lock and one shared number.
// Every goroutine needs the SAME lock to update the counter.
var (
	mutex   sync.Mutex
	counter int
)

// Step 2: The "high priority" job. It adds 10 to the counter.
// It holds the lock for a while (100 microseconds of sleep),
// so other goroutines have to wait longer for their turn.
func highPriority() {
	mutex.Lock()
	time.Sleep(100 * time.Microsecond) // pretend this job is slow
	defer mutex.Unlock()
	counter += 10
}

// Step 3: The "low priority" job. It adds only 1 to the counter.
// It is quick, but it must wait until the lock is free.
func lowPriority() {
	mutex.Lock()
	defer mutex.Unlock()
	counter += 1
}

func main() {

	// Step 4: Start 1000 goroutines of EACH type, all competing for ONE lock.
	// The idea: the slow high priority jobs keep grabbing the lock,
	// so the low priority jobs may keep waiting (starvation).
	for i := 0; i < 1000; i++ {
		go highPriority()
		go lowPriority()
	}

	// Step 5: Give them 1 second to finish, then print.
	time.Sleep(1 * time.Second)
	fmt.Println("Final counter value:", counter)
}

/*  Terminal output (changes every run):

Final counter value: 11000

*/