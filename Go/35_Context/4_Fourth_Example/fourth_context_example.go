package main

import (
	"context"
	"fmt"
	"time"
)

// makeChai takes 3 seconds to finish.
// But it also keeps listening to ctx: if ctx says "stop", it stops early.
func makeChai(ctx context.Context) {
	fmt.Println("Chai ban rahi hai...")

	// select = wait for whichever happens FIRST
	select {

	// Case 1: 3 seconds pass, the chai is ready
	case <-time.After(3 * time.Second):
		fmt.Println("Chai ready!")

	// Case 2: ctx says "stop" (patience over, or cancel() was called)
	case <-ctx.Done():
		fmt.Println("Chai cancel hui, reason:", ctx.Err())
	}
}

func main() {

	// Create a context with a patience limit of 1 second.
	// Background() is just the empty starting context.
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)

	// Always call cancel when done, it frees resources.
	defer cancel()

	makeChai(ctx)
}

/*  Terminal output:

Chai ban rahi hai...
Chai cancel hui, reason: context deadline exceeded

*/