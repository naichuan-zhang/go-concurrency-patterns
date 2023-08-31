package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Send a channel on a channel, making goroutine wait its turn.
// Receive all messages, then enable them again by sending on a private channel.

// Message We define a message type that contains a channel for the reply.
type Message struct {
	str  string
	wait chan bool
}

func fanInSimple(inputs ...<-chan Message) <-chan Message {
	c := make(chan Message)
	for i := range inputs {
		cm := inputs[i]
		go func() {
			for {
				c <- <-cm
			}
		}()
	}
	return c
}

func boring(msg string) <-chan Message {
	c := make(chan Message)
	// Shared between all messages.
	waitForIt := make(chan bool)
	go func() {
		// An infinite loop that sends messages to the channel.
		for i := 0; ; i++ {
			c <- Message{
				str:  fmt.Sprintf("%s %d", msg, i),
				wait: waitForIt,
			}
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)

			// Everytime the goroutine sends message.
			// This code waits until the value is received.
			<-waitForIt
		}
	}()
	return c
}

func main() {
	// Merge two channels into one.
	c := fanInSimple(boring("Joe"), boring("Ann"))
	for i := 0; i < 10; i++ {
		msg1 := <-c
		fmt.Println(msg1.str)
		msg2 := <-c
		fmt.Println(msg2.str)

		// Assign waiting channels signals to let them continue.
		msg1.wait <- true
		msg2.wait <- true
	}
	fmt.Println("You're boring. I'm leaving.")
}
