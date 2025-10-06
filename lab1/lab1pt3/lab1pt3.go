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
