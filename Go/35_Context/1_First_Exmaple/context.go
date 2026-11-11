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
	"time"
)

// Step 1: A task that runs forever until the context tells it to stop.
// By convention, ctx is the FIRST parameter (slide: General Rules).
func SimulateLongRunninTask(ctx context.Context) {
	for {
		// select waits on several channels and runs whichever is ready first.
		select {

		// ctx.Done() returns a channel that gets CLOSED when the context is
		// cancelled or its timeout is reached. Receiving from it means "stop now".
		case <-ctx.Done():
			fmt.Println("Context cancelled, task stopping...")
			return // exit the function, so the goroutine ends

		// default runs when ctx is NOT done yet: keep doing the work.
		default:
			fmt.Println("Simulating work...")
			time.Sleep(1 * time.Second) // pretend the work takes 1 second
		}
	}
}

func main() {

	// Step 2: Create a context that auto-cancels after 5 seconds.
	// context.Background() is the root (empty) context. We derive from it (slide: Derive Context).
	// WithTimeout returns the new ctx and a cancel function.
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	// Step 3: defer cancel() makes sure resources are released when main exits,
	// even if the timeout never fired.
	defer cancel()

	// Step 4: Start the task in a goroutine (runs in the background, concurrently).
	// We pass ctx so the task can see when it should stop.
	go SimulateLongRunninTask(ctx)

	// Step 5: Main does its own work for 2 seconds.
	// Meanwhile the goroutine prints "Simulating work..." about twice.
	time.Sleep(2 * time.Second)

	// Step 6: Cancel manually, BEFORE the 5 second timeout is reached.
	// This closes ctx.Done(), so the task will stop at its next loop check.
	fmt.Println("Cancelling context...")
	cancel()

	// Step 7: Give the goroutine a moment to print its stop message.
	// Without this, main could exit first and the message would never appear.
	// (In real code you would use sync.WaitGroup instead of Sleep.)
	time.Sleep(1 * time.Second)
}

/*  Terminal output:

Simulating work...
Simulating work...
Cancelling context...
Context cancelled, task stopping...

*/