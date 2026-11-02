/*
                      Error

• Go has a different way of error handling.
• There is no try-catch but uses if checks for error handling.
• The last return value is error by convention.
• The error string should not be capitalized and does not end with punctuations.

                   Creating Errors

• Implement the error interface.
• Use the New function from the errors package.
• Use the Errorf function from the fmt package.

*/

package main

import (
	"errors"
	"fmt"
)

// Step 1: Return the result AND an error (error is always the last return value).
func Divide(a, b float64) (float64, error) {
	if b == 0 {
		
		// Step 2: Create an error with errors.New().
		// Message is lowercase and has no punctuation at the end (Go convention).
		// Alternative: fmt.Errorf("cannot divide %v by zero", a) lets you format the message.
		return 0, errors.New("cannot divide by zero")
	}

	// Step 3: No problem happened, so return the result and nil (nil = no error).
	return a / b, nil
}

func main() {

	// Step 4: Receive both values from the function.
	result, err := Divide(10, 0)

	// Step 5: Check the error first (Go uses if checks, not try-catch).
	if err != nil {
		fmt.Println("error: ", err)
		return // stop here, do not use result when there is an error
	}

	// Only runs when err is nil
	fmt.Println("result: ", result)
}

/*  Terminal output:

error:  cannot divide by zero

*/