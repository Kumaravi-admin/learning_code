/*
                             Defer

  •  A "defer" statement invokes a function whose execution is deferred.
  •  Three condition for deferred function execution:

       •  Surrounding function returns.
       •  Code execution reaches the end of function  body.
       •  Function panics.

  •  The deferred call's arguments are evaluated immediately, but the function call is not executed right away.


                                Defer

  •  Multiple deferred functions are executed in last in first out manner (LIFO).
  •  Deferred functions are executed after any result parameters are set by that return statement but before the function returns to its caller.
  •  Used for cleanup and handling panics.

Example

┌ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ┐
  import "fmt"

  func main() {
      message := "Hello, "
      greetingFn := func(name string){
          fmt.Println(message + name)
      }
      defer greetingFn("Alice")
      defer greetingFn("Bob")
      message := "Hi, "
  }
└ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ┘


*/

package main

import "fmt"

// In this lesson, we will cover:
// 11. Defer Function



func main() {
	
	message := "Hello, "

	// 1st Example of defer function (defer function run l)
	greetingFn := func (name string){
		fmt.Println(message + name)
	}
	defer greetingFn("Anjesh")				// output: Hi, Anjesh
	defer greetingFn("Avyansh")				// output: Hi, Avyansh


	// 2nd Example of defer function 
	defer func(name string) {
		fmt.Println(message + name)

	}("Ankit")								// output: Hi, Ankit


	
	message = "Hi, "
	
}


/*				Terminal output

Hi, Ankit
Hi, Avyansh
Hi, Anjesh


*/