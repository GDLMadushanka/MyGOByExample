package main

import (
	"fmt"
	"time"
)

func main() {

	requests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		requests <- i
	}
	close(requests)

	// This limiter channel will receive a value every 200 milliseconds.
	limiter := time.Tick(200 * time.Millisecond)

	for req := range requests {
		<-limiter
		fmt.Println("request", req, time.Now())
	}

	/*
		We may want to allow short bursts of requests in our rate
		limiting scheme while preserving the overall rate limit.
		We can accomplish this by buffering our limiter channel. This
		burstyLimiter channel will allow bursts of up to 3 events
	*/

	burstlimiter := make(chan time.Time, 3)
	for i := 0; i < 3; i++ {
		burstlimiter <- time.Now()
	}

	//Every 200 milliseconds we’ll try to add a new
	//value to burstyLimiter, up to its limit of 3.
	go func() {
		for t := range time.Tick(2 * time.Second) {
			burstlimiter <- t
		}
	}()

	burstRequests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		burstRequests <- i
	}
	close(burstRequests)

	for req := range burstRequests {
		<-burstlimiter
		fmt.Println("request burst ", req, time.Now())
	}
}
