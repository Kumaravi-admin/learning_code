/*
                 errors.As()


• The errors.As function accepts two parameters the first is the error and the second a non-nil pointer
  to either a type that implements error, or to any interface type.
• Function signature: func As(err error, target any) bool
• Iterate through the errors in the error chain and
  matching target type, if found returns true and set the target to that error value.
• Own implementation: func As(interface{}) bool

*/

package main

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
)

// Step 1: Create a custom error type that carries extra data (Status).
type CustomError struct {
	Message string
	Status  int
}

// Step 2: Error() makes CustomError satisfy the error interface.
func (e CustomError) Error() string {
	return e.Message
}

// Step 3: This function returns the custom error as a normal "error".
// The caller only sees the error interface, not the Status field.
func SomeFunction() error {
	return CustomError{Message: "original error: something went wrong.", Status: 400}
}

func main() {
	err := SomeFunction()

	// Example 1: Extract a custom error type.
	// Step 4: Declare a variable of the TYPE you want to extract.
	var customErr CustomError

	// Step 5: errors.As checks if err (or anything wrapped inside it) is a CustomError.
	// If yes, it returns true AND copies the value into customErr.
	// You must pass a POINTER (&customErr), because As needs to fill the variable.
	if errors.As(err, &customErr) {
		// Now we can read the extra fields, which a plain error does not expose
		fmt.Println("Extracted CustomError: ", customErr.Message)
		fmt.Println("Extracted CustomError: ", customErr.Status)
	}

	// Example 2: Extract a standard library error type.
	// os.Open fails because the file does not exist.
	if _, err := os.Open("non-existing"); err != nil {
		// os.Open returns a *fs.PathError (a pointer type), so the variable is a pointer too.
		var pathError *fs.PathError
		if errors.As(err, &pathError) {
			// Read the Path field from the extracted error
			fmt.Println("Failed at path: ", pathError.Path)
		}
	}

}

/*  Terminal output:

Extracted CustomError:  original error: something went wrong.
Extracted CustomError:  400
Failed at path:  non-existing

*/