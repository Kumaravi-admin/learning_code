/*
		Recover

	  - The recover() function can be used to recover from panics.
	  - Function recover() needs to be called in a deferred function.
	  - The return value of calling recover() is the argument provided to panic().


	                    Use Of Recover

• Purpose of recover is to gracefully shutdown.
• Not designed to act like exception handling.
• If you recover, code can and most probably  will panic again.
• Used in packages to handle panic and translate to a meaningful error.
*/

package main

import "fmt"

func panicExample() {
	// Step 1: Register a deferred function BEFORE the panic happens.
	// recover() only works inside a deferred function.
	defer func() {
		// Step 3: recover() stops the panic and returns the value given to panic().
		// If there was no panic, it returns nil, so we check r != nil.
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
	}()

	// Step 2: Panic stops normal execution of this function.
	// Go now runs the deferred function (the one above) before leaving.
	panic("Something went wrong.")

	// Any code placed here would NOT run (the function already stopped).
}

func main() {

	fmt.Println("Start of the program")

	// Step 4: The panic is recovered inside panicExample,
	// so it does not crash the program.
	panicExample()

	// Step 5: main continues normally because the panic was recovered.
	fmt.Println("Exited function")
}

/*  Terminal output:

Start of the program
Recovered from panic: Something went wrong.
Exited function

*/