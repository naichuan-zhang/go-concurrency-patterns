package main

import (
	"fmt"
	"math/rand"
	"time"
)

// The boring function returns a channel that lets us communicate with the boring service it provides.
func boring(msg string) <-chan string { // Returns receive-only channel of strings.
	c := make(chan string)
	go func() { // We launch an anonymous goroutine from inside the function.
		// Simulating an infinite sender which puts a message to the channel.
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
	}()
	return c // Returns the channel to the caller.
}

func main() {
	// We can have more than one instances of the boring service.
	joe := boring("Joe") // Function returning a channel.
	ann := boring("Ann")

	for i := 0; i < 10; i++ {
		fmt.Println(<-joe)
		fmt.Println(<-ann)
	}

	fmt.Println("You're boring. I'm leaving.")
}
