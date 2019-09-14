package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan string, 1)
	go func() {
		time.Sleep(2 * time.Second)
		c1 <- "result 1"
	}()

	select {
	case res := <-c1:
		fmt.Println(res)
	case <-time.After(time.Second):
		fmt.Println("timeout 1")
	}

	go func() {
		time.Sleep(2 * time.Second)
		c1 <- "result 2"
	}()

	select {
	case res := <-c1:
		fmt.Println(res)
		// res := <-c1 awaits the result and <-Time.After
		// awaits a value to be sent after the timeout of 1s
	case <-time.After(time.Second * 3):
		fmt.Println("timeout 2")
	}
}
