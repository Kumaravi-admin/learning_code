package main

import (
	"errors"
	"fmt"
)

// 3. Pointer with struct
type student struct {
	firstName string
	lastName string
}

func main() {


	// 3. student object
	studentName := student{
		firstName: "Anjesh",
		lastName: "Kumar",
	}
	
	previousLastName, err := updateLastName(&studentName,"")
	if err !=nil {
		fmt.Println(err)
	}
	fmt.Println(previousLastName)

}


// 2. Pointer with struct: create update last student name function with handling error

func updateLastName(studentName *student, newLastName string) (*string, error) {
		if newLastName == "" {
			return nil, errors.New("empty new last name")
		}
		previous := studentName.lastName
		studentName.lastName = newLastName
		return &previous,nil

}


/*				Terminal output: 


empty new last name
<nil>

*/ 