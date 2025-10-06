package main

import (
	"fmt"
	"sync"
	"time"
)

// make struct containing channel
// add init, acquire and release
type semaphore struct {
	theCounter chan struct{}
}

func NewSemaphore(maxRoutines int) *semaphore {
	return &semaphore{
		theCounter: make(chan struct{}, maxRoutines),
	}
}

func Acquire(sem *semaphore) {
	sem.theCounter <- struct{}{}
}

func Release(sem *semaphore) {
	<-sem.theCounter
}

func main() {
	maxGoroutines := 5
	semaphore := NewSemaphore(maxGoroutines)

	var wg sync.WaitGroup
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			Acquire(semaphore)
			defer Release(semaphore)

			// Simulate a task
			fmt.Printf("Running task %d\n", i)
			time.Sleep(2 * time.Second)
		}(i)
	}
	wg.Wait()
}
