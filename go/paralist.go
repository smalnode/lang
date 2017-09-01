package main

import "fmt"

func paralist2( arg ...int ){
    for _, n := range arg {
        fmt.Printf( "%d\n", n )
    }
}

func paralist( arg ...int ) {
    for _, n := range arg {
        fmt.Printf( "%d\n", n )
    }

    paralist2( arg[:3]... )
}

func main() {
    paralist( 1, 2, 3, 4, 5, 6, 7, 8 )
}
