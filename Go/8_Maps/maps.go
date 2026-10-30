/*
                     Maps

• Maps are data structures used to store  key-value pairs.
• Behind the scenes Go uses a hash maps to implement maps.
• The len function can be used to get the number of key/value pairs in a map.


					Examples

			nameAge := map[string]int{
				"Foo": 25,
				"Bar": 20,
			}

			firstNameLastName := map[string]string{
				"Foo": "Bar",
				"Bar": "Foo",
			}
------------------------------------------------------------------

                   Nil Map

• The zero value for a map is nil.
• A nil map can be created using a variable declaration.
• Example: var nameAge map[string]int
• The len of nil map is 0.
• Reading from a nil map will return zero value for the value type.
• Writing to a nil map will cause panic.


                  Empty Map

• A empty map can be created via multiple  ways, we will look at some examples  next.
• The len of empty map is 0.
• You can read and unlike nil map write to it as well.


                  Examples

			nameAge := map[string]int{}

			nameAge := make(map[string]int)

			var nameAge map[string]int = map[string]int{}
------------------------------------------------------------------

                 Read from Map

• Value from a map can be accessed using the key.
• Example: nameAge["Foobar"]
• If key does not exist, returns the zero value.
• The comma ok idiom in Go:
    fooAge, ok := nameAge["foo"]
    if !ok {
        fmt.Println("not found")
    }
------------------------------------------------------------------

                 Write to Map

• Value can to written to a map using the   key.
• Example: nameAge["Foobar"] = 50
• A map can have unique keys only.
• Using duplicate keys will overwrite any existing key value pair.


               Some More Facts

• Map keys can only be of types.
• Values can be of any type.
• Map can only be compared to nil.
• Since keys are unique, maps can be used  as sets.


*/

package main

import "fmt"

// In this lesson we will cover:
// 1. A nil map.
// 2. a empty map.
// 3. Writing to a map.
// 4. Reading from map.
// 5. Map value as string key as slice
func main() {


	// Creation of a nil maps.
	var nameAge map[string]int
	fmt.Println("Name age: ",len(nameAge))									// output: 0 

	// Empty map using: Variable shorthand.
	shorthandEmptyMap := map[string]int{}
	fmt.Println("Shorthand empty map: ",len(shorthandEmptyMap))				// output: 0

	// Empty map using: Make funcation
	makeFunctionEmptyMap := make(map[string]int)
	fmt.Println("Make funcation empty map: ", len(makeFunctionEmptyMap))	// output: 0


	// Long or litral variable declaration
	var longVariableEmptyMap map[string]int = map[string]int{}
	fmt.Println("Long variable empty map: ",len(longVariableEmptyMap))		// output: 0



	// Map with initial value
	fruitsPrice := map[string]int{
		"Papaya" : 80,
		"Graps"	: 110,
		"Coconut" : 60,

	}

	// Add value to maps.
	fruitsPrice["Apple"] = 120
	fruitsPrice["Mango"] = 60
	fruitsPrice["Banana"]= 30

	fmt.Println("lenth of fruits price map: ", len(fruitsPrice)) 			// output: 6


	// Reading from maps.
	fmt.Println("Reding Mango Price: ", fruitsPrice["Mango"])				// output: 60
	fmt.Println("Reding Graps Price: ", fruitsPrice["Graps"])				// output: 110

	// Try to read key value that not exist in maps 
	fmt.Println("Reding Orange Price: ", fruitsPrice["Orange"])				// output: 0

	// comma ok idom
	studentGrade := map[string]int{
		"anjesh": 10,
		"ankit": 30,
		"ashitosh":40,
		"sanjeet": 30,
	}


	// Read studen grade maps value
	// Try to read key value that not exist in maps
	
	// 1. return 0 if key not exist in map
	fmt.Println("studen grade amir: ",studentGrade["amir"])					// output: 0

	// 2. key and ok, if key not exist then ok will return false, 
	student_amir_grade, ok := studentGrade["amir"]
	fmt.Println("student grade amir: ",student_amir_grade," Ok : ",ok)		// output :  0  Ok :  false

	// 3. if key exis then ok will return true
	student_anjesh_grade, ok := studentGrade["anjesh"]
	fmt.Println("student grade anjesh: ",student_anjesh_grade," Ok : ",ok) 	//	output : 10  Ok :  true


	// map decleration key as string and value as slice
	a := map[string][]int{
		"n": {1,2,3,4,5},
	}
	fmt.Println("Output of map a value of n: ",a["n"])						// output: [1 2 3 4 5]

}

/*				Terminal output

Name age:  0
Shorthand empty map:  0
Make funcation empty map:  0
Long variable empty map:  0
lenth of fruits price map:  6
Reding Mango Price:  60
Reding Graps Price:  110
Reding Orange Price:  0
studen grade amir:  0
student grade amir:  0  Ok :  false
student grade anjesh:  10  Ok :  true
Output of map a value of n:  [1 2 3 4 5]



*/