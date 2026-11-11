/*

        Marshalling/Unmarshalling

• Converting Go data structures to JSON data is called marshalling.
• Converting JSON data into Go data structures is called unmarshalling.
• Package encoding/json providedfunctions to marshal and unmarshal.
• Tags should be used for mapping.


                  JSON Tags

• JSON tags are used in Go to map a  field of a struct to it’s JSON  counterpart.
• Added inside backticks at end of the struct field declaration.
• Inside the tags you provide the  mapping and other options.


                 Prefer Using Tags

• If no tags are provided to map field  names, by default the field name
  should match it’s JSON counterpart.
• Go uses the case of the field name to decide if it’s exported or not.
• JSON Marshal & Unmarshal require the  struct fields to be exported.


         JSON Encoding & Decoding

• The io.Reader & the io.Writer  interface support reading and  writing byte data.
• It is efficient to use io.Reader or io.Writer to read JSON byte data directly into Go structs and  vice-versa.
• This is where encoding/json.Encoder  and encoding/json.Decoder comes into  play.

*/

package main

import (
	"encoding/json"
	"fmt"
)

// Step 1: Create a struct and add JSON tags.
// A tag is written inside backticks after the field type.
// Format: `json:"name_in_json,option"`
// All field names start with an uppercase letter (exported),
// because the json package can only see exported fields.
type Person struct {
	Name       string   `json:"full_name"`           // maps to the "full_name" key in JSON
	Age        int      `json:"years_old,omitempty"` // maps to "years_old"; omitted when marshalling if value is 0
	Occupation string   `json:"-"`                   // "-" = completely ignored by JSON (never read or written)
	Languages  []string `json:"spoken_languages"`    // maps to the "spoken_languages" array
}

func main() {

	// Step 2: JSON data as a string (backticks allow double quotes inside).
	// Notice: there is no "Occupation" key here, and even if there were, it would be ignored.
	jsonData := `{"full_name":"Anjesh Kumar", "years_old":25,"spoken_languages":["English","Hindi","Spanish"]}`

	// Step 3: Unmarshal = convert JSON into a Go struct.
	var person Person

	// json.Unmarshal needs:
	//   1. the JSON as []byte (so convert the string)
	//   2. a POINTER to the struct (&person), so it can fill the fields
	err := json.Unmarshal([]byte(jsonData), &person)
	if err != nil {
		fmt.Println("Error unmarshalling JSON: ", err)
		panic(err)
	}

	// Step 4: Print the unmarshalled data.
	fmt.Println("Name: ", person.Name)
	fmt.Println("Age: ", person.Age)
	fmt.Println("Languages: ", person.Languages)

	// Step 5: Occupation was tagged with "-", so it stays empty (zero value).
	// Accessing it is allowed, but it will never be filled from JSON (not recommended).
	fmt.Println("Occupation: ", person.Occupation)

}

/*  Terminal output:

Name:  Anjesh Kumar
Age:  25
Languages:  [English Hindi Spanish]
Occupation:

*/