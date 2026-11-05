/*
                   Introduction

• The underscore import and init functions are related.
• The main function and the init function are two special reserved functions in Go.
• The init function does not accepts any parameters and does not return anything.


                   Execution

• The init function is called implicitly when a package is imported.
• It is executed before the main function.
• You can have multiple init functions in your file.
• You can multiple init functions across files in your package.


                   Use Cases

• It is utilized to initialize resources or register a package.
• Also used for side effects like initializing metrics.
• It an also be used for one-time initialization of global objects.

                   Remember


• The init function runs once per package import.
• Order of execution cannot be relied upon.

*/

// Step 1: Run mod command (once, in project root): go mod init github.com/Kumaravi-admin/init-example

package main

import (
	"fmt"

	// Step 2: Underscore import.
	// We don't call anything from this package. We import it ONLY so that
	// its init() function runs (a side effect). Without "_", Go would give
	// an "imported and not used" error.
	_ "github.com/Kumaravi-admin/init-example/init"
)

// Step 3: init() is a special function.
// - No parameters, no return value
// - Never called by you, Go calls it automatically
// - Runs BEFORE main()
// - You can write many init() functions, even in the same file
func init() {
	fmt.Println("int 1 called")
}

func init() {
	fmt.Println("int 2 called")
}

func init() {
	fmt.Println("int 3 called")
}

// Step 4: main() runs last, after all init() functions are finished.
func main() {
	fmt.Println("main function called")
}

/*  Terminal output (when the imported package has no printing init):

int 1 called
int 2 called
int 3 called
main function called

*/