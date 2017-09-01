package main

import "fmt"

func main() {
    var array [100]int
    var slice1 = array[0:50]
    // var slice2 = array[0:100]
    for n := range slice1 {
        fmt.Printf( "%d\n", slice1[n] )
    }
    fmt.Printf( "%d\n", slice1[49] )
}
