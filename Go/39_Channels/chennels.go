/*
                  Introduction


  [ Goroutine 1 ]  ◄──────────────►  [ Goroutine 2 ]


• Communication between two or more goroutines.
• Sharing memory by communicating.


                  Creation

                chan string

• Channels are typed.
• Created using the make function.


                  Unidirectional Channels


• Channels we created earlier are bidirectional channels.
• Unidirectional channels can be created with <- operator.
• When <- is in front it indicates a read-only channel.

        <-chan string

• When <- is after it indicates a write-only channel.

        chan<- string

• Unidirectional channels are used as function parameters.
• Conversion from bi-directional to unidirectional channels is implicit.


                  Reading & Writing


• The <- operator is also used to read from and write to channels.
• To read, you place <- operator on the left and to write, you place <- operator on the right.

        value := <-chan
        chan<- 10


                  Blocking Behaviour


• Reading from or writing to channels is a blocking operation.
• Writing to a already full channel or reading from channel that has been written to is blocking.


                  Closing Channels


• The global function close can be used to close a channel.
• Reading from a closed channel is not an error.
• Reading from channel returns two values.

*/

package main

/*
We will cover in this lesson:

1. Creation of channels
2. Conversion (bidirectional to read-only / write-only)
3. Reading & Writing
4. Blocking behaviour
5. Closing channels
*/

import "fmt"

// ---------------------------------------------------------
// Used in example 2 (Conversion).
// Unidirectional channels are used as FUNCTION PARAMETERS.
// Functions must be declared outside main (Go has no nested named functions).
// ---------------------------------------------------------

// chan<- int = write-only. This function can only SEND into the channel.
// Trying to read from "out" here would be a compile error.
func writeOnly(out chan<- int) {
	out <- 10
}

// <-chan int = read-only. This function can only RECEIVE from the channel.
// Trying to write to "in" here would be a compile error.
func readOnly(in <-chan int) {
	value := <-in
	fmt.Println("2. readOnly received:", value)
}

func main() {

	// -----------------------------------------------------
	// 1. Creation of channels
	// -----------------------------------------------------
	// make creates the channel. Channels are typed: this one carries only int values.
	// (This is a normal bidirectional channel: it can read AND write.)
	intChannel := make(chan int)
	fmt.Printf("1. channel type: %T\n", intChannel)

	// -----------------------------------------------------
	// 2. Conversion (bidirectional to read-only / write-only)
	// -----------------------------------------------------
	// We pass the same bidirectional channel to both functions.
	// Go converts it automatically: to chan<- int for writeOnly, to <-chan int for readOnly.
	// The goroutine writes, main reads. Unbuffered channels need both sides at the same time.
	go writeOnly(intChannel)
	readOnly(intChannel)

	// -----------------------------------------------------
	// 3. Reading & Writing
	// -----------------------------------------------------
	// A channel with a buffer of 1 can hold one value without a reader waiting.
	buffered := make(chan int, 1)

	buffered <- 10      // WRITE: <- is on the right of the channel (data goes in)
	value := <-buffered // READ:  <- is on the left of the channel (data comes out)
	fmt.Println("3. value read from channel:", value)

	// -----------------------------------------------------
	// 4. Blocking behaviour
	// -----------------------------------------------------
	// A buffered channel holds up to 2 values (cap = 2).
	limited := make(chan int, 2)
	limited <- 1
	limited <- 2
	fmt.Println("4. channel is full: len =", len(limited), "cap =", cap(limited))

	// limited <- 3
	// ^ This would BLOCK forever, because the channel is full and nobody is reading.
	// With no other goroutine running, Go stops with:
	// fatal error: all goroutines are asleep - deadlock!
	//
	// Same on the other side: reading from an EMPTY channel blocks until someone writes.

	// -----------------------------------------------------
	// 5. Closing channels
	// -----------------------------------------------------
	done := make(chan int, 2)
	done <- 1
	done <- 2

	// close(ch) means: "no more values will be sent".
	// Only the SENDER should close a channel.
	close(done)

	// Reading returns TWO values: the value, and ok.
	// ok = true  -> a real value was received
	// ok = false -> channel is closed AND empty (value is the zero value, 0 for int)
	v, ok := <-done
	fmt.Println("5. read:", v, ok) // 1 true

	v, ok = <-done
	fmt.Println("5. read:", v, ok) // 2 true

	v, ok = <-done
	fmt.Println("5. read:", v, ok) // 0 false (closed and empty, but NOT an error)
}

/*  Terminal output:

1. channel type: chan int
2. readOnly received: 10
3. value read from channel: 10
4. channel is full: len = 2 cap = 2
5. read: 1 true
5. read: 2 true
5. read: 0 false

*/