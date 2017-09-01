package main

import (
	"time"
)

func say(s string) <-chan int {
	c := make(chan int)
	go func() {
		time.Sleep(300 * time.Millisecond)
		panic(s)
		c <- 0
	}()
	return c
}

func main() {
	say("hello")
	say("world")
}
