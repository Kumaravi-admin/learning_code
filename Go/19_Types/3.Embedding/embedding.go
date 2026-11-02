/*

           	 Composition & Promotion

• Go is object oriented language but does not have the concept of inheritance.
• It uses composition and promotion to promote code reuse.


                What is embedding?

• Inside a struct you write the name of the type that you want to embed.
• Any methods and fields of the embedded type can be directly accessed via the containing struct.
• If same fields exists on struct and embedded type, you need to explicitly specify the type when calling.



              	Embedding is inheritance?

• The two types still are two different types and cannot be used interchangeably.
• Methods of embedded types are promoted to the method set of the struct containing it.
• By embedding you can indeed satisfy interfaces.



                        Embedding

                                              ┌──► Details
  Details ◄───────┐                           │
                  │──────── ( A ) ────────────┼──► Size
  Size    ◄───────┘           ▲               │
                              │               └──► Report
                              │
                              │
  Report  ◄────────────────( B )



                        Summary

• Go enables code reuse with composition and promotion.
• Embedding another type inside a struct allows for composition and promotion.
• Embedding is not inheritance.
*/

package main

import "fmt"

// Engine represents an engine type
type Engine struct {
	HorsePower int
	Model      string
}

// Start initiates the engine
func (e *Engine) Start() {
	fmt.Println("Engine started!")
}

// Car represents a vehicle composed of a specific model and an embedded Engine
type Car struct {
	Model  string
	Engine // Struct Embedding: fields and methods of Engine are promoted to Car
}

// Drive moves the vehicle forward
// Note: Changed from pointer receiver (*Car) to value receiver (Car) to perfectly match 
// how myCar is instantiated and invoked as a value in main() without causing implicit conversion confusion.
func (c Car) Drive() {
	fmt.Printf("Driving my %s...\n", c.Model)
}

// GPS represents a GPS device
type GPS struct {
	Model string
}

func main() {

	// 1. Initializing an Embedded Struct
	// You must explicitly specify the embedded type name ('Engine') during literal initialization
	myCar := Car{
		Model:  "Toyota",
		Engine: Engine{Model: "V8", HorsePower: 200},
	}

	// 2. Field Resolution & Promotion Demonstration
	// Accessing Car's own field
	fmt.Println("Car Model:", myCar.Model)               // Output: Car Model: Toyota
	
	// Accessing an unambiguous promoted field directly from the parent struct
	fmt.Println("Car Horse power:", myCar.HorsePower)    // Output: Car Horse power: 200
	
	// Resolving a naming collision (both Car and Engine have a 'Model' field)
	// Accessing myCar.Model gives "Toyota", so we must specify .Engine to read "V8"
	fmt.Println("Car Engine model:", myCar.Engine.Model) // Output: Car Engine model: V8

	// 3. Method Promotion Demonstration
	// The Start() method from Engine is promoted directly into the method set of Car
	myCar.Start()                                        // Output: Engine started!

	// Executing Car's native method
	myCar.Drive()                                        // Output: Driving my Toyota...
}

/*              Terminal Output

Car Model: Toyota
Car Horse power: 200
Car Engine model: V8
Engine started!
Driving my Toyota...

*/
