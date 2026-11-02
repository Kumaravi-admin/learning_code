/*
                    Call By Value

• Call by value: Pass the value of the parameter.
• Inside function a copy of parameter is created.
• Any changes made inside the function do not have any effect on the actual value.
• In most languages call by value is used for most primitive data types.


                   Call By Reference

• Call by value: Pass the memory address of the  parameter.
• Inside the function you have access to the actual memory where data is stored.
• Any change made inside the function are visible on the actual variable.
• In most languages call by reference is used for complex data structure like classes and arrays.


       			   Go: Strictly call by value


• Go is strictly call by value.
• Even when you pass a pointer inside the function a copy of pointer is created.
• Maps and slices are curious cases in Go.
• Maps and slices are internally implemented as data structures that hold memory addresses.

                       Maps

• Maps are part of the reference types in Go.
• Data structure that holds the maps has a reference of memory where data is stored.
• When passing to function, a copy of the structure defining the map is created.
• The copy also includes the reference to memory  when data is stored.


                      Slices

• Slices are also reference data types along with maps and pointers.
• A slice is data structure that holds length, capacity and pointer to memory where data is stored.
• Inside a function again a copy of slice structure is created.
• Directly changing value updated the original slice but append is different.

------------------------------------------------------------------

                     Illustration

          Primitive    Pointers      Maps          Slices

                                                   Len: 2
  Copy    5            Addr:         Key,          Cap: 4
                       0x5431E       Values:       Values:
                                     0x5431E       0x3498D
                                        ▲
                                        │
  Value   5            Addr:         Key,          Len: 2
                       0x5431E       Values:       Cap: 4
                                     0x5431E       Values:
                                                   0x3498D
------------------------------------------------------------------



*/

package main

import "fmt"

func main() {

	// 1. Slice
	x := []int{1, 2}
	fmt.Println("Output before call somethingslice function: ",x)		// output:  [1 2]

	somethingSlice(x)
	fmt.Println("Output after call somethingslice function: ",x)		// output:  [100 2]



	// 2. Map
	myMap := map[int]string{
		0: "Learn",
		1: "Code",
	}
	fmt.Println("Output before calling somethingMap function: ",myMap)

	somethingMap(myMap)
	fmt.Println("Output after calling somthingMap function: ", myMap)

}
 
// 1. Slice
func somethingSlice(x []int ) {
	x[0] = 100
}

// 2. Map
func somethingMap(m map[int]string) {
	delete(m,0)

}

/*		Terminal output


Output after call somethingslice function:  [100 2]
Output before calling somethingMap function:  map[0:Learn 1:Code]
Output after calling somthingMap function:  map[1:Code]


*/