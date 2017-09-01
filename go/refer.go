package main

import "fmt"

func main() {
    var i *int
    i = new(int)
    *i = 100
    fmt.Println(*i)
    chg(i)
    fmt.Println(*i)
}

func chg(p *int) {
    *p  = -100
}
