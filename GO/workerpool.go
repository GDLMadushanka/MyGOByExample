package main

import (
	"fmt"
	"time"
)

// one receving channel one sending channel
func worker(id int, jobs <-chan int, reslts chan<- int) {
	for j := range jobs {
		fmt.Println("worker ", id, " started job ", j)
		time.Sleep(time.Second)
		fmt.Println("worker ", id, " finished job ", j)
		reslts <- j * 2
	}
}

func main() {
	jobs := make(chan int, 100)
	results := make(chan int, 100)

	for w := 1; w <= 8; w++ {
		go worker(w, jobs, results)
	}

	for j := 1; j <= 8; j++ {
		jobs <- j
	}
	close(jobs)

	for a := 1; a <= 8; a++ {
		<-results
	}
}
