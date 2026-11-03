package main

import (
	"fmt"

	// Import path = module name (from go.mod) + folder path of the package.
	// module: github.com/Kumaravi-admin/package-example
	// folder: math
	"github.com/Kumaravi-admin/package-example/math"
	// Alias: give the package a different name.
	// With this you would call internal.Add(...) instead of math.Add(...)
	// internal "github.com/Kumaravi-admin/package-example/math"
	// Period import: brings exported names into this file, so you can call Add(...) directly.
	// . "github.com/Kumaravi-admin/package-example/math"
	// Underscore import: only runs the package's init() side effects, you can't use its names.
	// _ "github.com/Kumaravi-admin/package-example/math"
)

func main() {

	// Call an exported function using packageName.FunctionName.
	// Add starts with an uppercase letter, so it is accessible here.
	sum := math.Add(1, 2)
	fmt.Println(sum) // Output: 3

	// These would NOT compile:
	// math.subtract(4, 2) // subtract is lowercase = unexported (private to the math package)
}


/*		Folder structure:

25_Packages/
├── go.mod          -> module name
├── main.go         -> package main (entry point)
└── math/
    ├── math.go     -> package math
    └── int.go      -> package math (same package, second file)

*/