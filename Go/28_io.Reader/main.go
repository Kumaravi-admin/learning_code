/*
                 io.Reader

• The io.Reader interface has one method Read.
• The io.Reader interface:
  type Reader interface {
      Read(p []byte) (n int, err error)
  }

• Takes a byte slice to read the data into.
• The caller of Read should always  process n bytes returned before  considering error.
• This way you can correctly handle I/O  errors that happen after reading some  bytes.
• Sentinel error io.EOF is returned when  end of file is reached.


              Common Implementations

• *os.File: used to read data from a file.
• strings.Reader: creates a new reader from the provided string.
• bufio.Reader: implements Read for improved performance, using buffers for reading.
• http.Request.Body: allows you to supply the request body via reader.
• bytes.Buffer: used for reading from in-memory byte buffers.

*/

// Step 1: Create a file named letters.txt in the same folder as main.go (put any text in it).

package main

import (
	"fmt"
	"io" // io.Reader interface and io.EOF live here
	"os"
	"strings"
)

func main() {

	// Example 1: Read from a string.
	// strings.NewReader turns a string into an io.Reader.
	r := strings.NewReader("Anjesh Kumar")

	number, err := countAlphabets(r)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Lattes from string: %d\n", number)

	// Example 2: Read from a file.
	// os.Open returns *os.File, which ALSO implements io.Reader.
	f, err := os.Open("letters.txt")
	if err != nil {
		panic(err)
	}
	// Good habit: close the file when main finishes
	defer f.Close()

	// Same function, different source. This is the power of io.Reader.
	n, err := countAlphabets(f)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Lattes from file: %d\n", n)

}

// Step 2: Accept the io.Reader INTERFACE, not a specific type.
// It works with string, file, network, buffer... anything that has Read().
func countAlphabets(r io.Reader) (int, error) {
	count := 0

	// Step 3: Create a byte slice (buffer). Read() fills it with data.
	buffer := make([]byte, 1024)

	for {
		// Step 4: Read returns n = number of bytes actually read, and an error.
		n, err := r.Read(buffer)

		// Step 5: Process the n bytes FIRST, before checking err (slide rule).
		// buffer[:n] = only the valid part of the buffer, ignore the rest.
		for _, l := range buffer[:n] {
			// Count only letters A-Z and a-z (spaces, digits, symbols are skipped)
			if (l >= 'A' && l <= 'Z') || (l >= 'a' && l <= 'z') {
				count++
			}
		}

		// Step 6: io.EOF is NOT a real failure, it just means "no more data".
		if err == io.EOF {
			return count, nil
		}

		// Any other error is a real I/O problem.
		if err != nil {
			return 0, err
		}
	}

}

/*  Terminal output:

Lattes from string: 11
Lattes from file: 10

*/