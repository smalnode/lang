package main

import "fmt"

func main() {
    c := make(chan int)

    go getInt(c)
    go getInt(c)
    go getInt(c)

    total := 0
    total += <-c
    total += <-c
    total += <-c
    fmt.Println("Total = ", total)
}

func getInt(c chan int) {
    c <- 10
}
