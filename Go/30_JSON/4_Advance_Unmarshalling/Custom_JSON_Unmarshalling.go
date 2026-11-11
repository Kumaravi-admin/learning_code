package main

import (
	"encoding/json"
	"fmt"
)

// Step 1: Create the enum (same as the Marshal part).
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

// Step 3: ONE lookup table: text -> number.
// Decoding only needs this direction.
var stringToStatus = map[string]Status{
	"pending":   Pending,
	"shipped":   Shipped,
	"delivered": Delivered,
}

// Step 4: UnmarshalJSON = JSON -> Go.
// json.Unmarshal calls this automatically when it fills a Status.
//
// Note the * (pointer): we need to CHANGE the original variable,
// so we must receive its address, not a copy.
func (s *Status) UnmarshalJSON(data []byte) error {

	// "data" is raw JSON like "shipped" (it still has the double quotes).
	// Decode it into a normal Go string to remove the quotes.
	var text string
	err := json.Unmarshal(data, &text)
	if err != nil {
		return err
	}

	// Find the number for this text
	value, found := stringToStatus[text]

	// Not found = unknown text (for example "banana"), so return an error
	if !found {
		return fmt.Errorf("invalid status: %s", text)
	}

	// Store the number into the original variable
	*s = value
	return nil
}

func main() {

	// Test 1: valid text
	// We never call UnmarshalJSON ourselves. Go calls it for the Status field.
	var order Order
	err := json.Unmarshal([]byte(`{"id":101,"status":"shipped"}`), &order)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Println(order.ID)                // 101
	fmt.Println(int(order.Status))       // 2 (the number stored inside)
	fmt.Println(order.Status == Shipped) // true

	// Test 2: invalid text
	var badOrder Order
	err = json.Unmarshal([]byte(`{"id":102,"status":"banana"}`), &badOrder)
	if err != nil {
		fmt.Println("error:", err)
	}
}

/*  Terminal output:

101
2
true
error: invalid status: banana

*/