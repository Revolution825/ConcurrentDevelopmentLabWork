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
// Channel example with a buffered channel

package main

import (
	"fmt"
)

func main() {
	c := make(chan string, 2) // creates a string channel c
	c <- "hello"              // send does not proceed until something is ready to receive (create a seperate go routine or buffered channel)

	msg := <-c
	fmt.Println(msg) // this example will not work
}
