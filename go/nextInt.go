package main

func nextInt( b []byte, i int ) ( int, int ) {
    x := 0
    for ; i < len( b ); i++ {
        x = x * 10 + int ( b[ i ] ) - '0'
    }

    return x, i
}

func main () {
    a := []byte{ '1', '2', '3', '4' }
    var x int
    for i := 0; i < len( a ); {
        x, i = nextInt( a, i )
        println( x )
    }
}
