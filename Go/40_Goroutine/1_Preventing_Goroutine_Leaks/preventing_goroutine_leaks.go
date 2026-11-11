/*

                  Introduction


      • Goroutines are easy to create.
      • Garbage collector does not clean Goroutines.
      • Responsibility of the creator to clean.


                Termination Paths


1. The job assigned to the goroutine is now completed.
2. There was an error when the goroutine was executed.
3. The parent goroutine needs to terminate it.



*/

package main

import (
	"fmt"
	"sync"
	"time"
)

func someWork(done <-chan struct{},wg *sync.WaitGroup) <-chan int {
	ch := make(chan int)
	go func(){
		defer wg.Done()
		select {
		case ch <-1:
		case <-done:
		}
		fmt.Println("Goroutne finished")
	}()
	return ch
}

func main() {
	wg := &sync.WaitGroup{}
	wg.Add(1)

	done := make(chan struct{})
	someWork(done, wg)
	time.Sleep(1 * time.Second)
	close(done)
	wg.Wait()
}