package main

import "fmt"

func deferDemo() {
    for i := 0; i !=10; i++ {
        defer fmt.Printf( "%d\n", i )
    }
}

func main() {
    deferDemo()
}
