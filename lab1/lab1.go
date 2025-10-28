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
// Example of using a wait group to wait for a go routine to finish

package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup // A counter we created for active go routines
	wg.Add(1)

	go func() { // the go keyword creates a go routine (this is an annonymous function that is invoked immediately)
		count("sheep")
		wg.Done()
	}()

	wg.Wait() // This will block until all active go routines are complete i.e. wg = 0
}

func count(thing string) {
	for i := 0; i <= 5; i++ {
		fmt.Println(i, thing)
		time.Sleep(time.Millisecond * 500)
	}
}
