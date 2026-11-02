package main

import (
	"errors"
	"fmt"
)

// Step 1: Create a custom error type that can HOLD another error inside it.
type CustomError struct {
	Message string
	Wrapped error // the inner error being wrapped
}

// Step 2: Error() makes CustomError satisfy the built-in error interface.
// %v prints the inner error's text after the message.
func (e CustomError) Error() string {
	return fmt.Sprintf("%s: %v", e.Message, e.Wrapped)
}

// Step 3: Unwrap() is what makes this a "wrapping" error.
// It returns the inner error, so errors.Unwrap / errors.Is / errors.As
// can look inside CustomError.
func (e CustomError) Unwrap() error {
	return e.Wrapped
}

func SomeFunction() error {
	// Step 4: Build a CustomError that wraps a simple error.
	return CustomError{
		Message: "original erro: something went wrong",
		Wrapped: errors.New("wrapped error"),
	}
}

func main() {

	// Step 5: Print the outer error. Output is "Message: inner error text".
	err := SomeFunction()
	fmt.Println("Error: ", err)

	// Step 6: Unwrap one layer. This calls our Unwrap() method
	// and gives back e.Wrapped (the "wrapped error").
	err = errors.Unwrap(err)

	// Step 7: Wrap that inner error once more using %w.
	// Now it is a new layer with the text "innermost: " added in front.
	innerError := fmt.Errorf("innermost: %w", err)
	fmt.Println("Innermost Error: ", innerError)
}

/*  Terminal output:

Error:  original erro: something went wrong: wrapped error
Innermost Error:  innermost: wrapped error

*/