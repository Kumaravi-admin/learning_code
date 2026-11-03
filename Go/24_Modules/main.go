/*
                   Go Modules

• Go modules are used for managing dependencies in your Go projects.
• Go modules uses go.mod and go.sum files to track dependencies of your Go project.
• One Go module per repository is preferred.
• The go.mod also defines your Go module.


            Initializing Go Modules

• The go mod init command is used to initialize a module.
• Name of your module is usually the whole path of your repository that hosts the module.
    ◦ "github.com/username/repo"
• Private modules can be accessed as long as your git client can access them.


                    Adding Modules

• The go get command is used to add module dependencies.
• By default it pulls the latest version.
• The go mod tidy command can be used to cleanup modules not used in your project.

                  The go.sum File

• The go.sum serves as lock file.
• It stores cryptographic hashes for  specific modules versions.
• The go.mod and go.sum file together ensure that same module versions are  downloaded every time.

                  The go list

Use go list all command to list of all imported packages in your modules.
You can supply the -m flag to only list modules imported.


                 The go mod tidy

• The go mod tidy command as the name suggest is used for tidying up  dependencies of your Go project.
• Running go mod tidy remove any unused  modules from your project.
• Running go mod tidy will also add any  missing modules that is imported but missing from go.mod file.
*/

package main

import (
	"fmt"

	// External package (not part of Go's standard library).
	// It comes from the module we download in Step 2.
	"github.com/fatih/color"
)

func main() {

	// Step 1: Create a module (run once, in your project folder).
	// This creates the go.mod file, which defines your module and tracks dependencies.
	// >> go mod init github.com/Kumaravi-admin/first-module

	// Step 2: Download an external module.
	// This adds it to go.mod and creates go.sum (hashes of the exact versions).
	// >> go get github.com/fatih/color

	// Step 3 (optional): Download a specific version instead of the latest.
	// >> go get github.com/fatih/color@v1.15.0

	// Step 4: List all packages used by your project.
	// Add the -m flag to list only modules.
	// >> go list all
	// >> go list -m all

	// Step 5: Clean up dependencies.
	// go mod tidy does two jobs:
	//   - removes modules that are in go.mod but not used in your code
	//   - adds modules that your code imports but are missing from go.mod
	// >> go mod tidy

	// Step 6: Use the imported package.
	// color.RedString returns the text wrapped in red color codes.
	str := color.RedString("Hello World!")
	fmt.Println(str)

}

/*  Terminal output:

Hello World!   (printed in red color)

*/