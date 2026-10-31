package main

import "fmt"

// 1. Increment function without pointer
func increment(x int) {
	x++
}

// 2. incremet function with pointer
func increment2(x *int) {
	*x++
}

func main() {

	a := 10
	increment(a)		
	fmt.Println(a)							// output: 10

	
	increment2(&a)
	fmt.Println(a)							// output: 11

}


/*			Terminal output:

10
11

*/
