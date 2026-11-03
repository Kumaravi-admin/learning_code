/*
                     Packages

• Packages are collection of source code files that logically grouped.
• First line of code is the package declaration in a Go file.
• Only one package can live inside a directory with exception of test packages.

                     Exports

• Go uses the case of constants, variables, functions, types, etc to decide if they are exported or not.
• Constants, variables, functions, types, etc that start with uppercase letters are exported.


                       Imports

• To import package you need the package path which is the module name plus the relative directory path of the package.
• The Go compiler does not allow packages to be imported but not use them.
• Each code file needs to have its own imports.

                    Package Naming

• Idiomatically a package directory should have the same name as the package.
• There are exceptions to this rule,example when you have different version of the same package.
• Package name should represents the functionality being offered by the package.
• Ambiguous package name such as utils are discouraged.
• Idiomatic to have multiple packages with clear name instead of having one package with ambiguous name.


            More About Package Imports

• Package alias can be used to give a  different name to package import.
• Package import with a period imports exported entities from package in current namespace.
• Underscore import is used side  effects like initialization.

*/

// Step 1: Run "go mod init <module-name>" once in the PROJECT ROOT (where go.mod lives).
// Step 2: Create a folder named "math" and put this file inside it.
// Step 3: The first line must be the package declaration.
// The package name matches the folder name (idiomatic Go).
package math

// Exported constant: starts with an uppercase letter, so other packages can use math.PI
const PI = 3.14159

// Exported function: uppercase "A", so main.go can call math.Add(...)
func Add(x, y int) int {
	return x + y
}

// Unexported function: lowercase "s", so it can only be used INSIDE the math package.
// Other packages (like main) cannot call it.
func subtract(x, y int) int {
	return x - y
}