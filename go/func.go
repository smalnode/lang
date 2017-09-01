package main

import "fmt"

func rec(i int) int {
    if i == 10 {
        return i
    }
    return rec(i + 1)
}

func main() {
    fmt.Printf("%d", rec(1))
}

