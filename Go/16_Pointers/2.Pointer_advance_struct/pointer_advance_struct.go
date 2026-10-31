package main

import "fmt"

// Pointer with struct
type student struct {
	firstName string
	lastName string
}

func main() {


	// student object
	studentName := student{
		firstName: "Anjesh",
		lastName: "Kumar",
	}
	
	// student last name before update
	fmt.Println(studentName)				// output: {Anjesh Ravi}

	// student last name after update
	updateLastName(&studentName,"Ravi")
	fmt.Println(studentName)				// output: {Anjesh Ravi}


	// Call Pretty print function (if we do note use point when function call each time fuction make copy of struc)
	prettyPrintStudent(&studentName)

	// output: First name: Anjesh
	//		   Last name: Ravi
}	


// 1. Pointer with struct: create update last student name function

func updateLastName(studentName *student, newLastName string) {
	studentName.lastName = newLastName

}

// 2. Function that use pointer to access struct 
func prettyPrintStudent(studentName *student) {
	fmt.Printf("First name: %s\nLast name: %s\n", studentName.firstName,studentName.lastName)
}

/*				Terminal output


{Anjesh Kumar}
{Anjesh Ravi}
First name: Anjesh
Last name: Ravi

*/