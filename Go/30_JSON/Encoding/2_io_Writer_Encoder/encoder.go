/*
         JSON Encoding & Decoding


• The io.Reader & the io.Writer  interface support reading and writing byte data.
• It is efficient to use io.Reader or io.Writer to read JSON byte data directly into Go structs and vice-versa.
• This is where encoding/json.Encoder and encoding/json.Decoder comes into play.

*/

package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// Step 1: Struct with JSON tags (Go field -> JSON key).
type Person struct {
	Name   string `json:"username"`
	Age    int    `json:"age"`
	Active bool   `json:"is_active"`
}

func main() {

	// Step 2: Create the Go value we want to write as JSON.
	p := Person{
		Name:   "Anjesh",
		Age:    25,
		Active: true,
	}

	// Step 3: Create (or overwrite) the file.
	// *os.File implements io.Writer, so an Encoder can write to it directly.
	f, err := os.Create("output.json")
	if err != nil {
		fmt.Println("err createing file: ", err)
		panic(err)
	}

	// Close the file when main finishes
	defer f.Close()

	// Step 4: Create an Encoder that writes to any io.Writer (here, the file).
	encoder := json.NewEncoder(f)

	// Step 5: Encode converts the struct to JSON and writes it into the file.
	// It also adds a newline at the end.
	err = encoder.Encode(p)
	if err != nil {
		fmt.Println("error encoding persion: ", err)
		panic(err)
	}

}

/*  Terminal output: (nothing is printed)

output.json file content:

{"username":"Anjesh","age":25,"is_active":true}

*/