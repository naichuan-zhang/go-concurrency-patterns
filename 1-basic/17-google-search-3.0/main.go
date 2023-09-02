package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Result string

type Search func(query string) Result

var (
	Web1   = fakeSearch("web 1")
	Web2   = fakeSearch("web 2")
	Web3   = fakeSearch("web 3")
	Image1 = fakeSearch("image 1")
	Image2 = fakeSearch("image 2")
	Image3 = fakeSearch("image 3")
	Video1 = fakeSearch("video 1")
	Video2 = fakeSearch("video 2")
	Video3 = fakeSearch("video 3")
)

// We can simulate the search function, much as we simulated conversation before.
func fakeSearch(kind string) Search {
	return func(query string) Result {
		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
		return Result(fmt.Sprintf("%s result for %q\n", kind, query))
	}
}

// First To avoid discarding results from slow servers by replicating the servers.
// Send requests to multiple replicas, and use the first response.
func First(query string, replicas ...Search) Result {
	c := make(chan Result)
	searchReplica := func(i int) {
		c <- replicas[i](query)
	}
	for i := range replicas {
		go searchReplica(i)
	}
	// Always waits for 1 time after receiving the result.
	return <-c
}

func Google(query string) (results []Result) {
	c := make(chan Result)
	go func() {
		c <- First(query, Web1, Web2, Web3)
	}()
	go func() {
		c <- First(query, Image1, Image2, Image3)
	}()
	go func() {
		c <- First(query, Video1, Video2, Video3)
	}()

	// A global timeout.
	// It ignores the result from the server that taking response greater than 80ms.
	timeout := time.After(80 * time.Millisecond)
	for i := 0; i < 3; i++ {
		select {
		case result := <-c:
			results = append(results, result)
		case <-timeout:
			fmt.Println("timed out")
			return
		}
	}
	return
}

func main() {
	rand.Seed(time.Now().UnixNano())
	start := time.Now()
	results := Google("golang")
	elapsed := time.Since(start)
	fmt.Println(results)
	fmt.Println(elapsed)
}
