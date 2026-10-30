/*
Type	Description	Range
int8	The set of all signed 8-bit integers	-128 to 127
int16	The set of all signed 16-bit integers	-32768 to 32767
int32	The set of all signed 32-bit integers	-2147483648 to 2147483647
int64	The set of all signed 64-bit integers	-9223372036854775808 to 9223372036854775807

--------------------------------------------------------------------------------------------

Topic	Description
uint	The uint data type is used to store non-negative integers (without fractions).
Examples	1, 2349087, 12, 23423545345
Types	uint8, uint16, uint32, uint64, etc.
Size and Range	Different uint types have different sizes in memory and different ranges of numbers they can store.
Range	All bits are used to store the value, so unsigned integers have a greater maximum range than signed integers of the same bit size.

----------------------------------------------------------------------------------------------------------------------------------


Go uint Types and Their Ranges

Type	Size	Range
uint8	8 bits (1 byte)	0 to 255
uint16	16 bits (2 bytes)	0 to 65,535
uint32	32 bits (4 bytes)	0 to 4,294,967,295
uint64	64 bits (8 bytes)	0 to 18,446,744,073,709,551,615

Note: Go's uint size is either 32 bits or 64 bits, depending on the target architecture. Also, uint cannot store negative numbers.

----------------------------------------------------------------------------------------------------------------------------------

Type	Description
byte	Alias for uint8
rune	Alias for int32
uint	Either 32 or 64 bits, depending on the system architecture
int	Either 32 or 64 bits, depending on the system architecture
*/

package main

import "fmt"

func main() {

	var myInt int		// int zero value is 0 (default value)
	fmt.Println(myInt) 	// this will print 0


	myInt = 42
	fmt.Println(myInt) 	// this will print 42



	var smallPositiveInt uint8		// uint zero value is 0 (default value)
	smallPositiveInt = 10	
	fmt.Println(smallPositiveInt)	// this will print 10


}