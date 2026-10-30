/*

                       First Class Functions

  •  In Go, functions are first-class citizens.
  •  First-class functions simply that functions can be treated as values.
       •  They are assignable to variables.
       •  They can be passed as arguments to other functions.
       •  They can also be returned from functions.
  •  This provides great flexibility when  writing code.


                        Function Signature

  •  Function signature has the parameters and return type.

  •  Example:

       1. func sum(a, b int) int =>
          func(int, int) int
       2. func printName(name string) =>


                             Example

┌─────────────────────────────────────────────────────────────────────────────┐
│                                                                             │
│  func main() {                                                              │
│       fullName := func(s1, s2 string) string {                              │
│            return fmt.Sprintf("%s %s", s1, s2)                              │
│       }                                                                     │
│       welcomeString := sayHello("code", "learn", fullName)                  │
│  }                                                                          │
│                                                                             │
│                                                                             │
│  func sayHello(first, last string, fn func(string, string) string) string { │
│       fullName := fn(first, last)                                           │
│       return fmt.Sprintf("Welcome %s", fullName)                            │
│  }                                                                          │
│                                                                             │
└─────────────────────────────────────────────────────────────────────────────┘

                              Example

┌──────────────────────────────────────────────────────────────────┐
│                                                                  │
│  func arithmeticOperation(op string) func(a, b int) int {        │
│       switch op {                                                │
│       case "add":                                                │
│            return func(a, b int) int {                           │
│                 return a + b                                     │
│            }                                                     │
│       case "subtract":                                           │
│            return func(a, b int) int {                           │
│                 return a - b                                     │
│            }                                                     │
│       ...                                                        │
│       }                                                          │
│  }                                                               │
│                                                                  │
└──────────────────────────────────────────────────────────────────┘

                          Function Type

  •  Passing around function signatures can lead to code that is less readable.
  •  Function signatures can be utilized to   define unique types.
  •  Example:

       func(a, b int) int =>
       type op func (int, int) int
  •  This enhances code readability.


*/

package main

import "fmt"

// In this lesson, we will cover:
// 6. Function Signature
// 7. Function return function
// 8. Function type


type operation func(int, int) int // declaring type for Function type example

func main() {


	// 6. Call Function Signature (take another function as input)
	double := apply(double,5)			
	fmt.Println("Result of double: ", double)							// output: 10

	triple := apply(triple,5)			
	fmt.Println("Result of triple: ", triple)			 				// output: 15

	fiveTime := apply(fiveTimes,5)			
	fmt.Println("Result of fiveTime: ", fiveTime)			 			// output: 25


	// 7. Call Function that return another function
	multiplyByTwo := multiplyBy(2)
	multiplyByThree := multiplyBy(3)

	result1 := multiplyByTwo(5)
	result2 := multiplyByThree(3)
	fmt.Printf("Result1: %d Result2: %d\n",result1,result2)				// output: Result1: 10 Result2: 9


	// 8. Call Function type
	var perform operation 
	perform = arthmeticOperation("add")
	result3 := perform(1,2)
	fmt.Println("Result3: ", result3)									// output: 3



}

// 6. Function Signature (take another function as input)
func apply(f func(int) int,x int) int {
	return f(x)
}

func double(x int) int {
	return x * 2
}

func triple(x int) int {
	return x * 3
}

func fiveTimes(x int) int {
	return x * 5
}


// 7. Function return function
func multiplyBy(multiplier int) func (int) int {
	return func(i int) int {
		return i * multiplier
	}
}


// 8. Function type
func arthmeticOperation(op string) operation {
	switch op {
	case "add":
		return func(n1, n2 int) int {
			return n1 + n2
		}
	case "subtract":
		return func(n1, n2 int) int {
			return n1 - n2
		}
	default:
		return func(n1, n2 int) int {
			return n1 + n2
		}
	}
}


/*				Terminal output:

Result of double:  10
Result of triple:  15
Result of fiveTime:  25
Result1: 10 Result2: 9
Result3:  3

*/