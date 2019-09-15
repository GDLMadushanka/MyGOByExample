package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var ops uint64
	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		wg.Add(1)

		go func() {
			for c := 0; c < 10000; c++ {
				// To atomically increment the counter we use
				// AddUint64, giving it the memory address
				// of our ops counter with the & syntax.
				atomic.AddUint64(&ops, 1)
			}
			wg.Done()
		}()
	}

	fmt.Println("here i am ---------------")
	wg.Wait()

	fmt.Println("ops : ", ops)
}
