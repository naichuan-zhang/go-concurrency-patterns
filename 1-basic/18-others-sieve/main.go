package main

// Generate sends the sequence 2, 3, 4, ... to channel
func Generate(ch chan<- int) {
	for i := 2; ; i++ {
		ch <- i
	}
}

// Filter copies the values from channel 'in' to channel 'out',
// and removes those divisible by 'prime'.
func Filter(in <-chan int, out chan<- int, prime int) {
	for {
		// Receive value from 'in'.
		i := <-in
		if i%prime != 0 {
			// Send 'i' to 'out'.
			out <- i
		}
	}
}

// The prime sieve: Daisy-chain filter processes.
func main() {
	c := make(chan int)
	go Generate(c)
	for i := 0; i < 10; i++ {
		prime := <-c
		print(prime, "\n")
		c1 := make(chan int)
		go Filter(c, c1, prime)
		c = c1
	}
}
