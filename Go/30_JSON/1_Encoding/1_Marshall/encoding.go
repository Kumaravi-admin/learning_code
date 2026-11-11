/*

        Marshalling/Unmarshalling


• Converting Go data structures to JSON data is called marshalling.
• Converting JSON data into Go data  structures is called unmarshalling.
• Package encoding/json provided  functions to marshal and unmarshal.
• Tags should be used for mapping.


                  JSON Tags

• JSON tags are used in Go to map a field of a struct to it’s JSON counterpart.
• Added inside backticks at end of the struct field declaration.
• Inside the tags you provide the mapping and other options.


          Common JSON Tag Options

• Tag ‘omitempty’ is used to omit the  struct field from marshalling when  zero-valued.
• The ‘-’ is used to indicate the  field should be completely ignored from marshalling.


                 Prefer Using Tags

• If no tags are provided to map field  names, by default the field name should match it’s JSON counterpart.
• Go uses the case of the field name  to decide if it’s exported or not.
• JSON Marshal & Unmarshal require the struct fields to be exported.



*/

package main

import (
	"encoding/json"
	"fmt"
)

// Target JSON shape (from the slide):
// {
//  “user_id”: 12345,
//  “name”: “User A”,
//  “age”: 35,
//  “password”: “my-password”,
//  “roles”: [“admin”, “collaborator”]
// }

// Step 1: Struct with JSON tags.
// Tags decide the JSON key names and special behavior.
type User struct {
	ID          int      `json:"User_id"`        // maps ID to "User_id" key (note: capital U, see note below)
	Name        string   `json:"name,omitempty"` // "name" key; skipped if Name is empty ("")
	Age         int      `json:"age"`            // "age" key
	Password    string   `json:"-"`              // "-" = never included in JSON (keeps secrets out)
	Permissions []string `json:"roles"`          // field is Permissions in Go, but "roles" in JSON
}

func main() {

	// Step 2: Create a Go struct value.
	u := User{
		ID:          1,
		Name:        "Anjesh Kumar",
		Age:         26,
		Password:    "my-password",
		Permissions: []string{"admin", "group-member"},
	}

	// Step 3: Marshal = convert Go struct into JSON.
	// It returns the JSON as []byte, plus an error.
	body, err := json.Marshal(u)
	if err != nil {
		fmt.Println("error marshalling JSON: ", err)
		panic(err)
	}

	// Step 4: Convert []byte to string to print it as readable text.
	fmt.Println(string(body))

}

/*  Terminal output:

{"User_id":1,"name":"Anjesh Kumar","age":26,"roles":["admin","group-member"]}

*/