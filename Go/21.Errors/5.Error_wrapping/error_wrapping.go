/*

                   Error Wrapping

• Wrap errors to add additional information.
• An error e wraps another error if e's type  has one of the methods:
    ◦ Unwrap() error
    ◦ Unwrap() []error
• An error chain refers to series of errors that are wrapped together.

• You can use fmt.Errorf() with %w to wraperrors.
• If you don’t want to wrap error but just want add error from, use the %v placeholder in fmt.Errorf().

*/

package main

import "fmt"

// Step 1: The innermost function creates the ORIGINAL error.
func fistFunction() error {
	return fmt.Errorf("original error: something went wrong infistFuction")
}

// Step 2: This function calls the first one and WRAPS its error.
func secondFunction() error {
	firstErr := fistFunction()
	if firstErr != nil {
		// %w wraps firstErr inside a new error and adds extra context.
		// The original error is kept inside, so it can be unwrapped later
		// (using errors.Unwrap, errors.Is or errors.As).
		// If you used %v instead, only the text would be copied and the
		// original error would be lost (no wrapping).
		return fmt.Errorf("Failed in first function: %w", firstErr)
	}
	return nil
}

func main() {

	// Step 3: Print the final error. It shows the whole chain in one message:
	// outer message first, then the inner (original) error.
	err := secondFunction()
	fmt.Println(err)

}

/*  Terminal output:

Failed in first function: original error: something went wrong infistFuction

*/