/*

				 Memory

• Programs are allocated memory by the operating system (OS) for execution.
• The allocated memory is divided into two key segments: Stack and Heap.
• In broad terms, the Stack manages function call execution, while the Heap handles dynamic storage of variables.
• Go runtime has built in Garbage Collector.


                  Pointers

• Initializing a variable in a program allocates  memory to store a value.
• A pointer is a variable that holds the memory  address of another variable.
• Pointers of any data type have the same size.
• The zero value for a pointer variable is nil.


                  Illustration

                     ┌────────┬────────┬────────┬─────────┐
  Variable Name      │  name  │ isEven │  age   │ section │
                     ├────────┼────────┼────────┼─────────┤
  Variable Values    │  "a"   │  true  │   10   │   "b"   │
                     ├────────┼────────┼────────┼─────────┤
  Memory Address     │   0    │   8    │   16   │   24    │
                     └────────┴────────┴────────┴─────────┘
                         ▲
                         │
                         │
                         │
       Pointer variable holds this value




                Pointer Operators

• When the asterisk (*) is used as a prefix with a  variable type, it signifies that it stores a pointer.
• It is also employed to retrieve the value stored  at the address pointed to by a pointer, known as  dereferencing.
• The ampersand (&) operator, when used as a prefix, provides the address of the variable.
• Pointers cannot be created for constants or primitive literals.
• Attempting to dereference a nil pointer will result in a runtime panic.



                Example
┌ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ┐
  var pointerToInt *int
  var pointerToString *string
  var pointerToBool *bool

  age := 10
  pointerToInt = &age

  ageValue := *pointerToInt
└ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ┘

              Performance Benefits


• Passing large structs as arguments in functions may lead to performance concerns.
• Pointers offer a way to mitigate this potential degradation.
• It's advisable to stick with passing values, unless the performance impact is substantial.
• Linting tools can assist in making informed decisions regarding this choice.

             Go is Garbage Collected

• Go is a garbage-collected language, relieving developers from manual memory deallocation concerns.
• The creation of pointer references results in heap allocations.
• Despite Go's garbage collection, it's important to be mindful of not overburdening the garbage  collector.
------------------------------------------------------------------

                 Mutability

• Variables can be categorized as either mutable or immutable.
• Mutable variables allow their values to be  changed.
• In contrast, immutable variables prohibit any  changes to their value.
• In Go within a function, you can enforce immutability with const and mutability with var.


              Mutability & Pointers

• Pointers are employed when functions need to alter the state of passed variables.
• With the address of the variable, a function can directly update its value.
• Caution must be exercised with mutability to prevent potential bugs; always document the behavior of your functions.


              Indication of no value

• Occasionally, it's important for the calling function to distinguish between a zero value and no value at all.
• When a function returns a pointer, it can use nil to signal the absence of a value.
• This pattern is widely used in Go code, particularly when a function returns both a value and a potential error.


                    Summary


• Pointers serve to indicate mutability, signal  absence of value, and optimize performance when necessary.
• Unless there's a specific need, prefer passing values over using pointers for simplicity and clarity.
• Optimizing memory usage to minimize garbage collector invocations enhances program performance.

.*/

package main

import "fmt" 

func main() {

	// 1. Create nil pointer
	var intPointer *int		
	fmt.Println(intPointer)							// output: <nil>


	// 2. Refrencing memory address
	age := 10
	intPointer = &age								// assign age memory address to intPointer
	fmt.Println(intPointer)							// output: 0x13ccb873a0a8


	// 3. Dereferencing (access value of pointer)
	fmt.Println(*intPointer)						// output: 10



	// 2nd example of pointer refrence and derefrencing
	name := "Anjesh Kumar"
	namePointer := &name							// Create pointer refer name momeory address

	fmt.Println(*namePointer)						// output: Anjesh Kumar


	

}


/*        Termnal output: 

<nil>
0x35db03bda0a8
10
Anjesh Kumar


*/