package main

import (
	"fmt"
	"math/rand"
	"time"
)

func boring(msg string) <-chan string {
	c := make(chan string)
	go func() {
		for i := 0; i < 10; i++ {
			c <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
	}()
	return c
}

func main() {
	c := boring("Joe")

	for {
		select {
		case s := <-c:
			fmt.Println(s)
		// The time.After function returns a channel that blocks for the specified duration.
		// After the interval, the channel delivers the current time, once.
		case <-time.After(1 * time.Second): // We have a timeout defined for each message.
			fmt.Println("You're too slow.")
			return
		}
	}
}
