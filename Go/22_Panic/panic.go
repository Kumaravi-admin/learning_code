/*
	More About Panic

• The panic function accepts an argument  of type any.
• When panic is called code execution stop and panic prints the error stack of where the panic occurred.
• Panic has a relationship with defer, it can be better understood with help of a illustration.

	                      Panic Execution Flow

					┌──────────────────────────────┐
				(6) │ func main() {                │ (1)
					│     defer func() {}()        │
					│     func1()                  │
					│ }                            │
					└──────────────────────────────┘
						▲                   │
						│                   ▼
					┌──────────────────────────────┐
				(5) │ func func1() {               │ (2)
					│     defer func() {}()        │
					│     func2()                  │
					│ }                            │
					└──────────────────────────────┘
						▲                   │
						│                   ▼
					┌──────────────────────────────┐
				(4) │ func func2() {               │ (3) ◄── Panic rais
					│     defer func() {}()        │
					│     panic("error")           │
					│ }                            │
					└──────────────────────────────┘
*/
package main

import "fmt"

func fuction1() {

	defer func() {
		fmt.Println("function1 deffered fuction called")
	}()
	function2()
}

func function2() {
	defer func() {
		fmt.Println("function2 deffered function called")
	}()
	panic("fuction2 panicked")
}

func main() {
	fuction1()	
}