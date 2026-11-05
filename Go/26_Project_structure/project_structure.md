# Simple Package

<!-- Heading of the slide: a simple Go package example -->

## Project Structure

<!-- Folder tree: shows which files live inside the package folder -->

    countries/
    ├── go.mod              # Defines the module name and tracks dependencies
    ├── go.sum              # Stores hashes of exact dependency versions (lock file)
    ├── countries.go        # Main source code of the package
    └── countries_test.go   # Test file (ends with _test.go, so Go treats it as a test)

Module name: github.com/codeandlearn199/countries

<!-- Module name = repository path. Set once with: go mod init github.com/codeandlearn199/countries -->

Package name: countries

<!-- Package name = first line in countries.go: package countries (matches the folder name) -->

<!----------------------------------------------------------------------------------------------------->

# Multi-File Simple Package

<!-- Heading of the slide: one package split across multiple files -->

## Project Structure

<!-- Folder tree: all .go files in this folder belong to the same package (countries) -->

    countries/
    ├── go.mod                       # Defines the module name and tracks dependencies
    ├── go.sum                       # Stores hashes of exact dependency versions (lock file)
    ├── countries.go                 # Main source code of the package
    ├── countries_test.go            # Tests for countries.go
    ├── countrycodes.go              # Country code logic, kept in its own file for clarity
    ├── countrycodes_test.go         # Tests for countrycodes.go
    ├── countrycontinent.go          # Country continent logic, kept in its own file
    └── countrycontinent_test.go     # Tests for countrycontinent.go

Module name: github.com/codeandlearn199/countries

<!-- Module name = repository path. Set once with: go mod init github.com/codeandlearn199/countries -->

Package name: countries

<!-- Every .go file in this folder starts with: package countries
     Files share the same package, so they can use each other's functions
     (even unexported ones) without any import. -->

<!----------------------------------------------------------------------------------------------------->

# A Simple Executable

<!-- Heading of the slide: a package that builds into a runnable program -->

## Project Structure

<!-- Folder tree: package main makes this an executable (it has a main() function as entry point) -->

    list/
    ├── go.mod               # Defines the module name and tracks dependencies
    ├── go.sum               # Stores hashes of exact dependency versions (lock file)
    ├── main.go              # Entry point: contains func main(), where the program starts
    ├── list.go              # List logic, kept in its own file for clarity
    ├── list_test.go         # Tests for list.go
    ├── sortlist.go          # Sorting logic for the list
    └── sortlist_test.go     # Tests for sortlist.go

Module name: github.com/codeandlearn199/list

<!-- Module name = repository path. Set once with: go mod init github.com/codeandlearn199/list -->

Package name: main

<!-- Every .go file in this folder starts with: package main
     "main" is special: it tells Go to build a runnable program instead of a library.
     Run it with: go run .   or build it with: go build -->

<!----------------------------------------------------------------------------------------------------->

# Multi-Package Project

<!-- Heading of the slide: one module that contains more than one package -->

## Project Structure

<!-- Folder tree: each sub-folder (auth, refresh) is its own separate package -->

    oauth2/
    ├── go.mod                 # Defines the module name and tracks dependencies
    ├── go.sum                 # Stores hashes of exact dependency versions (lock file)
    ├── oauth2.go              # Main source code of the root package (oauth2)
    ├── oauth2_test.go         # Tests for oauth2.go
    ├── auth/                  # Sub-folder = separate package: auth
    │   ├── auth.go            # Source code of the auth package
    │   └── auth_test.go       # Tests for the auth package
    └── refresh/               # Sub-folder = separate package: refresh
        ├── refresh.go         # Source code of the refresh package
        └── refresh_test.go    # Tests for the refresh package

Module name: github.com/codeandlearn199/oauth2

<!-- Module name = repository path. Set once with: go mod init github.com/codeandlearn199/oauth2
     Import paths for sub-packages = module name + folder path, for example:
     github.com/codeandlearn199/oauth2/auth
     github.com/codeandlearn199/oauth2/refresh -->

Package name: \*

<!-- "*" means there is no single package name here, because this project has many packages.
     Each folder has its own: oauth2 (root), auth, refresh.
     Rule: one folder = one package. -->

<!----------------------------------------------------------------------------------------------------->

# Internal Packages

<!-- Heading of the slide: packages that are hidden from outside the module -->

## Project Structure

<!-- Folder tree: the "internal" folder is special in Go -->

    oauth2/
    ├── go.mod                 # Defines the module name and tracks dependencies
    ├── go.sum                 # Stores hashes of exact dependency versions (lock file)
    ├── oauth2.go              # Main source code of the root package (oauth2)
    ├── oauth2_test.go         # Tests for oauth2.go
    ├── internal/              # Special folder: packages inside can only be imported from within this module
    │   └── client/            # Internal package: client
    │       ├── client.go      # Source code of the client package
    │       └── client_test.go # Tests for the client package
    ├── auth/                  # Public package: auth (other projects can import it)
    │   ├── auth.go            # Source code of the auth package
    │   └── auth_test.go       # Tests for the auth package
    └── refresh/               # Public package: refresh (other projects can import it)
        ├── refresh.go         # Source code of the refresh package
        └── refresh_test.go    # Tests for the refresh package

Module name: github.com/codeandlearn199/oauth2

<!-- Module name = repository path. Set once with: go mod init github.com/codeandlearn199/oauth2
     The Go compiler blocks other modules from importing:
     github.com/codeandlearn199/oauth2/internal/client
     but oauth2, auth and refresh (inside this module) can still use it. -->

Package name: \*

<!-- "*" means many packages here: oauth2 (root), client, auth, refresh.
     Rule: one folder = one package. -->

<!----------------------------------------------------------------------------------------------------->

# Multiple Executables

<!-- Heading of the slide: one module that builds more than one runnable program -->

## Project Structure

<!-- Folder tree: each folder inside cmd/ is a separate executable with its own main.go -->

    rest/
    ├── go.mod                       # Defines the module name and tracks dependencies
    ├── go.sum                       # Stores hashes of exact dependency versions (lock file)
    ├── cmd/                         # Convention: holds all executables (programs) of the project
    │   ├── server/                  # Executable 1: the server program
    │   │   └── main.go              # Entry point of the server (package main, func main())
    │   └── migratior/               # Executable 2: the migration program
    │       └── main.go              # Entry point of the migration tool (package main, func main())
    ├── internal/                    # Special folder: packages inside can only be imported from within this module
    │   └── handler/                 # Internal package: handler
    │       ├── handler.go           # Source code of the handler package
    │       └── handler_test.go      # Tests for the handler package
    └── migration/                   # Public package: migration (other projects can import it)
        ├── migration.go             # Source code of the migration package
        └── migration_test.go        # Tests for the migration package

Module name: github.com/codeandlearn199/rest

<!-- Module name = repository path. Set once with: go mod init github.com/codeandlearn199/rest
     Build or run each executable separately:
     go run ./cmd/server
     go run ./cmd/migratior -->

Package name: \*

<!-- "*" means many packages here: main (server), main (migratior), handler, migration.
     Rule: one folder = one package. Each folder in cmd/ is its own package main. -->
