package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan string)
	c2 := make(chan string)
	c3 := make(chan any)

	go func() {
		c1 <- "Hello c1!"
	}()
	go func() {
		c2 <- "Hello c2!"
	}()
	go func() {
		fmt.Printf("received %v from c3\n", <-c3)
	}()

	for {
		// The select statement provides another way to handle multiple channels.
		// It's like a switch, but each case is a communication:
		// - All channels are evaluated.
		// - Selection blocks until one communication cna proceed, which then does.
		// - If multiple can proceed, select chooses pseudo-randomly.
		// - A default clause, if present, executes immediately if no channel is ready.
		select {
		case v1 := <-c1:
			fmt.Printf("received %v from c1\n", v1)
		case v2 := <-c2:
			fmt.Printf("received %v from c2\n", v2)
		case c3 <- 23:
			fmt.Printf("sent %v to c3\n", 23)
		case <-time.After(5 * time.Second):
			fmt.Printf("You're boring. I'm leaving.\n")
			return
			//default:
			//	fmt.Printf("no one was ready to communicate\n")
			//	time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
	}
}
