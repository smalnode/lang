package main

import "fmt"

type T struct {
    a int
    b float32
    c string
}

func main() {
    t := &T{7, -3.25, "abc\tdef"}
    fmt.Printf("%T\n", t)
    fmt.Printf("%+T\n", t)
    fmt.Printf("%#T\n", t)
    tp := &t
    (*tp).a = -7
    tpp := &tp
    fmt.Printf("%v\n", t)
    fmt.Printf("%+v\n", t)
    fmt.Printf("%#T\n", &tpp)
}
