package main

import (
	"fmt"
	"math/rand"
	"time"
)

func boring(msg string) {
	for i := 0; ; i++ {
		fmt.Println(msg, i)
		time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
	}
}

// When main returns, the program exits and takes the boring function down with it.
// We can hang around a little, and on the way show that both main and the launched goroutine are running.
func main() {
	// the main goroutine is finished after this line is executed
	// it doesn't wait for boring goroutine finished
	// thus, the program exits, and we don't see anything
	go boring("boring!")

	fmt.Println("I'm listening.")
	// let the main goroutine to wait for 2s
	// thus, boring goroutine can be executed and have some messages printed out
	time.Sleep(2 * time.Second)
	// main goroutine exits and program exits
	fmt.Println("You're boring. I'm leaving.")

	// the main goroutine and boring goroutine doesn't communicate with each other
	// we need to introduce channel...
}
