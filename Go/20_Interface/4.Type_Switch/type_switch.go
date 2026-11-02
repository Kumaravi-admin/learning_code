/*

                  Type Switch

• Type switch can be used to check against multiple types.
• Type assertion is done in the switch statement instead of switching on the value.
• Inside each case the value will of that type if case matches, otherwise remains empty interface type.


                  Use Case

• Type assertion and type switch should be used  sparingly.
• Used to check if type implements a interface or when type is not known beforehand.

*/

package main

import "fmt"

// Step 1: Accept a value of ANY type (empty interface).
func checkType(value interface{}) {

	// Step 2: Type switch.
		switch v := value.(type) {
	case int:
		fmt.Println("type is an int: ", v)
	case string:
		fmt.Println("type is an string: ", v)
	case float64:
		fmt.Println("type is an float: ", v)
	default:
		// No case matched, so v stays as interface{}
		fmt.Println("unknown type")
	}

}

func main() {

	// Step 3: Call the same function with different types.
	checkType(42)          // matches int
	checkType("hello,Go!") // matches string
	checkType(3.14)        // matches float64 (decimals default to float64)
	checkType(true)        // bool has no case, so goes to default

}

/*				Terminal output:


type is an int:  42
type is an string:  hello,Go!
type is an float:  3.14
unknown type


*/