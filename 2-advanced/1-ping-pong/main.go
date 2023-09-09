package main

import (
	"fmt"
	"time"
)

type Ball struct {
	// Global count
	hits int
}

func main() {
	// Pointer to make sure it always points to the same ball.
	table := make(chan *Ball)
	go player("ping", table)
	go player("pong", table)

	// Game on...
	// The first ball is tossed.
	// Comment out the line below to cause deadlock.
	table <- new(Ball)
	// Main program waits for 1s before exit.
	time.Sleep(1 * time.Second)
	// Game over...
	// Snatch the ball.
	<-table
}

func player(name string, table chan *Ball) {
	for {
		// Player grabs the ball.
		ball := <-table
		ball.hits++
		fmt.Println(name, ball.hits)
		time.Sleep(100 * time.Millisecond)
		// Send the ball back to the adversary.
		table <- ball
	}
}
