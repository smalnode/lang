package main

import "time"
import "fmt"

func main() {
    c := time.Tick(1 * time.Second)
    i := 'a'
    for now := range c {
        fmt.Printf("%v %s \n", now.String()[:19], string(i))
        i++
    }
}
