package main

import (
	"fmt"
)

func main() {
	jobs := make(chan int, 100) // creates channel with a buffer of 100
	results := make(chan int, 100)

	go worker(jobs, results) // creates a worker

	for i := 0; i < 100; i++ {
		jobs <- i
	}
	close(jobs)

	for j := 0; j < 100; j++ {
		fmt.Println(<-results)
	}
}

func worker(jobs <-chan int, results chan<- int) { // jobs only send, results only receives
	for n := range jobs {
		results <- fib(n)
	}
}

func fib(n int) int {
	if n <= 1 {
		return n
	}

	return fib(n-1) + fib(n-2)
}
