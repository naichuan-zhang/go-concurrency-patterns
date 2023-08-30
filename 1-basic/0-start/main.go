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

func main() {

	// The boring function runs on forever, like a boring party guest.
	//boring("boring!")

	// The go statement runs the function as usual, but doesn't make the caller wait.
	// It launches a goroutine.
	// The functionality is analogous to the & on the end of a shell command.
	go boring("boring!")
}
