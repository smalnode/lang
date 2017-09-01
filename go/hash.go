package main

import (
    "fmt"
    "hash/crc32"
    "io"
    "os"
)

func main() {
    h := crc32.NewIEEE()
    w := io.MultiWriter( h, os.Stdout )
    fmt.Fprintf( w, "Hello, world!\n" )
    fmt.Printf( "hash=%#x\n", h.Sum32() )
}
