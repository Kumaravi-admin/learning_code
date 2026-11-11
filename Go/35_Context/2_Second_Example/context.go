/*
	Context

• Context allows you to signal that the operation can be cancelled.
• You can associate a deadline with a context, specifying the maximum time for an operation to complete.
• Context can be used to carry arbitrary data along with the function calls.

	                  Illustration


	                     ┌────────┐
	                     │  main  │
	                     └────────┘
	           ctx ╱                ╲ ctx
	              ╱                  ╲
	             ▼                    ▼
	      ┌──────────┐          ┌───────────┐
	      │  func 1  │          │    go     │
	      └──────────┘          │ routine 1 │
	     ctx ╱     ╲ ctx        └───────────┘
	        ╱       ╲                 │ ctx
	       ▼         ▼                ▼
	┌──────────┐ ┌──────────┐   ┌───────────┐
	│  func 2  │ │  func 3  │   │    go     │
	└──────────┘ └──────────┘   │ routine 2 │
	                            └───────────┘

	                 Uses

• HTTP Request context can carry values through middlewares and handlers.
• Database operations can rely on context for cancellation.
• Long running goroutines can be cancelled with context.

	Context Interface

	type Context interface {
	   Deadline() (deadline time.Time, ok bool)
	   Done() <-chan struct{}
	   Err() error
	   Value(key any) any
	}

	General Rules

• Context creation should be done at the root of the execution call.
• By convention context is the first parameter of a function.
• Context should not be saved in variable for later use.
• Context should not be mutated.

	Derive Context

• func WithCancel(parent Context) (ctx Context, cancel CancelFunc)
• func WithTimeout(parent Context, timeout time.Duration) (Context, CancelFunc)
• func WithValue(parent Context, key, val any) Context
*/
package main

import (
	"context"
	"fmt"
)

// Step 3: This function RECEIVES the context (ctx is the first parameter, by convention).
// It does not get userID as a separate parameter. It reads it from the context.
func ProcessOrder(ctx context.Context, orderId int) {

	// ctx.Value("userID") looks up the key and returns it as type "any".
	// .(int) is a type assertion: "I expect this value to be an int".
	// ok is true if the key exists AND the value is an int (comma ok pattern, no panic).
	userID, ok := ctx.Value("userID").(int)
	if !ok {
		fmt.Println("Error: userID not found in context")
		return
	}

	fmt.Printf("Processing order %d for user %d\n", orderId, userID)
}

// Step 1: Create the context at the ROOT of the call (slide: General Rules).
func GetUserAndProcessOrder(orderid int) {

	// WithValue creates a NEW context that carries a key-value pair.
	// Parent = Background() (empty root context), key = "userID", value = 123.
	ctx := context.WithValue(context.Background(), "userID", 123)

	// Step 2: Pass ctx down. ProcessOrder can now read userID from it.
	ProcessOrder(ctx, orderid)
}

func main() {
	GetUserAndProcessOrder(456)
}

/*  Terminal output:

Processing order 456 for user 123

*/