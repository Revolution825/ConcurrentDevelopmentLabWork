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
// This program demonstrates the use of semaphores in Go, to limit concurrent
// goroutines when performing parallel tasks.

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
