/*

                             Closures

  •  Closures are functions that close over values.
  •  It captures the values from its surrounding scope usually a function  block.
  •  Closure keeps access to the closed over variables ever after the surrounding scope has been executed.
  •  Closure can also been seen as function with state associated.

                             Benefits

  •  Closures can access and use variables that are not directly passed as arguments by closing over them.
  •  They can operate upon captured variables independently.
  •  Closures are used creating memoization functions, implementing callbacks, etc.


                             Example

┌─────────────────────────────────────────────────────────────┐
│                                                             │
│  import "fmt"                                               │
│                                                             │
│  func main() {                                              │
│       message := "Hello, "                                  │
│                                                             │
│       greet := func(name string) {                          │
│            fmt.Println(message + name)                      │
│       }                                                     │
│                                                             │
│       greet("Alice") // Hello, Alice                        │
│       greet("Bob") // Hello, Bob                            │
│  }                                                          │
│                                                             │
└─────────────────────────────────────────────────────────────┘
*/

package main

// In this lesson, we will cover:
// 09. Function closers

import "fmt"

// 09. Function closers
func counter() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
		
	}



func main() {

	// 09. Call Function closer (work as react memo its memorise values)
	increment := counter()
	fmt.Println(increment())			// output: 1
	fmt.Println(increment())			// output: 2


}


/*				Terminal Output: 

1
2


*/