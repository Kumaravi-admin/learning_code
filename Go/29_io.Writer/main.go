/*

                 io.Writer


• The io.Writer interface has one method Write.
• Takes the source data as slice of byte.
• The io.Writer interface:
  type Writer interface {
      Write(p []byte) (n int, err error)
  }


                Implementing io.Writer

• If written bytes are less than the  provided bytes you should return an error.
• Provided slice of byte is a source of writing and should not be modified.


              Common Implementations

• *os.File: used to read data from a file.
• strings.Reader: creates a new reader from the provided string.
• bufio.Reader: implements Read for improved performance, using buffers for reading.
• http.Request.Body: allows you to supply the request body via reader.
• bytes.Buffer: used for reading from in-memory byte buffers.

*/

package main

import (
	"fmt"
	"io" // io.Writer interface lives here
	"os"
)

func main() {

	// Step 1: Create a file (if it exists, it is emptied).
	// *os.File implements io.Writer because it has a Write() method.
	f, err := os.Create("writting.txt")
	if err != nil {
		panic(err)
	}
	// Step 2: Close the file when main finishes (good habit, frees the resource).
	defer f.Close()

	// Step 3: Pass the file as an io.Writer.
	n, err := writeString("Anjesh Kumar Ravi", f)
	if err != nil {
		panic(err)
	}

	// n = number of bytes written. "Anjesh Kumar Ravi" = 17 characters = 17 bytes.
	fmt.Printf("Written bytes: %d\n", n)

}

// Step 4: Accept the io.Writer INTERFACE, not a specific type.
// It works with a file, network connection, buffer, os.Stdout... anything with Write().
func writeString(s string, w io.Writer) (int, error) {

	// Step 5: Write needs a []byte, so convert the string first.
	// It returns n (bytes written) and an error.
	n, err := w.Write([]byte(s))
	if err != nil {
		// %w wraps the original error, so the caller can still inspect it
		// with errors.Is / errors.As.
		return 0, fmt.Errorf("error occurred while writing %w", err)
	}

	return n, nil
}

/*  Terminal output:

Written bytes: 17

writting.txt content:
Anjesh Kumar Ravi

*/