/*

                     Arrays

• Arrays are data structures used to store
  data in contiguous blocks of memory.
• Each piece of data of an array is called
  the element of the array.
• Each element of the array needs to be of
  the same type.
• We use the index of the element in the
  array to access the element.
• Arrays are zero indexed.

----------------------------------------------------------------------------------------------------------------------------------

                     Example


Elements     →     |  10  |  30  |  5  |  50  |  9  |
Index        →        0      1      2     3      4


• Integer array with length of 5.
• Accessing elements beyond the range (0
  to len - 1) of arrays lead to an out of
  range error (Go code panics).


----------------------------------------------------------------------------------------------------------------------------------

                Benefits of Arrays


• Arrays are very efficient to fetch
  elements from memory.
• Used to store collection of data, such a
  list of numbers, etc.
• Uses include search and sort algorithms,
  buffers, etc.
----------------------------------------------------------------------------------------------------------------------------------

               Declaring Arrays in Go

	• Declaring array:
    o var a [5]int
    o var a = [5]int{1, 2, 3, 4, 5}
    o var a = [5]int{5, 2:10, 50}
    o var a = [...]int{1, 4, 5}


----------------------------------------------------------------------------------------------------------------------------------

                    Limitations

• Once declared the size cannot be
  altered.
• The size or length of the array are
  stored in the type of the array.
• You cannot use a variable to declare the
  size of the array.
• Most cases slices are favoured over
  arrays.

*/

package main

import "fmt"

// In this lesson we will cover:
// 1. Different ways of array declaration.
// 2. The len function.
// 3. Accessing array elements by their index.
// 4. Two dimensional array.
// 5. Update the array

func main() {

	// Array declaration basic form
	var a  [5]int
	fmt.Println("Output of a: ",a)				// output: [0 0 0 0 0]


	// Array declaration with elements
	b := [5]int{1,2,3,4,5}
	fmt.Println("Output of b: ",b)				// output: [1 2 3 4 5]


	// Sparse array declaration
	c := [5]int{5,2:10,50}
	fmt.Println("Output of c: ",c)				// output: [5 0 10 50 0]


	// Implicit lenght declaration
	d := [...]int{3,4,1,2}
	fmt.Println("Output of d: ",d)				// output: [3 4 1 2]

	
	// Len function to calculate the lenght of the array
	e := [...]int{3,4,1,2}
	fmt.Println("Output of e: ",len(e))			// output: 4


	// Accessing array elements
	f := [...]int{3,4,1,2}
	g := f[0]
	fmt.Println("Output of g: ",g)				// output: 3
	fmt.Println("Output of f: ",f[2])			// output: 1


	// Two dimensional array
	h := [2][2]int{{1,2},{3,4}}
	fmt.Println("Output of h: ",h)				// output: [[1 2] [3 4]]
	fmt.Println("Output of h: ",h[1][1])		// output: 4
	

	// Update the array value
	i := [...]int{3,4,1,2}
	i[2] = 10
	fmt.Println("Output of i: ",i)				// output: [3 4 10 2]
}


/*			Terminal Output

Output of a:  [0 0 0 0 0]
Output of b:  [1 2 3 4 5]
Output of c:  [5 0 10 50 0]
Output of d:  [3 4 1 2]
Output of e:  4
Output of g:  3
Output of f:  1
Output of h:  [[1 2] [3 4]]
Output of h:  4
Output of i:  [3 4 10 2]



*/