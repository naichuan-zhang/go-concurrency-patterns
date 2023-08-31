package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Rewrite our original fanIn function.
// Only one goroutine is needed this time with select statement.
// NEW
func fanIn(input1, input2 <-chan string) <-chan string {
	c := make(chan string)
	go func() {
		for {
			select {
			case s := <-input1:
				c <- s
			case s := <-input2:
				c <- s
			}
		}
	}()
	return c
}

func main() {
	// Merge two channels into one
	c := fanIn(boring("Joe"), boring("Ann"))
	for i := 0; i < 10; i++ {
		fmt.Println(<-c)
	}
	fmt.Println("You're boring. I'm leaving.")
}

func boring(msg string) <-chan string {
	c := make(chan string)
	go func() {
		// An infinite loop to send messages to the channel.
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
	}()
	return c
}
