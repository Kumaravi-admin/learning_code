/*
                  Introduction


• Mutex stands for mutual exclusion.
• Enable access synchronization to a resource through memory.
• Used to guard critical sections of your code or the section that requires access to a shared resource.


                  Mutex in Go


• Mutex in Go provides you with the Lock and the Unlock methods.
• By convention, the Unlock method is called in a deferred statement.


                  Critical Section


• There is a reason why we call the section of code the critical section where mutexes are used.
• It is likely a bottleneck of your program.
• This is where RWMutex comes into play, let us have a look at it.


                  RWMutex


• Conceptually RWMutex is very similar to the Mutex struct that we have used in earlier examples.
• RWMutex provides a little more flexibility as compared to the Mutex.
• For the RWMutex the process requesting lock for reading only will be given access to the shared resource.

                  Summary


• RWMutex provides better performance for frequent reads compared to a regular mutex.
• Read locks are non-blocking, allowing multiple goroutines to read the shared resource concurrently.
• Write locks on the other hand are blocking, making sure that exclusive access is granted for modifying the resource.
*/

package main

import (
	"fmt"
	"sync"
	"time"
)

// Step 1: Shared data (count) + the lock that protects it.
// RWMutex has two modes:
//   - Lock / Unlock:   write lock, only ONE goroutine at a time
//   - RLock / RUnlock: read lock, MANY readers at the same time (but no writer)
type Counter struct {
	mu    sync.RWMutex
	count int
}

// Step 2: Writer. It changes the data, so it takes the WRITE lock.
func (c *Counter) Int() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.count++
}

// Step 3: Reader. It only reads, so the READ lock is enough.
// Many readers can enter together, which reduces the bottleneck.
func (c *Counter) Get() int {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.count
}

func main() {
	counter := &Counter{count: 0}

	wg := &sync.WaitGroup{}

	// Step 4: 10 writer goroutines. Each increments the counter 100 times.
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			for j := 0; j < 100; j++ {
				counter.Int()
			}
			wg.Done()
		}()
	}

	// Step 5: 10 reader goroutines. Add(10) matches the loop (10 times), so the counts are equal.
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			fmt.Println("Count: ", counter.Get())
			time.Sleep(time.Millisecond * 2)
			wg.Done()
		}()
	}

	// Step 6: Wait until all 20 goroutines (10 writers + 10 readers) have called Done().
	wg.Wait()
}

/*  Terminal output (the values change on every run):

Count:  932
Count:  932
Count:  932
Count:  932
Count:  932
Count:  932
Count:  933
Count:  933
Count:  931
Count:  933
...

*/