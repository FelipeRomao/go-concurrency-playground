package main

import (
	"fmt"
	"time"
)

func countWithChannel(thing string, c chan string) {
	for i := 1; i <= 5; i++ {
		c <- thing
		time.Sleep(time.Millisecond * 500)
	}
	close(c)
}

func main() {
	c := make(chan string, 2)
	c <- "Hello"
	c <- "World"

	msg := <-c
	fmt.Println(msg)

	msg = <-c
	fmt.Println(msg)
}
