/*
                     Strings

• Strings are stored as sequence of bytes.
• You can access each byte of the string   data using index notation.
• The slice expression works with strings.
• Strings in Go are immutable, you cannot   manipulate a string, once created.
• The string, rune, byte types can be  converted to each other.


                     Examples

• Given s is string "Code & Learn":


         Example                  Output

    1.   string(s[0])             "C"
    2.   string(s[:5])            "Code"
    3.   string(s[3:])            "e & Learn"
    4.   string(s[:])             "Code & Learn"
    5.   s[1] = 'a'               'UnassignableOperand' error
------------------------------------------------------------------

                     Example

• Example, we have string s := "朗퐿텟" which  includes unicode
  characters with code points of size more than 1 byte.
• Access a single byte (string(s[1])) will give you invalid output.
• Convert to rune slice, use strings or unicode/utf8 package or use for-range.
• s := "朗퐿텟" convert to rune slice r := []rune(s) and then access each
  character like string(r[2]).

*/

package main

import "fmt"

// In this lesson, we will cover:
// 1. Accessing byte data of a string.
// 2. Trying to update byte data of string.
// 3. Coverting byte, rune to string.
// 3. Slicing a string.

func main() {

	name := "Code & String"

	// Accessing byte data of a string.
	firstLetter := name[0]
	fmt.Println("Byte data of first latter: ",firstLetter)					// output: 67 
	fmt.Println("First letter of string: ",string(firstLetter))				// output: c

	lastLetter := name[len(name)-1]
	fmt.Println("Byte data of lastLatter:",lastLetter)						// output: 103
	fmt.Println("Last letter of string: ",string(lastLetter))				// output: g


	// Convert rune to string
	var b rune = 'x'
	fmt.Println("Output of b before convert rune to string: ",b)			// output: 120
	fmt.Println("outpute of b after convert rune to string: ",string(b))		// output: x


	// Convert byte into string
	var c byte = 'x'
	fmt.Println("Output of c before convert byte to string: ",c)			// output: 120
	fmt.Println("Output of c after convert byte to string: ",string(c))		// output: x


	// Convert string to rune slice
	runeName := []rune(name)
	fmt.Println("Output of rune name: ",runeName)							// output: [67 111 100 101 32 38 32 83 116 114 105 110 103]

	// Convert string to byte slice
	byteName := []rune(name)
	fmt.Println("Output of byte name: ",byteName)							// output: [67 111 100 101 32 38 32 83 116 114 105 110 103]


	
	// Slicing a string
	firstName := name[:]		// copy full string 
	fmt.Println("Output of slicing of string full: ",firstName)				// output: Code & String

	firstName = name[4:6]		// specify range 
	fmt.Println("Output of slicing of string range: ",firstName)			// output:  &


	// Working with UTF Characters
	name = "朗퐿텟"

	unicodeString := []rune(name)	// convert utf charactor to run slice
	fmt.Println("Output of rune unicode string: ",unicodeString)			// output: [26391 54335 53599]
	

	firstCharacter := unicodeString[0]	// access fist byte of rune slice
	fmt.Println("first charactor of unicode string in string :",string(firstCharacter))		// output: 朗


}

/*			Terminal Output

Byte data of first latter:  67
First letter of string:  C
Byte data of lastLatter: 103
Last letter of string:  g
Output of b before convert rune to string:  120
outpute of b after convert rune to string:  x
Output of c before convert byte to string:  120
Output of c after convert byte to string:  x
Output of rune name:  [67 111 100 101 32 38 32 83 116 114 105 110 103]
Output of byte name:  [67 111 100 101 32 38 32 83 116 114 105 110 103]
Output of slicing of string full:  Code & String
Output of slicing of string range:   &
Output of rune unicode string:  [26391 54335 53599]
first charactor of unicode string in string : 朗

*/