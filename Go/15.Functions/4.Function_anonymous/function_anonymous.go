/*

                        Anonymous Functions

  •  Anonymous functions like anonymous structs don't have a name.
  •  Anonymous functions are useful for defining a small piece of functionality inline without the need to give it a name.
  •  Anonymous functions are also used to provide callbacks.


                                Example

┌──────────────────────────────────────────────────────────────────────┐
│                                                                      │
│  func main() {                                                       │
│       func() {                                                       │
│            fmt.Println("This is an anonymous function.")             │
│       }()                                                            │
│                                                                      │
│                                                                      │
│       func(msg string) {                                             │
│            fmt.Println("Message:", msg)                              │
│       }("Hello")                                                     │
│  }                                                                   │
│                                                                      │
└──────────────────────────────────────────────────────────────────────┘
.*/

package main

// In this lesson, we will cover:
// 10. Anonymous Functions

import "fmt"




func main() {


	// 10. Basic Anonymous Functions (first example)
	func() {
		fmt.Println("This is an anonymous function.")				// output: This is an anonymous function.
	}()
	
	// 10. Anonymous Functions With perameter (second example example)
	func(message string) {
		fmt.Println(message)
	}("Hello Go Learner")											// output: Hello Go Learner

	
	// 10. Anonymous Fuctions 3rd example
	numbers := []int{1, 2, 3, 4, 5}
	squared := sliceOperation(numbers, func(i int) int {
		return i * i
	} )  
	fmt.Println("Output of squared: ",squared)						// output: [1 4 9 16 25]

	

}

// 10. Anonymous Fuctions 3rd example
func sliceOperation(s []int, op func(int) int) []int{
	result := make([]int,len(s))
	for i, v := range s {
		result[i] = op(v)
	}
	return result
}

/*				Terminal Output: 


This is an anonymous function.
Hello Go Learner
Output of squared:  [1 4 9 16 25]

*/