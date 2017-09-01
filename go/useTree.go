package main

import (
    "tree"
    "math/rand"
)


func main() {
    t := tree.Tree{nil}
    for cnt := 0; cnt != 25; cnt++ {
        t.Insert(int(rand.Int31()))
    }
    println(t.Root == nil)
    t.PreTra()
}
