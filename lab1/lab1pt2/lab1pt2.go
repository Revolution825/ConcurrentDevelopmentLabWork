//Copyright (C) 2024 Diarmuid O'Neill

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
// Channel example with a go routine sending messages to main

package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan string) // creates a string channel c
	go count("sheep", c)

	for msg := range c {
		// msg, open := <-c // receives message from channel (sending and receiving are blocking operations) (channels can be used to synchronise/communicate between go routines)

		fmt.Println(msg)
	}
}

func count(thing string, c chan string) {
	for i := 0; i <= 5; i++ {
		c <- thing
		time.Sleep(time.Millisecond * 500)
	}

	close(c) // closes channel, only sender should close its channel
}
