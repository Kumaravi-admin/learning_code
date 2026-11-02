/*
                 The Error Interface

• The error type is an interface type in Go.
• It has single method called error  returning a string.
• type error interface {
      Error() string
  }


*/

package main

import (
	"fmt"
)

// Step 1: Create your own error type (a struct).
// It can carry extra data like a code, not just a message.
type CustomError struct {
	Code    int
	Message string
}

// Step 2: Add the Error() string method.
// This is all Go needs. Because CustomError has Error() string,
// it automatically satisfies the built-in error interface.
func (c CustomError) Error() string {
	// Sprintf builds a formatted string and returns it (does not print)
	return fmt.Sprintf("Error %d: %s", c.Code, c.Message)
}

func Divide(a, b float64) (float64, error) {
	if b == 0 {
		// Step 3: Return the custom error.
		// It is allowed here because CustomError implements error.
		return 0, CustomError{Code: 1, Message: "cannot devide by zero"}
	}

	return a / b, nil
}

func main() {

	result, err := Divide(10, 0)

	// Step 4: Check the error as usual.
	// Println calls err.Error() automatically to get the text.
	if err != nil {
		fmt.Println("error: ", err)
		return
	}
	fmt.Println("result: ", result)
}

/*  Terminal output:

error:  Error 1: cannot devide by zero

*/