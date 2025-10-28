//Copyright (C) 2025 Diarmuid O'Neill

// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

// Author: Diarmuid O'Neill (C00282898@setu.ie)
// Example of worker pool using channels and go routines

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
