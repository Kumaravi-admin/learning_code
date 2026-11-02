/*


                        Iota


• Within a constant declaration, the predeclared
  identifier iota represents successive untyped
  integer constants.
• It be used to defined set of pre-defined
  limited values but it is not equivalent to
  enumerations.



                  Example

┌──────────────────────────────────────┐
│                                      │
│  type Size int                       │
│                                      │
│  const (                             │
│      ExtraSmall Size = iota          │
│      Small                           │
│      Medium                          │
│      Large                           │
│      ExtraLarge                      │
│  )                                   │
│                                      │
└──────────────────────────────────────┘


                   Summary


• Deriving types do not account for inheritance.
• Pre-declared identifier iota can be used to
  create increase constant value.
• Go do not support enumerations or inheritance.


*/

package main

import "fmt"

// Derived type from int to add type safety for sizes
type Size int

// Using iota to automatically assign incrementing integer constants starting at 0
const (
	ExtraSmall Size = iota // 0
	Small                  // 1 (Type and iota expression are implicitly repeated)
	Medium                 // 2
	Large                  // 3
	ExtraLarge             // 4
)

// printSize inspects the underlying integer value to print human-readable text
func printSize(s Size) {
	switch s {
	case ExtraSmall:
		fmt.Println("Extra Small")
	case Small:
		fmt.Println("Small")
	case Medium:
		fmt.Println("Medium")
	case Large:
		fmt.Println("Large")
	case ExtraLarge:
		fmt.Println("Extra Large")
	default:
		fmt.Printf("Unknown Size (%d)\n", s)
	}
}

func main() {
	// 1. Printing constant values
	fmt.Println("Constant values:", ExtraSmall, Small, Medium, Large, ExtraLarge) // Output: Constant values: 0 1 2 3 4

	// 2. Calling printSize with a valid integer mapping (4 matches ExtraLarge)
	printSize(4) // Output: Extra Large

	// 3. Handling values outside the defined constant range
	// Your previous comment said output: "0 1 2 3 4", which was incorrect for this call.
	// Since 10 matches nothing in the switch, it falls through safely.
	printSize(10) // Output: Unknown Size (10)
}

/*              Terminal Output

Constant values: 0 1 2 3 4
Extra Large
Unknown Size (10)

*/
