/*
                  Two Ways


• Use a struct to couple result & error.
• Use a separate error channel.

*/

package main

/*
We will cover in this lesson:

1. Way 1: Use a struct to couple result & error
*/

import (
	"fmt"
	"net/http"
)

// ---------------------------------------------------------
// 1. Way 1: Use a struct to couple result & error
// ---------------------------------------------------------
// A goroutine cannot "return" a value and an error like a normal function.
// It can only send things into a channel. So we put BOTH the result and the error
// inside one struct, and send that struct. The receiver checks both fields.

// Step 1: The struct that carries the result AND the error together.
type result struct {
	resp *http.Response // the result (nil if the request failed)
	err  error          // the error (nil if the request succeeded)
}

// Step 2: A function that checks many URLs inside a goroutine.
// done is read-only (<-chan): the parent uses it to tell the goroutine "stop now".
// It returns a read-only channel of result, so the caller can only receive.
func checkIfExist(done <-chan struct{}, urls ...string) <-chan result {
	results := make(chan result)

	go func() {
		// Step 5: Close results when this goroutine ends (even on early return).
		// Closing tells the receiver "no more values", so its "for range" loop can stop.
		defer close(results)

		for _, url := range urls {
			select {

			// The parent closed done, so stop immediately (termination path 3).
			case <-done:
				return

			// Otherwise, make the request.
			default:
				// Step 3: http.Get returns two values: response and error.
				// We do NOT stop on error, we pack both into the struct and send it.
				res, err := http.Get(url)
				results <- result{resp: res, err: err}
			}
		}
	}()

	return results
}

func main() {
	done := make(chan struct{})

	// Second URL uses port 300, which has no server, so it will fail on purpose.
	// This lets us see both the success and the error case.
	results := checkIfExist(done, "https://google.com", "http://localhost:300")

	// Step 4: "for range" on a channel keeps receiving until the channel is closed.
	for r := range results {

		// Success: resp is set, so print the status.
		if r.resp != nil {
			fmt.Println(r.resp.Status)
			r.resp.Body.Close() // always close the body to avoid a connection leak
		}

		// Failure: err is set, so print the error.
		if r.err != nil {
			fmt.Println(r.err)
		}
	}

	// Tell the goroutine to stop if it is still running.
	close(done)
}

/*  Terminal output (the error text depends on your OS):

200 OK
Get "http://localhost:300": dial tcp [::1]:300: connectex: No connection could be made because the target machine actively refused it.

*/