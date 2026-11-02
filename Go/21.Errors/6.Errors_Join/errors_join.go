/*
                 errors.Join()

• Go 1.20 introduced errors.Join() function  to wrap multiple errors.
• Join wraps provided error, discarding nil errors.
• If all errors are nil, it returns nil.
• A non-nil error returned by Join implements the Unwrap() []error method.

*/

package main

import (
	"errors"
	"fmt"
)

// Step 1: The first function creates an error.
func fistFunction() error {
	return fmt.Errorf("original error: something went wrong infistFuction")
}

func secondFunction() error {
	firstErr := fistFunction()
	if firstErr != nil {
		// Step 2: Create a second, separate error.
		secondErr := errors.New("failed in second fucntion")

		// Step 3: Join both errors into ONE error.
		// Unlike %w (which wraps a single error), Join keeps MULTIPLE errors together.
		// nil errors passed to Join are ignored automatically.
		return errors.Join(firstErr, secondErr)
	}
	return nil
}

func main() {

	// Step 4: Print the joined error.
	// Each error is printed on its own line (Join separates them with a newline).
	err := secondFunction()
	fmt.Println(err)

}

/*  Terminal output:

original error: something went wrong infistFuction
failed in second fucntion

*/