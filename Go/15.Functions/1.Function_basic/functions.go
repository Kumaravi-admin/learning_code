/*

                              Functions

  •  Functions are modular, reusable blocks of code designed to accomplish specific tasks.
  •  Functions take input (optional), process it, and return output (optional).
  •  Functions enable us to structure and organize code.
  •  A program is basically a orchestration of function calls.

                              Functions

               ┌──────────┐
  Input ──────►│   Code   ├──────► Output
               └──────────┘
                 Function


         ┌──────────┐      ┌──────────┐      ┌──────────┐
  Input ─┤   Code   ├─────►┤   Code   ├─────►┤   Code   ├─► Output
         └──────────┘      └──────────┘      └──────────┘
           Function          Function          Function
              1                 2                 3

------------------------------------------------------------------

                        More About Functions

  •  In Go, the entry point function main takes no parameters and returns no value.
  •  We have utilized the fmt.Println() function to output statements to the console.
  •  Functions in Go can return multiple  values.

                               Syntax

		┌──────────────────────────────────────────────────┐
		│                                                  │
		│  func functionName(parameters) returnType{       │
		│       // Code block                              │
		│       return                                     │
		│  }                                               │
		│                                                  │
		└──────────────────────────────────────────────────┘
------------------------------------------------------------------

                          Variadic Parameters

  •  Go permits the passing of variadic parameters to a function.
  •  Variadic parameters, in essence, allow for the passing of any number of function arguments.

  •  Rules:
       •  Employ three dots ...
       •  It must be the final parameter.
  •  Within the function, these parameters are accessible as a slice.


                  	             Syntax

┌──────────────────────────────────────────────────────────────┐
│                                                              │
│  func functionName(paramName ...paramterType) returnType{    │
│       // Code block                                          │
│       return                                                 │
│  }                                                           │
│                                                              │
└──────────────────────────────────────────────────────────────┘
------------------------------------------------------------------

                           Default Params

  •  Go lacks named or default parameters, a feature found in some other languages.
  •  Aim to limit the number of parameters in a function to just a couple for clarity and simplicity.
  •  When dealing with lengthy parameter lists  or the need for optional values, consider utilizing a struct.


                           Example

┌─────────────────────────────────────────────────────────────┐
│                                                             │
│  type MathOperatonOptions struct {                          │
│       Operation string                                      │
│       ValueOne int                                          │
│       ValueTwo int                                          │
│       ValueThree int                                        │
│  }                                                          │
│                                                             │
│                                                             │
│  func mathOperation(opts MathOperatonOptions){              │
│       // Code block                                         │
│  }                                                          │
│                                                             │
└─────────────────────────────────────────────────────────────┘

*/

package main

import "fmt"

// In this lesson, we will cover:
// 1. Basic fuction declaration
// 2. Function with parameter
// 3. Function with return value
// 4. Function with multiple return value
// 5. Function with Variadic Parameters

// Other fuction lesson in other folder
// 6. Function Signature
// 7. Function return function
// 8. Function type
// 9. Function closers
// 10. Anonymous Functions
// 11. Defer Function


func main() {

	// 1. Call Basic fuction declaration			 					// output: Hi!
	sayHi()


	// 2. Call Function with parameter 
	sayHiToSomeone("Anjesh")											// output: Hi! Anjesh


	// 3. Call Function with return value 
	full_name := fullName("Anjesh","Kumar")
	fmt.Println(full_name)												// output: Anjesh Kumar


	// 4. Call Function with multiple return value
	fn,name_lenth := fullNameWithLenth("Anjesh","Kumar")
	fmt.Printf("Full name: %s Lenth: %d\n",fn,name_lenth)  				// output: Full name: Anjesh Kumar Lenth: 13


	// 5. Call Function with Variadic Parameters
	total := sum(1, 2, 3, 4, 5, 6, 7, 8, 9)
	fmt.Println("Total: ", total)										// output: 45

	nums := []int{1, 2, 3, 4, 5}
	sum := sum(nums...)
	fmt.Println("Sum: ", sum)											// output: 15


	

}	

// 1. Basic Function
func sayHi() {
	fmt.Println("Hi!")
}

// 2. Function with parameter
func sayHiToSomeone(name string) {
	fmt.Println("Hi!",name)
}

// 3. Function with return value with parameter
func fullName(firstName string, lastName string) string {
	return fmt.Sprintf("%s %s\n", firstName, lastName)
}

// 4. Function with multiple return value
func fullNameWithLenth(firstName string, lastName string) (string,int) {
	fn := fullName(firstName, lastName)
	return fn,len(fn)
}

// 5. Function with Variadic Parameters
func sum(nums ...int) int {
	total := 0
	for _, num := range nums {
		total += num
	}

	return total
}

/*		Terminal output



Hi!
Hi! Anjesh
Anjesh Kumar

Full name: Anjesh Kumar  Lenth: 13
Total:  45
Sum:  15
*/