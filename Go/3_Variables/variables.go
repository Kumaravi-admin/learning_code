/*

Variable Rule		Description

Purpose				Variables are declared to store values that can change throughout.
Scope				There can be package level variables or function level variables.
Type Declaration	Variable type declaration can be explicit or implicit.
Grouping			Variables declaration can be grouped with braces ().
Inline Declaration	Multiple variables can be declared in single line.

----------------------------------------------------------------------------------------------------------------------------------


Type Assignment				Description / Example

Explicit Type Assignment	The type is assigned in declaration.
Explicit Examples			var age int or var age int = 10
Implicit Type Assignment	Type is inferred from right side value.
Implicit Rules				The default type of the value on right is used to assign type to the variable.
Implicit Example			var age = 10

----------------------------------------------------------------------------------------------------------------------------------

Short Variable Declaration	Description / Example
Omitting Keywords			The keyword var can be omitted along with the type which is inferred from value on right in variable declaration.
Multiple Variables			Can be used to declare multiple variables as long as one variable on the left is new.
Example						age := 10

----------------------------------------------------------------------------------------------------------------------------------

Examples

var age int
var age int = 10
var age = 10
var age, name = 10, "Code & Learn"
age := 25
age, name := 10, "Code & Learn"


*/

package main

import "fmt"

// (01) Package lavel variable
var myInt = 10


func main() {

	fmt.Println(myInt)	// this will print: 10


	// (02) Group variable
	var (
		age int = 26
		name string = "Anjesh Kumar"
	)
	fmt.Println(age)	// this will print: 26
	fmt.Println(name)	// this will print: Anjesh Kumar



	// (03) Explicit type assigment
	var pinCode int = 110001
	var city string = "New Delhi"

	fmt.Println(pinCode)	// this will print: 110001
	fmt.Println(city)		// this will print: New Delhi


	// implicit type assignment
	var studentAge = 25
	var studentCity = "Mumbai"
	fmt.Println(studentAge)		// this will print: 25
	fmt.Println(studentCity)	// this will print: Mumbai


	// Mulltiple variable declaration
	var year, month = 2026, "Octuber"
	fmt.Println(year,month)		// this will print: 2026 Octuber


	// short hand declaration
	firstName := "Anjesh"
	lastName := "Kumar"
	birthYear := 2001

	fmt.Printf("%s\n%s\n%v\n",firstName,lastName,birthYear)

	/* this will print
	Anjesh
	Kumar
	2001
	*/


	// // All Examples:

	// var apartmentNumber int // zero value ot the variable
	// var apartmentNumber int = 2000
	// var apartmentNumber = 2000
	// var apartmentNumber,streetName = 2000,"Bayshore"
	// apartmentNumber := 2000
	// apartmentNumber,strestreetName := 2000,"Bayshore"
}