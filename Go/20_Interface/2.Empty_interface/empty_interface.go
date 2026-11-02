/*

                  Empty Interface


• The empty interface type is written as interface{}.
• It can store any value of type that has one or more method.
• This matches all types in Go.
• It can be used to represent values where type is not known.

*/

package main

import "fmt"

// Step 1: Accept a value of ANY type.
// interface{} has no methods, so every type in Go matches it.
// (any is just a newer, shorter name for interface{}. Both work the same.)

// func printValue(value any) {
// 	fmt.Println(value)
// }

func printValue(value interface{}) {
	fmt.Println(value)
}

func main() {

	// Step 2: A slice of interface{} can hold different types together.
	// Here: int, string, float64 and bool in one slice.
	mixedSlice := []interface{}{42, "Hello", 3.14, true}

	// Step 3: Pass it to the function.
	// Output: [42 Hello 3.14 true]
	printValue(mixedSlice)

}