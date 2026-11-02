/*

               Methods & Functions

• Methods are attached to types while functions are standalone.
• Methods can be used for functions if the signature of method is same as function type.
• Similar to function methods can be saved in variables and passed around.

*/

package main

import "fmt"

// 1. Passing Methods as First-Class Functions
type Calculator struct {
	// Fields of the struct can be empty
}

// Add uses a pointer receiver but matches the signature: func(int, int) int
func (c *Calculator) Add(a, b int) int {
	return a + b
}

// AddFunction is a standard standalone function matching: func(int, int) int
func AddFunction(a, b int) int {
	return a + b
}

// ArithmeticOperation is a higher-order function that accepts any func(int, int) int
func ArithmeticOperation(f func(int, int) int, a, b int) int {
	return f(a, b)
}

// 2. Storing Method Values in Variables
type Printer struct{}

// Print is attached to *Printer but can be bound to a variable as a Method Value
func (p *Printer) Print(m string) {
	fmt.Println(m)
}

func main() {
	
	// 1. Executing Standard Functions
	fmt.Println(AddFunction(10, 15))                      // Output: 25
	fmt.Println(ArithmeticOperation(AddFunction, 10, 15)) // Output: 25

	// 2. Passing a Method Value to a Higher-Order Function
	// Go automatically binds the instance 'c' to create a 'Method Value' function
	c := Calculator{}
	fmt.Println(ArithmeticOperation(c.Add, 10, 15))       // Output: 25

	// 3. Storing a Method Value in a Variable 
	// This captures both the receiver instance ('printer') and the method ('Print')
	printer := &Printer{}
	printFunction := printer.Print 

	printFunction("Hello go Learner!")                     // Output: Hello go Learner!
}

/*              Terminal Output

25
25
25
Hello go Learner!

*/
