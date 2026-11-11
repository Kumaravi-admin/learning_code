package main

import (
	"fmt"
	"sync"
)

// The object we want to reuse (our "cup").
type Cup struct {
	ID int
}

// Counts how many cups were really created.
var created = 0

// Step 1: Create the pool.
// New is called ONLY when the pool is empty, to make a brand new object.
var pool = sync.Pool{
	New: func() any {
		created++
		fmt.Println("New cup banaya, total created:", created)
		return &Cup{ID: created}
	},
}

func main() {

	for i := 1; i <= 3; i++ {

		// Step 2: Get() takes a cup from the pool.
		// Pool empty? It calls New and gives a new cup.
		// Get returns type "any", so .(*Cup) converts it back to *Cup (type assertion).
		cup := pool.Get().(*Cup)

		fmt.Println("Customer", i, "using cup", cup.ID)

		// Step 3: Put() returns the cup, so the next customer can reuse it.
		pool.Put(cup)
	}
}

/*  Terminal output:

New cup banaya, total created: 1
Customer 1 using cup 1
Customer 2 using cup 1
Customer 3 using cup 1

*/