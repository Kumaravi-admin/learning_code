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
------------------------------------------------------------------

                       First Class Functions

  •  In Go, functions are first-class citizens.
  •  First-class functions imply that functions can be treated as values.
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

 .*/

package main

import "fmt"

// In this lesson, we will cover:
// 1. Basic fuction declaration
// 2. Function with parameter
// 3. Function with return value
// 4. Function with multiple return value
// 5. Function with Variadic Parameters
// 6. Function Signature
// 7. Function return function
// 8. Function type

	type operation func (int, int) int // declaring type for Function type example

func main() {
	// 1. Call the basic function 			 							// output: Hi!
	sayHi()


	// 2. Call function pass paramenter to the function 
	sayHiToSomeone("Anjesh")											// output: Hi! Anjesh


	// 3. Call funciton it's return a full 
	full_name := fullName("Anjesh","Kumar")
	fmt.Println(full_name)												// output: Anjesh Kumar


	// 4. Call fuction it's return multiple value full name and lenth 
	fn,name_lenth := fullNameWithLenth("Anjesh","Kumar")
	fmt.Printf("Full name: %s Lenth: %d\n",fn,name_lenth)  				// output: Full name: Anjesh Kumar Lenth: 13


	// 5. Call Function with Variadic Parameters
	total := sum(1, 2, 3, 4, 5, 6, 7, 8, 9)
	fmt.Println("Total: ", total)										// output: 45

	nums := []int{1, 2, 3, 4, 5}
	sum := sum(nums...)
	fmt.Println("Sum: ", sum)											// output: 15


	// 6. Call Function Signature (take another function as input)
	double := apply(double,5)			
	fmt.Println("Result of double: ", double)							// output: 10

	triple := apply(triple,5)			
	fmt.Println("Result of triple: ", triple)			 				// output: 15

	fiveTime := apply(fiveTimes,5)			
	fmt.Println("Result of fiveTime: ", fiveTime)			 			// output: 25


	// 7. Call Function return function
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
func arthmeticOperation( op string) operation {
	switch op {
	case "add":
		return func(n1,n2 int) int{
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



/*		Terminal output

Hi!
Hi! Anjesh
Anjesh Kumar
Full name: Anjesh Kumar Lenth: 13

*/