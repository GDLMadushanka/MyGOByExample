package main

import (
	"fmt"
	"time"
)

func main() {
	timer1 := time.NewTimer(2 * time.Second)

	// The <-timer1.C blocks on the timer’s channel C
	// until it sends a value indicating that the timer expired.
	<-timer1.C
	fmt.Println("timer 1 expired")

	timer2 := time.NewTimer(time.Second)
	go func() {
		<-timer1.C
		fmt.Println("timer 2 expired")
	}()

	// If you just wanted to wait, you could have used time.
	// Sleep. One reason a timer may be
	// useful is that you can cancel the timer before it expires
	stop2 := timer2.Stop()
	if stop2 {
		fmt.Println("timer 2 stopped")
	}

}
