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
