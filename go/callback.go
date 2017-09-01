package main

import "fmt"

func callback( y int, f func( int ) int ) int {
    return f( y )
}

func f( i int ) int  {
    fmt.Printf( "%d\n", i )
    return i
}

func main() {
    println( callback( 100, f ) )
}
