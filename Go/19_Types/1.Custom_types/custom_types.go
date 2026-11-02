/*
                 Derived Types

• We have already seen that you can derive types from other built in types. Example: type Age int
• You can also derive types from user defined types. Example: type Student Human
• Derived types from built-in types can be operated upon by type operators for those types.


                 Is it inheritance?

• The derived type and base type are two different types.
• They don’t share method or anything else.
• Go is object oriented language but it does not use inheritance but relies on composition for code reuse.

*/

package main

import "fmt"

// 1. Derived type from a built-in primitive type
type Age int

// 2. Custom Base Type
type Human struct {
	Age int
}

// Method attached to the base type (Human)
func (h Human) printAge() {
	fmt.Println(h.Age)
}

// 3. Derived type from a custom type
// Student is a completely unique type; it does NOT inherit methods from Human!
type Student Human


func main() {

	// 1. Primitive Derived Type Operations
	// Derived types can still use standard operators (+) of their underlying type (int)
	var young Age = 10
	var old Age = 60
	fmt.Println(young + old)             // Output: 70

  
	// 2. Passing Derived Type to a Function
	// This works perfectly because the type matches exactly
	sum(young)                           // Output: Age: 10


	// 3. Working with Custom Derived Types
	// Student keeps fields (Age) but DOES NOT inherit the printAge() method!
	var s = Student{Age: 10}
	fmt.Println("\nStudent Age:", s.Age) // Output: Student Age: 10


	// Proving No Inheritance:
	// s.printAge() -> This would throw a COMPILATION ERROR! 
	// To call it, you must explicitly convert it back to the base type:
	fmt.Print("Calling method via Type Conversion: ")
	Human(s).printAge()                  // Output: 10
}


// sum accepts our custom 'Age' type
func sum(a Age) {
	fmt.Printf("Age: %d", a)
}

/*              Terminal Output

70
Age: 10
Student Age: 10
Calling method via Type Conversion: 10

*/
