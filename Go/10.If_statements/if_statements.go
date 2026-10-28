/*
               The ‘if’ Statement

• The 'if' statement is employed to   conduct conditional checks based on
  boolean expressions, which are commonly  referred to as conditions.
• The code block within the 'if' statement  will be executed only when the boolean
  expression or condition evaluates to  true.


                       Example


The ‘if’ keyword              The boolean expression
     │             /────────   or condition
     │            /
     ▼         	 /
				▼
     if age >= 18 { <───────────────────────── The ‘if’ code
                                                  block
         fmt.Println("You are an adult!")

     }


------------------------------------------------------------------

              The ‘if-else’ Statement

• The 'if-else' statement serves as an extension of the 'if' statement.
• It is employed when there are precisely two conditions that need to be evaluated.
• This construct allows for distinct code blocks to be executed depending on
  whether the condition is true or false.


                     Example


     if age >= 18 { <───────────────────────── The ‘if’ code
                                                  block
         fmt.Println("You are an adult!")

     } else { <─────────────────────────────── The ‘else’ code
                                                  block
         fmt.Println("You are not an adult!")

     }
------------------------------------------------------------------

             The ‘if-else-if’ Statement

• An extension of the 'if-else' statement is the 'if-else-if' statement.
• This statement is employed when there are more than two conditions you want to  evaluate.
• You can create a chain of 'if-else-if' statements by adding blocks of 'else-if' statements as needed.


                     Example




     if age >= 18 { <───────────────────────── The ‘if’ code
                                                  block
         fmt.Println("You are an adult!")

     } else if age >= 13 { <────────────────── The ‘else-if’
                                                  code block
         fmt.Println("You are teenager!")

     } else { <─────────────────────────────── The ‘else’ code
                                                  block
         fmt.Println("You are a child!")

     }

------------------------------------------------------------------

             The ‘if’ Scoped Variables

• In Go, you have the capability to initialize a new variable directly within the 'if' statement.
• The variable declared in this manner is  scoped exclusively to the 'if' (or 'if-else') block.
• This feature aids in writing more  concise and organized code.


                     Example

even variable                  Condition
Initialize the				      /                                             │
      ▼                          /
								▼
     if even := isEven(age); even {

         fmt.Println("Age is even!")

     }

------------------------------------------------------------------

					Summary

• The 'if' statement in Go enables you to assess a condition and execute specific actions based on the result.
• This construct can be expanded to handle multiple conditions using 'if-else' and  'if-else-if' statements.
• Go provides the flexibility to initialize variables that are scoped exclusively within the 'if' code block.

*/

package main

import "fmt"

// In this lesson we will cover:
// 1. The if statement
// 2. the if-else statment
// 3. the if-else-if statement
// 4. Variable declaration within if.

func main() {

        age := 10


        if age >= 18 {
            fmt.Println("You are an adult!")
        } else if age >= 13 {
            fmt.Println("You are teenager!")
        }else {
            fmt.Println("You are a child!")
        }



        // Variable declaration within if. (Shorthand declaration)
        if even := isEven(age); 
           even == true {
            fmt.Println("Age is even!")
        }



}

func isEven(n int) bool {
    return n&1 == 0
}