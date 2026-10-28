/*
                           Break Statement

  •  The 'break' statement exits the current loop prematurely.
  •  When a specified condition is met, it  allows you to leave the loop and
     continue with subsequent code.
  •  It can be utilized with a for loop to  mimic the behavior of a 'do-while' loop.

								Syntax
				┌──────────────────────────────────────┐
				│                                      │
				│  for {                               │
				│       // Code block                  │
				│       condition {                    │
				│            break                     │
				│       }                              │
				│  }                                   │
				│                                      │
				└──────────────────────────────────────┘

------------------------------------------------------------------


                         Continue Statement

  •  The 'continue' statement is employed to bypass a specific iteration of a for loop.
  •  This is useful when you wish to skip executing the loop's code block under certain conditions.

								Syntax
				┌──────────────────────────────────────┐
				│                                      │
				│  for condition {                     │
				│       // Code block                  │
				│       condition {                    │
				│            continue                  │
				│       }                              │
				│  }                                   │
				│                                      │
				└──────────────────────────────────────┘

------------------------------------------------------------------

                                Label

  •  The 'continue' statement by itself only affects the loop in which it's placed.
  •  When dealing with nested loops, 'labels' allow you to continue to an outer loop.
  •  In Go, the usage of 'labels' is infrequent and considered rare.
  •  'Labels' can be used with 'break', 'continue' and 'goto' statements.


								Syntax
				┌──────────────────────────────────────────────┐
				│                                              │
				│  outer:                                      │
				│       for for-clause {                       │
				│            for for-clause {                  │
				│                 if condition {               │
				│                      continue outer          │
				│                 }                            │
				│                 // Code block                │
				│            }                                 │
				│       }                                      │
				│                                              │
				└──────────────────────────────────────────────┘
------------------------------------------------------------------

                                Goto

  •  The 'goto' statement is employed to redirect the flow of code.
  •  In Go, there are limitations on the use  of 'goto':
       •  It is not allowed to jump over  variable declarations.
       •  It is not allowed to jump into inner or same-level code blocks.


								Syntax
				┌──────────────────────────────────────┐
				│                                      │
				│  if condition {                      │
				│       goto end                       │
				│  }                                   │
				│  // skip code                        │
				│  end:                                │
				│  // code                             │
				│                                      │
				└──────────────────────────────────────┘

*/

package main

import "fmt" 

func main() {


	// 1. Break statement
	a := 100
	for {
		fmt.Println("Ouput of Break Statement: ",a)			// output: 100 99 98 97 96 95 94 93 92 91
		a--
		if a == 90 {
			break
		}
	}

	// 2. Continue statment
	for i := 0; i < 5; i++ {
		if i == 2 {
			continue
		}
		fmt.Println("Output of Continue Statement: ",i) 	// 0 1 3 4 (skip 2)
	}

	// 3. Lable for for Loop statement
	outerLoop:
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if i == 1 && j == 1 {
				continue outerLoop
			}
			fmt.Println(i,j)
		}
	}

	// 4. Goto Statement 
	n := 10
	
	if n == 10 {
		goto end
	}
	fmt.Println("Output of n: ",n)
	
	end:
	fmt.Println("The end of the program.")


}

/*			Terminal output


Ouput of Break Statement:  99
Ouput of Break Statement:  98
Ouput of Break Statement:  97
Ouput of Break Statement:  96
Ouput of Break Statement:  95
Ouput of Break Statement:  94
Ouput of Break Statement:  93
Ouput of Break Statement:  92
Ouput of Break Statement:  91

Output of Continue Statement:  0
Output of Continue Statement:  1
Output of Continue Statement:  3
Output of Continue Statement:  4

0 0
0 1
0 2
1 0
2 0
2 1
2 2

The end of the program.

*/