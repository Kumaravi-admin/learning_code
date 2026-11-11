package main

import (
	"context"
	"fmt"
	"time"
)

// counter keeps counting every second until ctx says "stop"
func counter(ctx context.Context) {
	i := 1
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Counter ruk gaya")
			return
		default:
			fmt.Println("Count:", i)
			i++
			time.Sleep(1 * time.Second)
		}
	}
}

func main() {

	// WithCancel = no time limit, we stop it manually
	ctx, cancel := context.WithCancel(context.Background())

	go counter(ctx) // run counter in the background

	time.Sleep(3500 * time.Millisecond) // let it count for ~3 seconds

	cancel() // tell the counter: "stop now"

	time.Sleep(1 * time.Second) // give it a moment to print its stop message
}

/*  Terminal output:

Count: 1
Count: 2
Count: 3
Counter ruk gaya

*/