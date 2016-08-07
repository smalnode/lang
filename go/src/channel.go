package main

import "fmt"
import "time"

func main() {
	chn := make(chan int)

	go func() {
		for i := 0; i != 100; i++ {
			chn <- i
		}
	}()

	go func() {
		for i := 0; i != 100; i++ {
			fmt.Printf("%d\n", <-chn)
		}
	}()

	time.Sleep(1e9)
}
