//Barrier.go Template Code
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

//--------------------------------------------
// Author: Joseph Kehoe (Joseph.Kehoe@setu.ie)
// Created on 30/9/2024
// Modified by: Diarmuid O'Neill (C00282898@setu.ie)
//--------------------------------------------

package main

import (
	"context"
	"fmt"
	"sync"
	"time"

	"golang.org/x/sync/semaphore"
)

// Place a barrier in this function --use Mutex's and Semaphores
func doStuff(goNum int, wg *sync.WaitGroup, bar *chan bool, lock *sync.Mutex, count *int, totalRoutines int) bool {
	time.Sleep(time.Second)
	fmt.Println("Part A", goNum)
	//we wait here until everyone has completed part A

	lock.Lock()
	*count++

	if *count == totalRoutines {
		lock.Unlock()
		for i := 0; i < totalRoutines-1; i++ {
			*bar <- true
		}

	} else {
		lock.Unlock()
		<-*bar
	}
	fmt.Println("PartB", goNum)
	wg.Done()
	return true
}

func main() {
	totalRoutines := 10
	count := 0
	var wg sync.WaitGroup
	barrier := make(chan bool)
	wg.Add(totalRoutines)
	//we will need some of these
	ctx := context.TODO()
	var theLock sync.Mutex
	var threadLock sync.Mutex
	sem := semaphore.NewWeighted(int64(totalRoutines))
	theLock.Lock()
	sem.Acquire(ctx, 1)
	for i := range totalRoutines { //create the go Routines here
		go doStuff(i, &wg, &barrier, &threadLock, &count, totalRoutines)
	}
	sem.Release(1)
	theLock.Unlock()

	wg.Wait() //wait for everyone to finish before exiting
}
