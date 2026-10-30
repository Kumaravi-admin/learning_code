/*
	                            Code Blocks

	  •  In Go, variables can be created at  different levels within a program's structure.
	  •  These levels are known as code blocks,  or simply, blocks.
	  •  Variables defined in an outer block are  accessible within inner blocks, but the
	     reverse is not true.


		                              Example

┌──────────────────────────────────────────────────────────────┐
│                                                              │
│  // package level variable.                                  │
│  var a = 10                                                  │
│                                                              │
│                                                              │
│  func main() {                                               │
│       // function level variable.                            │
│       b := 20                                                │
│       {                                                      │
│            // block variable.                                │
│            c := 30                                           │
│            // a, b & c can be accessed here.                 │
│       }                                                      │
│       // a & b can be accessed here                          │
│       // but c cannot be accessed here                       │
│  }                                                           │
│                                                              │
└──────────────────────────────────────────────────────────────┘

	                          Go Code Blocks


	 ┌──────────────────────────────────────────┐
	 │         Universal Block                   ─► Functions like len, cap, etc
	 │ ┌──────────────────────────────────────┐ │
	 │ │       Package Block                  ────► Package level variables
	 │ │ ┌──────────────────────────────────┐ │ │
	 │ │ │  	 File Block	                  │ │ │
	 │ │ │ ┌──────────────────────────────┐ ──────► Imported packages
	 │ │ │ │                              │ │ │ │
	 │ │ │ │                              │ | | |
	 │ │ │ │   Function Block             │ │ │ │
	 │ │ │ │                              │ │ │ │	  Function parameters
	 │ │ │ │                              ────────► Any code blocks, created with {}
	 │ │ │ └──────────────────────────────┘ │ │ |
	 │ │ └──────────────────────────────────┘ │ |
	 │ └──────────────────────────────────────┘ |
	 └──────────────────────────────────────────┘



	                            Shadowing

	•  Shadowing occurs when you create a   variable with the same name in an inner
	   code block as in an outer code block.
	•  Inner variable takes precedence over the  outer variable.
	•  Shadowing should be avoided for code  clarity.
	•  Tools can help detect and fix shadowing.


	                            Example

┌────────────────────────────────────────────────────────────────────────┐
│                                                                        │
│  package main                                                          │
│                                                                        │
│  func main() {                                                         │
│       x := 10                                                          │
│                                                                        │
│       if true {                                                        │
│            x := 20 // This is shadowing the outer 'x'.                 │
│            fmt.Println(x) // This will print 20, not 10.               │
│       }                                                                │
│                                                                        │
│       fmt.Println(x) // This will print 10, the outer 'x'.             │
│  }                                                                     │
│                                                                        │
└────────────────────────────────────────────────────────────────────────┘
*/
package main

import "fmt"

var  (
	// Package level variable
	a = 10 
)

func main() {
	fmt.Println("Package lavel variable a: ",a)  // output: 10

   {
      // block scoped variable a
      a := 15
      fmt.Println("Block scoped variabl a: ",a) // output: 15
   }
}
