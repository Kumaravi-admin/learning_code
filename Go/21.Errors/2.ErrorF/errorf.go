package main

import (
	"fmt"
)

func Divide(a, b float64) (float64, error) {
	if b == 0 {
		// Step 1: fmt.Errorf creates an error with a FORMATTED message.
		// Unlike errors.New, it lets you insert values using verbs like %f (float), %d (int), %s (string), %v (any).
		// Here %f is replaced by the value of b.
		return 0, fmt.Errorf("cannot divide by zero %f", b)
	}

	// Step 2: No problem happened, so return the result and nil (no error).
	return a / b, nil
}

func main() {

	// Step 3: Receive both the result and the error.
	result, err := Divide(10, 0)

	// Step 4: Always check err first.
	if err != nil {
		fmt.Println("error: ", err)
		return
	}

	fmt.Println("result: ", result)
}

/*  Terminal output:

error:  cannot divide by zero 0.000000

*/