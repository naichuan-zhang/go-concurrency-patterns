package main

import (
	"fmt"
	"math/rand"
	"time"
)

// the boring function adding a parameter to accept a string channel
func boring(msg string, c chan string) {
	for i := 0; ; i++ {
		// send the string format value to channel `c`
		// it also waits for receiver to be ready
		c <- fmt.Sprintf("%s %d", msg, i)
		time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
	}
}

// A channel connects the main and boring goroutines, so they can communicate.
// When the main function executes <-c, it will wait for a value to be sent.
// Similarly, when the boring function executes c <- value, it waits for a receiver to be ready.
// A sender and receiver must both be ready to play their part in the communication. Otherwise we wait until they are.
// Thus channels both communicate and synchronize.
func main() {
	// init channel
	c := make(chan string)
	// start our boring goroutine
	go boring("boring!", c)

	for i := 0; i < 5; i++ {
		// <-c receives the value from boring function
		// <-c blocks until a value to be sent
		fmt.Printf("You say: %q\n", <-c)
	}
	fmt.Println("You're boring. I'm leaving.")
}
