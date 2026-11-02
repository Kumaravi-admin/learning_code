/*
		Types

	  - A user-defined types can be created with the type keyword.
	  - Derive a type from existing primitive or composite types.
	    ◦ type TodoList []string
	  - In the past lesson we have created struct types.
	    ◦ type student struct {...}

	    Abstract & Concrete Types

• Abstract type defines what the type should do.
• Example of abstract type is function type that we have created in previous lessons.
• Concrete defines what and how the type should do it.
• Example of concrete type is a struct with methods.

	Methods

• Methods are functions that are associated and belong to a type.
• You can define methods for both custom or derived types.
• Syntax is slightly different from functions.

	More About Methods

• There are no overloaded methods in Go.
• Methods can be accessed just like field i.e. dot notation.
• Methods are of two kind pointer receiver and value receiver.

	Which Receiver Type To Use?

• General rules of when to use pointers apply to  method receivers as well.
• Method modifying the receiver will need to use  pointer receiver.
• Such methods should also handle the case for nil  pointer receiver.
• Method accessing a value on receiver can just use value receiver.

	Method Sets

• With pointer you can call pointer receiver methods as well as value receiver methods.
• If value is addressable you can call pointer receiver methods as well as value receiver methods.
• If value isn't addressable, Go can't get a pointer to it, so you can only call value receiver methods.

		Method Expression

	  - There is another way of using methods defined  on types called method expression.
	  - This requires you to save the method into a variable.
	  - The first argument to this function will be the receiver.

	  - Examples:
	    ◦ f1 := Person.GetDetails; f1(p)
	    ◦ f2 := Rectangle.SetLength; f2(r, 1)

	    Summary

• You can create new type or derive from  existing type.
• Abstract type defines what a type should do.
• Concrete type define what and how is should be done.
• Methods are functions associated with types.
• Method type to use depends on the use case.
• Expect the pointer receiver to be nil and code accordingly.
• Pointer receiver methods can be used for values if the value is addressable.
*/
package main

import "fmt"

// 1. Value Receiver Method
type Person struct {
	Name string
	Age  int
}

// GetDetails uses a Value Receiver (p Person) to read data without modifying it
func (p Person) GetDetails() string {
	return fmt.Sprintf("Name: %s, Age: %d", p.Name, p.Age)
}

// 2. Custom Type and Area Method
type Rectangle struct {
	Length float64 
	Width  float64
}

// Area uses a Value Receiver to calculate and return a value
func (r Rectangle) Area() float64 {
	return r.Length * r.Width
}

// 3. Pointer Receiver Method (Mutator)
// SetLength uses a Pointer Receiver (*Rectangle) so it can modify the original struct's data
func (r *Rectangle) SetLength(l float64) {
	r.Length = l
}

func main() {

	// 1. Calling a Value Receiver Method
	developer := Person{Name: "Anjesh", Age: 25}
	fmt.Println("Output of developer: ", developer.GetDetails())            // Output: Name: Anjesh, Age: 25

	// 2. Calling a Method on a Custom Type
	rect := Rectangle{Length: 5.0, Width: 3.0}
	fmt.Println("Output of rect Area: ", rect.Area())                       // Output: 15

	// 3. Calling a Pointer Receiver Method to Update Data
	fmt.Println("Rect length before update: ", rect.Length)                 // Output: 5
	rect.SetLength(10)                                                      // Go automatically handles the pointer transformation here
	fmt.Println("Rect length after update: ", rect.Length)                  // Output: 10
	
	// 4. Method Expression Example
	// Demystifying Method Expressions: Turning a method into a regular function by calling it directly from the Type.
	// The receiver type (Person) becomes the first argument of the function variable (f1).
	f1 := Person.GetDetails
	p := Person{Name: "Anjesh", Age:30}
	fmt.Println("Output of Function Expression: ", f1(p))                    // Output: Name: Anjesh, Age: 30
}

/*              Terminal Output

Output of developer:  Name: Anjesh, Age: 25
Output of rect Area:  15
Rect length before update:  5
Rect length after update:  10
Output of Function Expression:  Name: Anjesh, Age: 30

*/
