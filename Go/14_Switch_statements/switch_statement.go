/*

                           Switch Statement

  •  The switch statement is employed when there are multiple potential code paths.
  •  Like the for loop, in Go the switch statement enables variable  initialization.
  •  Each case operates within its independent code block.
  •  It doesn't necessitate a break statement at the end of each case.


							Syntax

┌──────────────────────────────────────┐  ┌──────────────────────────────────────────────┐
│                                      │  │                                              │
│  // switch on some value             │  │  // switch on initialized value              │
│  switch value {                      │  │  switch wordLen := len(word); wordLen {      │
│       // cases                       │  │       // cases                               │
│  case match-condition:               │  │  case match-condition:                       │
│       // case block                  │  │       // case block                          │
│  case match-condition:               │  │  case match-condition:                       │
│       // case block                  │  │       // case block                          │
│  default:                            │  │  default:                                    │
│  // if no case matches               │  │  // if no case matches                       │
│  }                                   │  │  }                                           │
│                                      │  │                                              │
└──────────────────────────────────────┘  └──────────────────────────────────────────────┘

------------------------------------------------------------------------------------------


                          More about switch

  •  The switch statement allows you to match  multiple conditions with a single case statement.
  •  Additionally, there is a 'fallthrough' keyword that enables matching multiple cases in sequence.
  •  Prefer cases with multiple conditions over fallthrough for clearer code.


                          Syntax			Not Recommended

┌──────────────────────────────────────┐  ┌──────────────────────────────────────┐
│                                      │  │                                      │
│  // switch on some value             │  │  // switch on some value             │
│  switch value {                      │  │  switch value {                      │
│       // cases                       │  │       // cases                       │
│  case match-condition,               │  │  case match-condition:               │
│       match-condition:               │  │       fallthrough                    │
│       // case block                  │  │  case match-condition:               │
│  case match-condition:               │  │       // case block                  │
│       // case block                  │  │  case match-condition:               │
│  default:                            │  │       // case block                  │
│  // if no case matches               │  │  default:                            │
│  }                                   │  │  // if no case matches               │
│                                      │  │  }                                   │
│                                      │  │                                      │
└──────────────────────────────────────┘  └──────────────────────────────────────┘



.*/

package main

import "fmt"

func main() {

	// 1. Basic Swich statement
	day := "Monday"

	switch day {
	case "Monday":
		fmt.Println("It's Monday")
	case "Tuesday":
		fmt.Println("It's Tuesday")
	default:
		fmt.Println("It's note Monday or Tuesday")
	}

	// 2. Switch Statement with initialized value
	word := "Test"

	switch wordLen := len(word); wordLen {
	case 3:
		fmt.Println("word lenth is 3")
	case 4:
		fmt.Println("word lenth is 4")
	default:
		fmt.Println("word lenth not 3 or 4, but",word)		
	}



	// 3. Swich statement with multiple conditions
	word = "a"

	switch word {
	case "a","b":
		fmt.Println("a or b")
	case "d":
		fmt.Println("d")
	default:
		fmt.Println("not a,b, or d")
	}
}

/*			
Terminal Output: 

It's Monday
word lenth is 4
a or b

*/