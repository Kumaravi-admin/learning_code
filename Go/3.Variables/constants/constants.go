/*

Constants Rule			Description
Immutability			Values that never change during execution of program or are immutable.
Keyword					Declared with the const keyword.
Compile Time			In Go constants can only hold values that the compiler can infer to be constant at the compile time.
Runtime Restriction		You cannot make a runtime value constant.

----------------------------------------------------------------------------------------------------------------------------------

Constant Value Type		Examples / Details

Booleans				Boolean values i.e true or false
Numeric					Numeric values for example, 1, 100, 122.01, etc
Global Functions		Global function like complex, real, imag, len and cap.
Strings & Runes			A string or rune value for example, "code and learn", 'a'
Expressions				Expression that can be evaluated at compile time for example 1 < 2.

----------------------------------------------------------------------------------------------------------------------------------

Constant Type Rule		Description
Categories				There are two types of constants and they are typed and untyped constants.
Typed Constants			A typed constant can only be assigned to directly to a variable of the same kind.
Untyped Constants		For untyped constant type can be inferred or it fallbacks to default type for the constant value.

----------------------------------------------------------------------------------------------------------------------------------

Category				Details
Similarities			package-level and function-level.
Similarities			Can be grouped using braces ().
Differences				Constant values cannot change but variables can.
Differences				Constants can remain un-used.

*/

package main

import "fmt"

// package lavel constants

const x = 10 	// untype constants
const y int32 = 15 	// type constants

// group const

const (
	n = 10
	s = len("hello World")
)


func main() {

	fmt.Println(n)  // this will print: 10
	fmt.Println(s)	// this will print lent of strint: 11


	var a int
	a = x			// asign untype constants value to int type variable
	fmt.Println(a) 	// this will print: 10

	var b float64
	b = x			// asign untype constants value to float64 type variable
	fmt.Println(b)	// this will print 10

	var c int
	c = int(y)		// assign type constant int32 to int 
	fmt.Println(c)	// this will pirnt: 15


	var d int32		// another wayt ot assign type constant to var variable (both variable type match) 
	d = y
	fmt.Println(d)	// this will aslos print: 15



}