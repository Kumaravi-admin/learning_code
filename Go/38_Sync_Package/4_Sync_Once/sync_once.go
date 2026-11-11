/*

                  sync.Once


• Ensure that a piece of code is only executed once.
• Thread safe concurrency primitive.


                  Uses


• Package level initialization.
• Lazy initialization.


                  Using it


• Create a instance of sync.Once.
• Pass function to the Do method.
• First call to Do method executes the function.
• Subsequent calls will have no implications.



*/

package main

import (
	"fmt"
	"sync"
)

// Step 1: Create a sync.Once. Think of it as a "one-time switch".
// It remembers whether the function has already been run.
var once sync.Once
var initialized bool

// Step 2: The setup code that must run only ONE time.
// (Example: connect to a database, load config, create a cache.)
func initSomething() {
	fmt.Println("initalizeing...")
	initialized = true
}

func doSomething(wg *sync.WaitGroup) {

	// Step 3: Pass the setup function to Do().
	// - The FIRST goroutine that reaches this line runs initSomething.
	// - Every other goroutine skips it (they do not run it again).
	// - Do() also makes the other goroutines WAIT until the first one finishes,
	//   so nobody uses the data before it is ready.
	once.Do(initSomething)

	// Step 4: This line runs for EVERY goroutine, because it is outside once.Do.
	fmt.Println("Doing somehting...")
	wg.Done()
}

func main() {
	wg := &sync.WaitGroup{}

	// Step 5: Start 3 goroutines. All 3 call doSomething at about the same time.
	wg.Add(3)

	for range 3 {
		go doSomething(wg)
	}

	wg.Wait()
}

/*  Terminal output:

initalizeing...
Doing somehting...
Doing somehting...
Doing somehting...

*/