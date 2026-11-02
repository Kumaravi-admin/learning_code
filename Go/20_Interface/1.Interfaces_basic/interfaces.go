/*
              Uses / Use Cases

• Grouping types by shared behaviour.
• Writing functions that accept many different types.
• Making code easier to extend without changing existing functions.
• Swapping implementations for testing.
• Designing clear and simple program structures.


/*
                     Interfaces


• Interfaces are abstract types that specify what a type should do or the behaviour.
• To implement the interface, a type must implement the methods in the interface.
• Names of interfaces from the standard lib usually ends with ...er like io.Writer.



                 Example

┌──────────────────────────────────────┐
│                                      │
│                                      │
│    type Shape interface {            │
│        Area() float64                │
│    }                                 │
│                                      │
│    type Car interface {              │
│        Start()                       │
│        Stop() bool                   │
│    }                                 │
│                                      │
│                                      │
└──────────────────────────────────────┘

              Implementing Interface

• Interfaces in Go are implicit which is   different from other languages.
• If types implements the methods listed in the  interface, it meets the interface.
• A type can implement multiple interfaces.


                 Use of interface?


• Interfaces are very practical as your code  works with functionality and not how it is   implemented.
• Database interface implements methods but does not matter if it uses Postgres, MySQL, etc.
• This makes for very maintainable code.


                 Example

┌──────────────────────────────────────┐
│                                      │
│                                      │
│    type Shape interface {            │
│        Area() float64                │
│    }                                 │
│                                      │
│                                      │
│    type ColoredShape interface {     │
│        Shape                         │
│        Color() string                │
│    }                                 │
│                                      │
└──────────────────────────────────────┘
*/

package main

import "fmt"

// Step 1: Create an interface (the contract).
// Any type that has an Area() float64 method is automatically a Shape.
type Shape interface {
	Area() float64
}

// Step 2: Create types (structs).
type Circle struct {
	Radius float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

// Step 3: Add the Area() method to each type.
func (c Circle) Area() float64 {
	return 3.14 * c.Radius * c.Radius
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Step 4: Write ONE function that accepts the interface.
// It works with any Shape, so we never change it when new shapes are added.
func printArea(s Shape) {
	fmt.Println(s.Area())
}

func main() {

	// Step 5: Create values of different types.
	c := Circle{Radius: 3}
	r := Rectangle{Width: 4, Height: 5}

	// Step 6: Pass them to the same function.
	printArea(c) // output: 28.26
	printArea(r) // output: 20


}