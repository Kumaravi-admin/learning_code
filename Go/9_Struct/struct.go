/*
                    Struct

• Production applications require much more than primitive and composite data
  types provided by Go.
• Go provides struct to allow you to create your own custom data types.
• A struct helps group similar data  together.


                   Example

			type student struct {
				firstName string
				lastName  string
				age       int
				classes   []string
			}
------------------------------------------------------------------

				Zero value for struct

• Zero value for a struct is a struct with  all its fields having zero values for
  their corresponding types.
• An zero value struct be created using a literal declaration: var foo student.
• It can also be created with variable shorthand: foo := student{}.


          Initializing field values

• There are two ways to initialize struct  field values during struct initialization.
• The first is implicit initialization,  where values are provided inside the
  curly braces without field names.
• The second is explicit initialization  when values are provided inside curly
  braces along with field names.

                   Example

				foo := student{
					"anjesh",
					"kumar",
					25,
					[]string{"maths", "english"},
				}

				foo := student{
					firstName: "anjesh",
					lastName:  "kumar",
					age:25,
					subject: []string{"maths", "english"},
				}

------------------------------------------------------------------

               The Dot Notation

• The dot notation is used to read or   write field values of a struct.
• For example, for a student struct named   foo, you can read values with  foo.firstName.
• Similarly, to update the value, you can do foo.firstName = "foobar".


                    Summary

• With structs, you can create your own data type that groups different field   types.
• You can create and initialize struct  field values with implicit or explicit   initialization.
• Dot notation is used to access or modify a single field.
*/

package main

import "fmt"

// In this lesson we well cover:
// 1. Defining a struct
// 2. Initializing a struct.
// 3. Read and writing struct fileds
// 4. anounymous struct.


func main() {
	
	// 1. Defining a struct
	type student struct {
		first_name string
		last_name  string
		age        int
		subjects   []string
	}

	var student1 student
	fmt.Printf("Defining a struct: %+v \n", student1)					// output: {first_name: last_name: age:0 subjects:[]}


	// 2. Initializing a struct 

	// Implicit struct
	student1 = student{"Anjesh","Kumar",25,[]string{"maths","science"}}
	fmt.Printf("Initializing a implicit struct: %+v \n",student1)		// output: {first_name:Anjesh last_name:Kumar age:25 subjects:[maths science]} 

	// Explicit stuct
	student2 := student{
		first_name: "Ankit",
		last_name: "Kumar",
		age:24,
	}


	// Print student2 all fields
	fmt.Printf("Initializing a explicit struct: %+v \n",student2)		// output: {first_name:Ankit last_name:Kumar age:24 subjects:[]}  
	
	// Print stuct specific fields (student fist name )
	fmt.Println("First name of student2:",student2.first_name)			// output: Ankit


	// write struct fields
	student2.subjects = append(student2.subjects,"arts")				
	fmt.Printf("Student2 struct after append values: %+v \n",student2)	// output:  {first_name:Ankit last_name:Kumar age:24 subjects:[arts]} 

	student2.first_name = "Avyansh"
	fmt.Printf("Student2 struct after update name field: %+v\n",student2) // output: {first_name:Avyansh last_name:Kumar age:24 subjects:[arts]}


	// Anonymous struct
	guardian := struct{
		firstName string
		lastName string
	
	} {
		firstName: "Anjesh",
		lastName: "Kumar",
	}

	fmt.Printf("Anonymous struct: %+v\n", guardian)				// output: {firstName:Anjesh lastName:Kumar}
}

/*				Terminal output


Initializing a implicit struct: {first_name:Anjesh last_name:Kumar age:25 subjects:[maths science]} 
Initializing a explicit struct: {first_name:Ankit last_name:Kumar age:24 subjects:[]} 
First name of student2: Ankit
Student2 struct after append values: {first_name:Ankit last_name:Kumar age:24 subjects:[arts]} 
Student2 struct after update name field: {first_name:Avyansh last_name:Kumar age:24 subjects:[arts]}
Anonymous struct: {firstName:Anjesh lastName:Kumar}

*/