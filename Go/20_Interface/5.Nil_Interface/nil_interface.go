// Nil interface

package main

import "fmt"

func main() {

	// Step 1: A nil pointer of type *int.
	// It points to nothing, so it is nil.
	var a *int

	// Step 2: An empty interface with no value and no type.
	// Inside, an interface holds two things: (type, value). Here both are empty.
	var i interface{}

	fmt.Println(a == nil) // true  -> the pointer is nil
	fmt.Println(i == nil) // true  -> interface has no type and no value

	// Step 3: Put the nil pointer inside the interface.
	// Now the interface holds (type = *int, value = nil).
	// It has a TYPE now, so it is no longer a nil interface.
	i = a
	fmt.Println(i == nil) // false -> interface is nil ONLY when both type and value are nil
}

/*  Terminal output:

true
true
false

*/