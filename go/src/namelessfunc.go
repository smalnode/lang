package main

import "fmt"

func main() {
    a := func( i int ) {
        fmt.Printf( "Hello, this message from nameless function, %d\n", i )
    }

    a( 10 )

    fmt.Printf( "%T\n", a )
}
