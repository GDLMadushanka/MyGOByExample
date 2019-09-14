/*

To wait for multiple goroutines to finish,
we can use a wait group.
*/

package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

//WaitGroup must be passed to functions by pointer
func worker(id int, wg *sync.WaitGroup) {
	fmt.Println("worker ", id, " started")
	var a int = rand.Intn(5)
	fmt.Println("sleep ", a, " seconds")
	time.Sleep(time.Duration(a) * time.Second)
	fmt.Println("worker ", id, " done")
	// Notify the WaitGroup that this worker is done.
	wg.Done()
}

func main() {
	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		// Launch several goroutines and increment the WaitGroup counter for each.
		wg.Add(1)
		go worker(i, &wg)
	}

	wg.Wait()
}
