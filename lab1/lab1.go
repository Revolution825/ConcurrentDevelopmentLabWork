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
