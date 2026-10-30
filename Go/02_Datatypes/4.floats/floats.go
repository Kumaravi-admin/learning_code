/*
(02) Go Data Type: float

Type	Description	Range
float32	The set of all IEEE-754 32-bit floating-point numbers	Approximately 1.401298464324817e-45 to 3.4028234663852886e+38 for positive values
float64	The set of all IEEE-754 64-bit floating-point numbers	Approximately 4.9406564584124654e-324 to 1.7976931348623157e+308 for positive values

Note: In Go, float32 uses 32 bits (4 bytes), and float64 uses 64 bits (8 bytes). Both types support positive and negative values, as well as zero.


----------------------------------------------------------------------------------------------------------------------------------

(02) Go Data Type: complex
Type	Description
complex64	Uses float32 for both the real and imaginary parts.
complex128	Uses float64 for both the real and imaginary parts.


Built-in Functions
Function	Description
complex()	Creates a complex number from its real and imaginary parts.
real()	Returns the real part of a complex number.
imag()	Returns the imaginary part of a complex number.

*/

package main

import "fmt"


func main() {

	var smallFloat float32		// float32 zero value is 0 (default value)
	fmt.Println(smallFloat)		// this will print 0 (default value)

	smallFloat = 23.0234324
	fmt.Println(smallFloat)		// this will print 23.023432


	var bigFloat float64		// float64 zero value is 0 (default value)
	fmt.Println(bigFloat)		// this will print 0

	bigFloat = 43.08094938590453
	fmt.Println(bigFloat) 		// this will print 43.08094938590453

	bigFloat = 43.9834829759823759827587443
	fmt.Println(bigFloat)		// this will print 43.983482975982376

	
	var myComplex complex128
	myComplex = complex(bigFloat,bigFloat)

	fmt.Println(myComplex)		// this will print (43.983482975982376+43.983482975982376i)

	var myRealPart, myImaginaryPary float64
	myRealPart = real(myComplex)
	myImaginaryPary = imag(myComplex)
	fmt.Println(myRealPart)
	fmt.Println(myImaginaryPary)

}
