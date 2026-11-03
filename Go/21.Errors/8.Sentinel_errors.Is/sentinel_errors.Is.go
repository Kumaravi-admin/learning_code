/*
                 errors.Is()


• The errors.Is function accepts two parameters, thefirst is the error and second the instance of error we are checking against.
• Function signature: func Is(err, target error) bool
• Iterate through the errors in the error chain and  returns true if matching error instance found.
• Own implementation: func Is(error) bool

*/

package main

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
)

// Example 1: Sentinel error (a fixed, package-level error value used for comparison).
// Step 1: Define a custom error type.
type CustomError struct {
	Message string
}

// Step 2: Error() makes CustomError satisfy the error interface.
func (c CustomError) Error() string {
	return c.Message
}

// Step 3: Create ONE shared error value. Others will compare against it.
// By convention sentinel error names start with "Err".
var ErrFirstError = CustomError{Message: "first error"}

func someFunction() error {
	return ErrFirstError
}

func main() {

	err := someFunction()
	fmt.Println(err) // output: first error

	if err != nil {
		// Step 4: errors.Is checks if err (or any error wrapped inside it)
		// matches ErrFirstError. Better than err == ErrFirstError because
		// it still works when the error has been wrapped with %w.
		if errors.Is(err, ErrFirstError) {
			fmt.Println("sentinel error found")
		}
	} // output: sentinel error found

	// Example 2: Sentinel error from the standard library.
	// os.Open fails because the file does not exist.
	if _, err := os.Open("non-existing"); err != nil {
		// fs.ErrNotExist is a ready-made sentinel error.
		// errors.Is looks through the chain inside err and finds it.
		if errors.Is(err, fs.ErrNotExist) {
			fmt.Println("file does not exist")
		} else {
			fmt.Println(err)
		}
	} // output: file does not exist
}

/*  Terminal output:

first error
sentinel error found
file does not exist

*/