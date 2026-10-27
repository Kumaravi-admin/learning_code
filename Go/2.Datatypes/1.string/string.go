/*
Introduction to Strings in Go

Topic				Description

string				The string data type is used to store text in Go.
String Literals		You can define a string using double quotes (" ") or backticks (``    ).
Examples			"Hello World" or `Hello World`
Zero Value			The zero value of a string is an empty string ("").
Unicode Support		Go supports Unicode characters, symbols, and emojis.
Encoding			Go source code uses UTF-8 encoding for Unicode text.


----------------------------------------------------------------------------------------------------------------------------------

Feature / Topic	Core Concept

Internal Storage			Strings are stored as a byte array internally and are immutable.
Modification				The variable holding the string can be updated, but the string value itself cannot be modified directly.
Encoding					Uses UTF-8 encoding, meaning each Unicode character can take from 1 to 4 bytes in memory.
Character Representation	The rune data type is the best way to represent an individual character in a string.

----------------------------------------------------------------------------------------------------------------------------------

String Type				Syntax					Key Characteristics

Interpreted String		Double quotes ""		Creates a one-line string. Requires the \n escape character to manually introduce a new line break.
Raw String Literal		Backticks ` `			Creates multi-line strings. Preserves actual line breaks and ignores escape characters completely.
Formatted String		String Formatters		Allows dynamically formatting and injecting values into the text structure.

*/

package main

import "fmt"

func main() {

	var myString string		// string zero value is "" (empty sring)
	fmt.Println(myString)	// this will print blank line

	myString = "Hello World"
	fmt.Println(myString)	// this will print: Hello World

	myString = "Hello\nWorld"
	fmt.Println(myString) 	// this will print:  Hello 
							// this will print:  World					 

	myString = `Hello
	
	World`
	fmt.Println(myString)	
	/* this will print: Hello

        World
	*/


	// contact string
	var firstName, lastName string
	firstName = "Anjesh"
	lastName = "Kumar"

	var fullName string
	fullName = firstName + " " + lastName
	fmt.Println(fullName)		// this will Print: Anjesh Kumar


	// another way to contact string
	fmt.Printf("%s %s",firstName, lastName) 	// this will also Print: Anjesh Kumar


}