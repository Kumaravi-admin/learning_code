/*

                  Type Assertion

• Type assertion is used to assert type for an empty interface.
• Type assertion is different from type conversion and can panic.
• Type conversion: string(typeDerivedFromString)
• Type assertion: emptyInterface.(string)



.*/

package main

import "fmt"

func main() {

	// Step 1: Create an empty interface variable (any = interface{}).
	// It can hold a value of any type.
	var emptyinterface any

	// Step 2: Store a string inside it.
	// Go now knows it only as "any", not as "string".
	emptyinterface = "Hello World!"

	// Step 3: Type assertion with the "comma ok" pattern.
	// emptyinterface.(string) asks: "Is the real value a string?"
	// ok is true if yes, false if no. Because we check ok, the program will NOT panic.
	if str, ok := emptyinterface.(string); ok {
		// str is now a real string, so we can use it as one
		fmt.Println("The underlying value is a string: ", str)
	}

	// Without the ok check, a wrong guess panics:
	// n := emptyinterface.(int) // PANIC: value is a string, not an int
}

/*  Terminal output:

The underlying value is a string:  Hello World!

*/