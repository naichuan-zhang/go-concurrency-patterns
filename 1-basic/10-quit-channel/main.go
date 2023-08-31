package main

import (
	"fmt"
	"math/rand"
	"time"
)

func boring(msg string, quit chan bool) <-chan string {
	c := make(chan string)
	go func() {
		for i := 0; ; i++ {
			select {
			case c <- fmt.Sprintf("%s %d", msg, i):
				// do nothing
			case <-quit:
				fmt.Println("quit")
				return
			}
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
	}()
	return c
}

func main() {
	// We can use the quit channel to turn this around and tell Joe to stop
	// when we're tired of listening to him.
	quit := make(chan bool)
	c := boring("Joe", quit)
	for i := rand.Intn(10); i >= 0; i-- {
		fmt.Println(<-c)
	}
	quit <- true
}
