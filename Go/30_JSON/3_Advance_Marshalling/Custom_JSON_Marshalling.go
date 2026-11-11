package main

import (
	"encoding/json"
	"errors"
	"fmt"
)

// Step 1: Create the enum.
// Go has no "enum" keyword, so we make a new type based on int.
type Status int

const (
	_         Status = iota // 0 is skipped (means "no value chosen")
	Pending                 // 1
	Shipped                 // 2
	Delivered               // 3
)

// Step 2: A normal struct that uses the enum as a field.
type Order struct {
	ID     int    `json:"id"`
	Status Status `json:"status"`
}

// Step 3: ONE lookup table: number -> text.
// Encoding only needs this direction.
var statusToString = map[Status]string{
	Pending:   "pending",
	Shipped:   "shipped",
	Delivered: "delivered",
}

// Step 4: MarshalJSON = Go -> JSON.
// json.Marshal calls this automatically when it sees a Status.
func (s Status) MarshalJSON() ([]byte, error) {

	// Find the text for this number
	name, found := statusToString[s]

	// Not found = invalid value (for example Status(99)), so return an error
	if !found {
		return nil, errors.New("invalid status provided")
	}

	// JSON text must be inside double quotes, so we add them: "shipped"
	return []byte(`"` + name + `"`), nil
}

func main() {

	// Test 1: valid value
	// We never call MarshalJSON ourselves. Go calls it for the Status field.
	order := Order{ID: 101, Status: Shipped}
	jsonBytes, err := json.Marshal(order)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Println(string(jsonBytes))

	// Test 2: invalid value (99 is not in the map)
	badOrder := Order{ID: 102, Status: Status(99)}
	_, err = json.Marshal(badOrder)
	if err != nil {
		fmt.Println("error:", err)
	}
}

/*  Terminal output:

{"id":101,"status":"shipped"}
error: json: error calling MarshalJSON for type main.Status: invalid status provided

*/