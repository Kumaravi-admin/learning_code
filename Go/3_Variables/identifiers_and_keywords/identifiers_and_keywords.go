/*

Identifiers Rule		Description / Examples

Definition				Name that you give to identify entities like variables, constants, types, functions, etc.
Composition				Formed of one or more Unicode letters and digits.
First Character Rule	As per language spec the first character in an identifier must be a letter.
Underscore Rule			Underscore is considered a letter.
Examples				myInt, _myString, \(\alpha\beta\), MyName

----------------------------------------------------------------------------------------------------------------------------------

Predeclared Identifier Category		Items

Types								any bool byte comparable complex64 complex128 error float32 float64 int int8 int16 int32 int64 rune string uint uint8 uint16 uint32 uint64 uintptr
Constants							true false iota
Zero value							nil
Functions							append cap close complex copy delete imag len make new panic print println real recover

----------------------------------------------------------------------------------------------------------------------------------

Keywords Feature	Rules / Items

Meaning				Keywords have a specific meaning and function.
Restriction			Cannot be used as identifiers.
List of keywords	break, default, func, interface, select, case, defer, go, map, struct, chan, else, goto, package, switch, const, fallthrough, if, range, type, var



*/

package main

import "fmt"


func main() {


	var myInt int = 10 		//
	fmt.Println(myInt)

}