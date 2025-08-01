// add code for sample concurrency tester
package main

import (
	"fmt"
	"sync"
)

func main() {

	var wg sync.WaitGroup
	var mu sync.Mutex
	counter := 0

	// Number of goroutines to run
	numGoroutines := 10

	wg.Add(numGoroutines)

	for i := 0; i < numGoroutines; i++ {
		go func(id int) {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				mu.Lock()
				counter++
				mu.Unlock()
			}
			fmt.Printf("Goroutine %d finished\n", id)
		}(i)
	}

	wg.Wait()
	fmt.Printf("Final counter value: %d\n", counter)
}
