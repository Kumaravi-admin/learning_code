/*

                  net/http

• The net/http package exports default  HTTP client.
• The package also provides function to  make GET, HEAD, POST requests.
• Highly recommended to not use these.


Package net/http: [https://pkg.go.dev/net/http]

*/

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// Sample API: GET https://www.boredapi.com/api/activity
// It returns JSON like this:
// {
//   "activity": "Learn Express.js",
//   "type": "education",
//   "participants": 1,
//   "price": 0.1,
//   "link": "https://expressjs.com/",
//   "key": "3943506",
//   "accessibility": 0.25
// }

// Step 1: Create a struct that matches the JSON response.
// JSON tags map each JSON key to a Go field (all fields start with uppercase, so they are exported).
type boringResponse struct {
	Activity      string  `json:"activity"`
	Type          string  `json:"type"`
	Participants  int     `json:"participants"`
	Price         float64 `json:"price"`
	Link          string  `json:"link"`
	Key           string  `json:"key"`
	Accessibility float64 `json:"accessibility"`
}

func main() {

	// Step 2: Create a context. Background() means "no special rules".
	// (Context lets you cancel a request or set a deadline.)
	ctx := context.Background()

	// Step 3: Create our OWN http client with a timeout.
	// The slide says not to use the default helpers (http.Get, http.Post...),
	// because the default client has NO timeout. If the server never answers,
	// your program would wait forever.
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	// Step 4: Build the request (method + URL + no body, so nil).
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "https://www.boredapi.com/api/activity", nil)
	if err != nil {
		fmt.Println("error creating request: ", err)
		return
	}

	// Step 5: Send the request and get the response.
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("error doing request: ", err)
		return
	}

	// Step 6: Always close the response body when done, otherwise the connection leaks.
	defer resp.Body.Close()

	// Step 7: Check the status code. 200 (StatusOK) means success.
	if resp.StatusCode != http.StatusOK {
		fmt.Println("error invalid status code: ", resp.Status)
		return
	}

	// Step 8: Decode the JSON body into our struct.
	// resp.Body is an io.Reader, so json.NewDecoder can read from it directly.
	var boringRes boringResponse
	if err := json.NewDecoder(resp.Body).Decode(&boringRes); err != nil {
		fmt.Println("error reading json response body: ", err)
		return
	}

	// Step 9: Print the struct. %+v prints field names with values.
	fmt.Printf("%+v\n", boringRes)

}

/*  Terminal output (your values will be different):

{Activity:Learn Express.js Type:education Participants:1 Price:0.1 Link:https://expressjs.com/ Key:3943506 Accessibility:0.25}

*/