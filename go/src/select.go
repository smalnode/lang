package main

import (
	"time"
)

func intChan(base int) chan int {
	c := make(chan int)

	go func() {
		for i := 1; i < 10; i++ {
			c <- (i * base)
			time.Sleep(100 * time.Millisecond)
		}
		close(c)
	}()
	return c
}

func main() {
	c1 := intChan(1)
	/*
		c2 := intChan(-1)
		var t int
	*/
	for {
		if v, ok := <-c1; ok {
			println(v)
		} else {
			break
		}
	}

}
