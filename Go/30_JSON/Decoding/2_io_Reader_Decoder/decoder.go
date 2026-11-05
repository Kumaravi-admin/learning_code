/*
         JSON Encoding & Decoding


• The io.Reader & the io.Writer  interface support reading and  writing byte data.
• It is efficient to use io.Reader or  io.Writer to read JSON byte data  directly into Go structs and vice-versa.
• This is where encoding/json.Encoder  and encoding/json.Decoder comes into  play.

*/

// Step 1: Create a file named output.json in the same folder as main.go with this content:
// {"username":"Anjesh","age":25,"is_active":true}

package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// Step 2: Struct with JSON tags.
// Tags map JSON keys (username, age, is_active) to Go fields (Name, Age, Active).
type Person struct {
	Name   string `json:"username"`
	Age    int    `json:"age"`
	Active bool   `json:"is_active"`
}

func main() {
	var p Person

	// Step 3: Open the file. *os.File implements io.Reader,
	// so a Decoder can read from it directly.
	f, err := os.Open("output.json")
	if err != nil {
		fmt.Println("error opening file: ", err)
		panic(err)
	}

	// Close the file when main finishes
	defer f.Close()

	// Step 4: Create a Decoder that reads from any io.Reader (here, the file).
	decoder := json.NewDecoder(f)

	// Step 5: Decode reads JSON from the file and fills the struct.
	// Pass a POINTER (&p) so Decode can set the fields.
	err = decoder.Decode(&p)
	if err != nil {
		fmt.Println("error decoding persion: ", err)
		panic(err)
	}

	// Step 6: Print the struct.
	fmt.Println(p)

}

/*  Terminal output (for the sample output.json above):

{Anjesh 25 true}

*/