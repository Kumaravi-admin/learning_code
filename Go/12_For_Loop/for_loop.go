/*
                               Loops

  •  Loops in programming serve to repetitively execute a block of code.
  •  Throughout history, programming languages have featured constructs such
     as the for loop, while loop, and  do-while loop.
  •  In Go, the singular type of loop available is the for loop.


                          Types of for loop

  •  There are three kind of for loops in Go.
       •  For with a condition.
       •  For with for clause.
       •  For with range clause.
------------------------------------------------------------------

                       For loop with condition

  •  The 'for with condition' loop type  requires a single condition.
  •  It will continue executing as long as the condition remains true.
  •  If the condition is omitted, the loop will default to true, resulting in an infinite loop.
  •  To exit an infinite loop, we can use the 'break' statement.

									Syntax
						┌──────────────────────────────┐
						│                              │
						│  for condition {             │
						│       // Code block.         │
						│  }                           │
						│                              │
						└──────────────────────────────┘
------------------------------------------------------------------


                        For Loop with for clause

  •  This is the most standard form of a for  loop, present in nearly all programming
     languages that support this construct.
  •  It comprises three components separated by semicolons:

       •  Initialization.
       •  Condition.
       •  Post Statement.
  •  Any of these parts can be left empty.

										Syntax
						┌──────────────────────────────────────┐
						│                                      │
						│  for init; condition; post {         │
						│       // Code block.                 │
						│  }                                   │
						│                                      │
						└──────────────────────────────────────┘
------------------------------------------------------------------

                       For Loop with range clause

  •  A 'for loop with a range clause'  iterates through elements of an array,
     slice, string, map, channel, or any  derived type from these composite structures.
  •  It yields two iteration variables for  arrays, slices, strings, and maps, but
     only a single iteration variable for channels.
  •  When used with an array or slice, range  provides the index of the element as the
     first iteration variable, and the value  as the second.
  •  When applied to a string, range iterates  over the Unicode code points within the  string.
  •  In a map, the first value represents the key, while the second value corresponds
     to the associated value for that key.
  •  However, it's important to note that the  order of iteration is not guaranteed.
  •  If the map is modified during  iterations, the presence of newly added
     or deleted values may or may not be observed.



										Syntax
						┌──────────────────────────────────────┐
						│                                      │
						│  for val1, val2 := range x {         │
						│       // Code block.                 │
						│  }                                   │
						│                                      │
						└──────────────────────────────────────┘

------------------------------------------------------------------


.*/

package main

import "fmt"

func main() {

	// 1. For loop with condition
	a := 10

	for a > 0 {
		fmt.Println("Output of a: ",a)		// output : 10 9 8 7 6 5 4 3 2 1
		a--
	}


	// 2. For loop with break statement
	b := 10 

	for {
		fmt.Println("Output of b: ",b)		// output : 10 9 8 7 6 5 4 3 2 1
		b--
		if b == 0 {
			break
		}
	}


	// 3. For Loop with for clause
	for i := 0; i < 10; i++ {
		fmt.Println("Output of i: ", i)		// output: 0 1 2 3 4 5 6 7 8 9
	}


	// 4. For Loop with range clause
	n := []int{1, 2, 3, 4, 5}

	for index, value := range n{
		fmt.Printf("Ouput of n: index: %d, value: %d\n", index,value)		// output: 1 2 3 4 5
	}


	// 5. For Loop with map
	studentsAge := map[string]int {
		"Anjesh":26,
		"Avyansh":3,
		"Ankit": 25,
	}

	for name, age := range studentsAge {
		fmt.Printf("Output of student age: %s is %d year old.\n", name,age) 	// Anjesh is 26 year old. Avyansh is 3 year old. Ankit is 25 year old.
	}

	// 6. For Loop with unicode string
	s := "朗퐿텟"
	for index, runeValue := range s {
		fmt.Printf("Output of s: Index: %d, Rune Value : %c\n", index, runeValue) 	// output: 朗 퐿 텟
	}

}


/*				Terminal Output

Output of a:  10
Output of a:  9
Output of a:  8
Output of a:  7
Output of a:  6
Output of a:  5
Output of a:  4
Output of a:  3
Output of a:  2
Output of a:  1

Output of b:  10
Output of b:  9
Output of b:  8
Output of b:  7
Output of b:  6
Output of b:  5
Output of b:  4
Output of b:  3
Output of b:  2
Output of b:  1

Output of i:  0
Output of i:  1
Output of i:  2
Output of i:  3
Output of i:  4
Output of i:  5
Output of i:  6
Output of i:  7
Output of i:  8
Output of i:  9

Ouput of n: index: 0, value: 1
Ouput of n: index: 1, value: 2
Ouput of n: index: 2, value: 3
Ouput of n: index: 3, value: 4
Ouput of n: index: 4, value: 5

Output of student age: Ankit is 25 year old.
Output of student age: Anjesh is 26 year old.
Output of student age: Avyansh is 3 year old.

Output of s: Index: 0, Rune Value : 朗
Output of s: Index: 3, Rune Value : 퐿
Output of s: Index: 6, Rune Value : 텟


*/