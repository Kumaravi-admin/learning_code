/*

                  Unwrapping Errors

• If e.Unwrap() returns a non-nil error w or  a slice containing w, then we say that e wraps w.
• It is invalid for an Unwrap method to return an []error containing a nil error value.
• A nil error returned from e.Unwrap() indicates that e does not wrap any error.
*/

package main

import (
	"errors"
	"fmt"
)

// Step 1: The innermost function creates the ORIGINAL error.
func fistFunction() error {
	// Simulating an error condition
	return fmt.Errorf("original error: something went wrong")
}

// secondFuction calls FirstFunction and wraps its error with additional context.
func secondFunction() error {
	err := fistFunction()
	if err != nil {
		// Step 2: Wrap the error with additional context using %w.
		// The original error stays inside, so it can be unwrapped later.
		return fmt.Errorf("failed in AnotherFunction: %w", err)
	}
	return nil
}

func main() {

	// Step 3: Get the outer (wrapped) error.
	// Printing it shows the full chain: outer message + inner message.
	err := secondFunction()
	fmt.Println("Original error: ", err)

	// Step 4: Unwrap it to get the error stored INSIDE.
	// errors.Unwrap peels off one layer and returns the inner error.
	// If the error wraps nothing, Unwrap returns nil.
	innerError := errors.Unwrap(err)
	fmt.Println("inner error: ", innerError)
}

/*  Terminal output:

Original error:  failed in AnotherFunction: original error: something went wrong
inner error:  original error: something went wrong

*/