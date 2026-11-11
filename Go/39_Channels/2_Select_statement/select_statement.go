/*
                  Introduction


• Similar syntax to the switch statement.
• select {
  case value := <-ch1:
      fmt.Println("Received from ch1:", value)
  case ch2 <- struct{}{}:
      fmt.Println("Written to ch2")
  }


                  Switch vs Select


• Functioning of switch & select is very different.
• select {
  case value := <-ch1:
      fmt.Println("Received from ch1:", value)
  case ch2 <- struct{}{}:
      fmt.Println("Written to ch2")
  }
• Each case in select is evaluated simultaneously.


                  Simultaneous Evaluation


• Go uses pseudo-random uniform selection.
• Each case that is ready has equal probability of getting executed.


                  Unblocking Select


• The select statement unblocks as soon as the channels get unblocked.
• To continuously evaluate channels use for-select statement.
• Empty select statement blocks forever.

*/

package main

/*
We will cover in this lesson:

1. Basic select (waits for the first ready channel)
2. for-select (keep receiving until all values arrive)
3. for-select with timeout (time.After)
4. Empty select (blocks forever, commented out)
*/

import (
	"fmt"
	"time"
)

// ---------------------------------------------------------
// Helper used by examples 1, 2 and 3.
// It creates two channels and starts two goroutines.
// Each goroutine sends 0, 1, 2 (one value, then waits 1 second).
// Every example calls it, so each example gets its own fresh channels.
// ---------------------------------------------------------
func startSenders() (chan int, chan int) {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		for i := range 3 {
			ch1 <- i
			time.Sleep(1 * time.Second)
		}
	}()

	go func() {
		for i := range 3 {
			ch2 <- i
			time.Sleep(1 * time.Second)
		}
	}()

	return ch1, ch2
}

// ---------------------------------------------------------
// 1. Basic select
// ---------------------------------------------------------
// select waits until ONE of its cases is ready, runs that case, and then ends.
// If both channels are ready, Go picks one at random (pseudo-random uniform selection).
// So the output can be ch1 or ch2, it changes between runs.
func example1BasicSelect() {
	ch1, ch2 := startSenders()

	select {
	case value := <-ch1:
		fmt.Println("1. Received from ch1:", value)
	case value := <-ch2:
		fmt.Println("1. Received from ch2:", value)
	}
}

// ---------------------------------------------------------
// 2. for-select (keep receiving)
// ---------------------------------------------------------
// A single select handles only ONE event. To keep handling events, put it inside a loop.
// Each sender sends 3 values, so we loop 6 times to receive all of them.
func example2ForSelect() {
	ch1, ch2 := startSenders()

	for range 6 {
		select {
		case value := <-ch1:
			fmt.Println("2. Received from ch1:", value)
		case value := <-ch2:
			fmt.Println("2. Received from ch2:", value)
		}
	}
}

// ---------------------------------------------------------
// 3. for-select with timeout
// ---------------------------------------------------------
// time.After(d) returns a channel that receives a value after duration d.
// Used as a select case, it works as a timeout: "if nothing arrives in 2 seconds, stop".
// The loop has no fixed count, it ends only when the timeout case fires.
func example3SelectTimeout() {
	ch1, ch2 := startSenders()

	for {
		select {
		case value := <-ch1:
			fmt.Println("3. Received from ch1:", value)
		case value := <-ch2:
			fmt.Println("3. Received from ch2:", value)
		case <-time.After(2 * time.Second):
			fmt.Println("3. returning after timeout")
			return
		}
	}
}

// ---------------------------------------------------------
// 4. Empty select (blocks forever)
// ---------------------------------------------------------
// "select {}" has no cases, so it waits forever.
// If no other goroutine is running, Go stops with:
// fatal error: all goroutines are asleep - deadlock!
// Same thing happens with nil channels (var ch1, ch2 chan any), they are never ready.
// That is why this example is commented out in main.
func example4EmptySelect() {
	fmt.Println("4. waiting forever...")
	select {}
}

func main() {
	example1BasicSelect()
	example2ForSelect()
	example3SelectTimeout()
	// example4EmptySelect() // uncomment to see the deadlock error
}

/*  Terminal output (ch1 / ch2 order changes between runs):

1. Received from ch2: 0
2. Received from ch1: 0
2. Received from ch2: 0
2. Received from ch1: 1
2. Received from ch2: 1
2. Received from ch2: 2
2. Received from ch1: 2
3. Received from ch1: 0
3. Received from ch2: 0
3. Received from ch2: 1
3. Received from ch1: 1
3. Received from ch1: 2
3. Received from ch2: 2
3. returning after timeout

*/