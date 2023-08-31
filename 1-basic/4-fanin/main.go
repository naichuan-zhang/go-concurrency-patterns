package main

import (
	"fmt"
	"math/rand"
	"time"
)

// These programs make Joe and Ann count in lockstep.
// We can instead use a fan-in function to let whoever is ready to talk.
func fanIn(input1, input2 <-chan string) <-chan string {
	c := make(chan string)
	go func() {
		for {
			//v1 := <-input1
			//c <- v1

			// For simplicity
			c <- <-input1
		}
	}()
	go func() {
		for {
			//v2 := <-input2
			//c <- v2

			// For simplicity
			c <- <-input2
		}
	}()
	return c
}

func fanInSimple(inputs ...<-chan string) <-chan string {
	c := make(chan string)
	for _, ci := range inputs {
		go func(cv <-chan string) {
			for {
				c <- <-cv
			}
		}(ci)
	}
	return c
}

func main() {
	// Merge two channels into one
	//c := fanIn(boring("Joe"), boring("Ann"))
	c := fanInSimple(boring("Joe"), boring("Ann"))
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
