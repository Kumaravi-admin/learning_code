package main

import "fmt"

func main() {

	a := 10
	b := 5
	c := 15.50


	// Addition opration	
	sum := a + b		
	fmt.Println("Addition Operater: ",sum)	


	// you can also concat string using + sign
	firstName := "Anjesh"
	lastName := "Kumar"
	fullName := firstName +" "+ lastName	
	fmt.Println("Full name: ",fullName)					

	// Subtraction operation
	subtraction := a - b		
	fmt.Println("Subtraction: ",subtraction)

	// Multiplication Operation 
	multiplication := a * b			
	fmt.Println("Multipication: ",multiplication)	

	// Division operator
	division := a / b
	fmt.Println("Division: ", division)

	// Remainder operator
	remainder := a % b			
	fmt.Println("Remainder: ",remainder)	

	//Type Conversion
	newSum := a + int(c)
	fmt.Println("New Sum: ",newSum)

	// Computation shorthand
	n := 10
	
	n += 5
	fmt.Println("n += : ",n)

	n -= 5
	fmt.Println("n -= : ",n)

	n *= 5
	fmt.Println("n *= : ",n)

	n /= 5
	fmt.Println("n /= : ",n)

	// Incriment Operator
	n++
	fmt.Println("Incriment: ",n)

	// Decrement Operator
	n--
	fmt.Println("Decrement: ",n)


}


/* Terminal output

Addition Operater:  15
Full name:  Anjesh Kumar
Subtraction:  5
Multipication:  50
Division:  2
Remainder:  0
New Sum:  25
n += :  15
n -= :  10
n *= :  50
n /= :  10
Incriment:  11 
Decrement:  10

*/