/*

Bits and Bitwise Operators		Details
Bits Definition					Bits are the smallest unit of computer memory and the language of computers.
Bitwise Operators				The bitwise operators are used to perform low-level bit manipulation operations.
Examples						cryptography algorithms, system programming, network programming, device drivers, etc.

----------------------------------------------------------------------------------------------------------------------------------

Bitwise OR Operator

A	B	OR(|)
0	0	0
0	1	1
1	0	1
1	1	1


            Bitwise OR Operation

Example of bitwise OR operation written as a | b.


              Binary          Decimal

           1 1 0 0 1 0 1 0      202

           1 0 0 0 1 1 0 1      141
           ---------------
  Result:  1 1 0 0 1 1 1 1      207


----------------------------------------------------------------------------------------------------------------------------------

Bitwise AND Operator

A	B	AND (&)
0	0	0
0	1	0
1	0	0
1	1	1


           Bitwise AND Operation

Example of bitwise AND operation written as a & b.


              Binary          Decimal

           1 1 0 0 1 0 1 0      202

           1 0 0 0 1 1 0 1      141
           ---------------
  Result:  1 0 0 0 1 0 0 0      136


----------------------------------------------------------------------------------------------------------------------------------

Bitwise XOR Operatorj

A	B	XOR (^)
0	0	0
0	1	1
1	0	1
1	1	0

The exact text extracted from the image is:
• Sets the resultant bit if one bit is set and the other is not.
• One bit is 1 and the other is 0.


       Bitwise XOR Operation

Example of bitwise XOR operation written as a ^ b.


              Binary          Decimal

           1 1 0 0 1 0 1 0      202

           1 0 0 0 1 1 0 1      141
           ---------------
  Result:  0 1 0 0 0 1 1 1       71

----------------------------------------------------------------------------------------------------------------------------------

Bitwise AND NOT Operator

A	B	AND NOT (&^)
0	0	0
0	1	0
1	0	1
1	1	0

• Sets the resultant bit, if the bit is set
  in the first number but not in second,
  order matters.


Example of bitwise AND NOT operation written as
                      a &^ b.

              Binary          Decimal

           1 1 0 0 1 0 1 0      202

           1 0 0 0 1 1 0 1      141
           ---------------
  Result:  0 1 0 0 0 0 1 0       66

----------------------------------------------------------------------------------------------------------------------------------

Bitwise INVERT Operator

A	INVERT (^)
0	1
1	0


• Unary operator that inverts the bits.
• Sets the bit that was not set and un-sets
  the bit that was set.
• Also known as the NOT operator.


          Bitwise INVERT Operation

Example of bitwise INVERT operation written as
                       ^a.


              Binary          Decimal

           1 1 0 0 1 0 1 0      202
           ---------------
  Result:  0 0 1 1 0 1 0 1       53

----------------------------------------------------------------------------------------------------------------------------------

           Binary Shift Operators

• Bitwise shift operators are used to remove or
  append bits to a integer value.
• There are two shift operators, the right shift
  operator denoted by >> and the left shift
  operator denoted by <<.
• Used for low level bit manipulation like
  generating a bit mask.


             Right Shift Operator

• The right shift operator remove bits from
  number or moves the bits to the right.
• a >> i operation removes i bits from the number
  a shifting i bits from the number to right.
• Operations 1 1 1 1 >> 2 will give you the
  result 0 0 1 1.
• Right shifting i bits of the number a will
  result in a number that is equal to a divided
  by 2^i.


               Left Shift Operator

• The left shift operator appends zero bits to a
  number.
• The a << i operation appends i zero bits to the
  number a.
• Operations 1 1 1 1 << 2 will give you the
  result 1 1 1 1 0 0.
• The left shift operation effectively multiplies
  the number by 2^i (where i is the number of
  bits shifted)


               More Interesting Facts

• 1 << i will result in the a single bit being
  set in the i^th position while rest of the bits
  are zero.
• The operation a | (1<<i) will set the i^th bit of
  a to 1.
• The operation a &^ (1<<i) will set the i^th bit
  of a to 0 (bit clear).
• The operation a ^ (1<<i) will inverts the i^th
  bit of a.



                      Examples

                                              Result

Binary:  0 0 0 1 << 3                   =     1 0 0 0

Binary:  0 1 0 1 |  (1 << 1)            =     0 1 1 1

Binary:  0 1 1 1 &^ (1 << 1)            =     0 1 0 1

Binary:  0 1 0 1 ^  (1 << 0)            =     0 1 0 0

----------------------------------------------------------------------------------------------------------------------------------

               Right Shift Example

Example of bitwise >> operation written
                     as a >> i.

                                          Result

Binary:     1 0 1 0 0 >> 2          =      1 0 1

Decimal:          20                =        5

----------------------------------------------------------------------------------------------------------------------------------

               Left Shift Example

Example of bitwise << operation written
                     as a << i.

                                          Result

Binary:       1 0 1 << 2            =    1 0 1 0 0

Decimal:          5                 =        20

----------------------------------------------------------------------------------------------------------------------------------





*/

package main

import (
	"fmt"
	"strconv"
)

func main() {
	
	var a uint8 = 202
	var b uint8 = 141
	var x uint8 = 20 // 10100 in bit


	// The OR operator
	c := a | b
	fmt.Println("--------Resulet of: OR Operation--------")
	fmt.Println("c: ",c)
	fmt.Println("a: ",strconv.FormatUint(uint64(a),2))
	fmt.Println("b: ",strconv.FormatUint(uint64(b),2))
	fmt.Println("c: ",strconv.FormatUint(uint64(c),2))



	// The AND operator
	d := a & b
	fmt.Println("--------Resulet of: AND Operation--------")
	fmt.Println("d: ",d)
	fmt.Println("a: ",strconv.FormatUint(uint64(a),2))
	fmt.Println("b: ",strconv.FormatUint(uint64(b),2))
	fmt.Println("d: ",strconv.FormatUint(uint64(d),2))



	// The XOR operator
	e := a ^ b
	fmt.Println("--------Resulet of: XOR Operation--------")
	fmt.Println("e: ",e)
	fmt.Println("a: ",strconv.FormatUint(uint64(a),2))
	fmt.Println("b: ",strconv.FormatUint(uint64(b),2))
	fmt.Println("e: ",strconv.FormatUint(uint64(e),2))


	// The AND NOT or Bit Clear operator
	f := a &^ b
	fmt.Println("--------Resulet of: AND NOT Operation--------")
	fmt.Println("f: ",f)
	fmt.Println("a: ",strconv.FormatUint(uint64(a),2))
	fmt.Println("b: ",strconv.FormatUint(uint64(b),2))
	fmt.Println("f: ",strconv.FormatUint(uint64(f),2))


	// The INVERT or the NOT operator
	g := ^a
	fmt.Println("--------Resulet of: INVERT or the NOT Operation--------")
	fmt.Println("g: ",g)
	fmt.Println("a: ",strconv.FormatUint(uint64(a),2))
	fmt.Println("g: ",strconv.FormatUint(uint64(g),2))


	// Right shift operation
	fmt.Println("--------Resulet of: Right shift operation--------")
	fmt.Println(strconv.FormatUint(uint64(x),2))
	x = x >> 2 // Remove the last to bits -> 101 or 5
	fmt.Println(strconv.FormatUint(uint64(x),2))


	// Left shift operation
	fmt.Println("--------Resulet of: Left shift operation--------")
	fmt.Println(strconv.FormatUint(uint64(x),2))
	x = x << 2	// Appedn two zero bits --> 10100
	fmt.Println(strconv.FormatUint(uint64(x),2))


	// Create a bit mask
	fmt.Println("--------Resulet of: Create a bit mask--------")
	x = 1
	fmt.Println(strconv.FormatUint(uint64(x),2))
	x = x << 3	// bitmask where 3rd bit is set to 1
	fmt.Println(strconv.FormatUint(uint64(x),2))


	// Set a bit a given possiton
	fmt.Println("--------Resulet of: Set a bit a given possiton--------")
    fmt.Println(strconv.FormatUint(uint64(x),2))
	x = 5 | (1 << 1)	// 101 -> 111 set bit a position 1 to 1
	fmt.Println(strconv.FormatUint(uint64(x),2))


	// Unset bit at given possiton
	fmt.Println("--------Resulet of: Unset bit at given possiton--------")
	fmt.Println(strconv.FormatUint(uint64(x),2))
	x = x &^ (1<<1)		// 111 -> 101 clear bit a position 1
	fmt.Println(strconv.FormatUint(uint64(x),2))


	// Toggle a bit at a given possiton
	fmt.Println("--------Resulet of: Toggle a bit at a given possiton--------")
	fmt.Println(strconv.FormatUint(uint64(x),2))
	x = x ^ (1 << 0)	// 101 -> 100
	fmt.Println(strconv.FormatUint(uint64(x),2))


	
}		

/*			Tearmina output

--------Resulet of: OR Operation--------
c:  207
a:  11001010
b:  10001101
c:  11001111
--------Resulet of: AND Operation--------
d:  136
a:  11001010
b:  10001101
d:  10001000
--------Resulet of: XOR Operation--------
e:  71
a:  11001010
b:  10001101
e:  1000111
--------Resulet of: AND NOT Operation--------
f:  66
a:  11001010
b:  10001101
f:  1000010
--------Resulet of: INVERT or the NOT Operation--------
g:  53
a:  11001010
g:  110101
--------Resulet of: Right shift operation--------
10100
101
--------Resulet of: Left shift operation--------
101
10100
--------Resulet of: Create a bit mask--------
1
1000
--------Resulet of: Set a bit a given possiton--------
1000
111
--------Resulet of: Unset bit at given possiton--------
111
101
--------Resulet of: Toggle a bit at a given possiton--------
101
100


*/