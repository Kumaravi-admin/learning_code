/*

                    Slices

• Slices are wrappers around arrays.
• They manage the underlying arrays for you.
• They help you overcome the limitations of arrays.
• Length is not attached to the type of the slice.
• Slices have length and capacity.
------------------------------------------------------------------

                     Example

               - - - - - Slice - - - - -
              (                         )
              v                         v

                      Length: 3
              <------------------------->
                                Capacity: 5
              <------------------------------------->

             [  10  ][  30  ][   5  ][      ][      ]
                 0       1       2       3       4


• Length of a slice is the length of  elements stored in the slice (use len).
• Capacity is total size of the underlying array (use cap).
------------------------------------------------------------------

               Behind The Scenes

• When you create a slice of certain length underlying array is usually
  created of double the length.
• When you reach the capacity of a slice, it will create a array of 50 to 25 %
  larger than the existing array and copy values from old to new array.
------------------------------------------------------------------

            Declaring Slices in Go

• Declaring slice:
    o var a []int
    o a := []int{1, 2, 3, 4, 5}
    o a := []int{5, 2:10, 50}
    o a := make([]int, 0)
------------------------------------------------------------------

              The make functio

• The universal make function can be used  to create slices.
• It takes type of slice as first   argument, length as second and capacity
  as the optional third argument.
• Example:
    o a := make([]int, 5)
    o a := make([]string, 5, 10)
    o a := make([]int, 0, 10)
------------------------------------------------------------------

             The append function

• The append universal function can be used to append elements to a slice.
• It is an error to not assign the return value of append.
• Example:
    o a = append(a, 5)
    o a = append(a, 4, 2, 5, 6)
    o a = append(a, b...)
------------------------------------------------------------------

           Similarities with Arrays

• Slices like arrays use index to fetch or  manipulate value stored at given index.
• Both arrays and slices are zero-indexed.
• You cannot have index -1 for a slice like array and like array cannot access
  element greater than length - 1.
------------------------------------------------------------------

              Empty & Nil Slice

• nil in Go is different from null of other languages.
• nil means a lack of value and type.
• nil is the zero value for slice.
• A nil slice has no underlying array.
• A empty slice has underlying with no
  values stored.
• A nil slice: var a []int
• A empty slice: a := []int{}
------------------------------------------------------------------

                    Slicing

• Go allows you to slice a slice or array to create a new slice.
• Range of indices are used to slice a  slice or array.
• Example: b := a[1:10]
• Range specified is a half-open range  which includes the first element, but excludes the last one.

                    Examples

• Given a is a slice or array here are few
  different examples:

    Example    		  Actual Values Sliced

	b := a[2:5]  	  index 2 to index 4
	b := a[:5]        index 0 to index 4
	b := a[3:]        index 3 to last index
	b := a[:]         first index to last index
------------------------------------------------------------------

               Careful with Slicing

• New slice still refers to the same  underlying array the original slice uses.
• Manipulating or appending values in the  child slice can cause a ripple effect.
• Be careful when slicing a slice or array to create a new slice when you want to
  manipulate the child slice.

                 Passing Capacity

• The slice expression accepts a third value which is the capacity of the new slice.
• Passing capacity equal to the length of  slice will make sure new underlying is
  created when you append.
• Example: b := a[0:10:10]
• Manipulating slice directly would still lead to the same problem.
------------------------------------------------------------------

                The copy Function

• copy(dst, src []Type) int
• Copy takes the destination slice and the source slice, returning the number of elements copied.
• It will copy as many elements it can.
• First create a empty slice of given length and copy of the values.

                    Examples


    Example                 Description
    copy(a, b)         		copy elements from slice b to a
  	copy(a, b[1:10]) 		copy range of elements from slice b to a
    copy(a, b[:])     		copy element from array b to a
------------------------------------------------------------------

*/

package main

import "fmt"

// In this lesson we will cover:
// 1. Creating a slice.
// 2. The make function.
// 3. The append function.
// 4. nil and empty slice.
// 5. Different slice expressions.
// 6. The problemwith slicing.
// 7. Passing a capacity to the slice expression.
// 8. the copy function


func main() {

	// Nil slice
	var a []int
	fmt.Println("Output of nil slice a: ",a == nil)			// Output: true
	
	// Empty slice
	a = []int{}
	fmt.Println("Output of empty slice a:", a == nil)		// output: false


	// Different way of declaring and inializing slices
	b := []int{1,2,3,4,5}
	fmt.Println("Output of slice b: ", b)					// output: [1 2 3 4 5]

	b = []int{5, 2: 10, 50}
	fmt.Println("Output of slice b: ", b)					// output: [5 0 10 50]


	// Make fuction
	c := make([]int,5)
	fmt.Println("Output of slice c lenth: ", len(c))		// output: 5
	fmt.Println("Output of slice c capacity: ", cap(c))		// output: 5

	d := make([]string,5,10)
	fmt.Println("Output of slice d: ", d)					// output: [    ]
	fmt.Println("Output of slice d lenth: ", len(d))		// output: 5
	fmt.Println("Output of slice d capacity: ", cap(d))		// output: 10


	// Append function
	e := make([]int,0)
	fmt.Println("Output of slice e: ",e)					// []

	e = append(e,10)
	fmt.Println("Output of slice e after append: ",e)		// output: [10]

	e = append(e,10,15,20)
	fmt.Println("Output of slice e after append: ",e)		// output: [10 10 15 20]

	e = append(e,10,15,20)
	fmt.Println("Output of slice e after append: ",e)		// output: [10 10 15 20 10 15 20]



	// Append slice using spred operator
	f := []int{1, 2, 3, 4}
	g := []int{10, 20, 30}
	h := append(f, g...)
	fmt.Println("Output of slice h after append: ", h)		// output: [1 2 3 4 10 20 30]


  // Create slice from other slice value passing index [start index:end index]
  i := []int{1, 2, 3, 4, 5, 6}

  j := i[2:5]   // value from 2 index to 5 index
  fmt.Println("Output of slice j: ",j)                  // output: [3 4 5]

  j = i[:] // capy all value from i slice
  fmt.Println("Output of slice j: ",j)                  // output: [1 2 3 4 5 6] 

  j = i[2:] // all value after 2 index
  fmt.Println("Output of slice j: ",j)                  // output: [3 4 5 6]

  j = i[:3] // value from start to till index 3
  fmt.Println("Output of slice j: ",j)                  // output: [1 2 3]

  j = i[2:6:6]   //[start index:end index:capacity]
  j = append(j,100)
  fmt.Println("Output of slice j after append: ",j)     // output:  [3 4 5 6 100]



  // Copy function
  j = make([]int,6)    // make(slice,capacity)
  copy(j,i)
  fmt.Println("Output of copy function i to j: ",j)     // output:   [1 2 3 4 5 6]


   j = make([]int,3)    // copy till index 3
  copy(j,i)
  fmt.Println("Output of copy function i to j: ",j)     // output:   [1 2 3]

  

  // Copy Value form array to slice
  k := [6]int{1, 2, 3, 4, 5, 6}
  l := make([]int,6)

  copy(l,k[:])    // copy all value from array to slice 
  fmt.Println("Output of copy function k to l: ",l)     // output:   [1 2 3 4 5 6]

  m := make([]int,6)
  copy(m,k[2:])   // copy all array value after index 2 to slice
  fmt.Println("Output of copy function k to m: ",m)     // output:    [3 4 5 6 0 0]


}

/*			Terminal Output

Output of nil slice a:  true
Output of empty slice a: false
Output of slice b:  [1 2 3 4 5]
Output of slice b:  [5 0 10 50]
Output of slice c lenth:  5
Output of slice c capacity:  5
Output of slice d:  [    ]
Output of slice d lenth:  5
Output of slice d capacity:  10
Output of slice e:  []
Output of slice e after append:  [10]
Output of slice e after append:  [10 10 15 20]
Output of slice e after append:  [10 10 15 20 10 15 20]
Output of slice h after append:  [1 2 3 4 10 20 30]

*/