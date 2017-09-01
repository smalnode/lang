package main

import (
    "crypto/rand"
    "math/big"
    "fmt"
)

func main() {
    c := 10
    b := make([]byte, c)
    _, err := rand.Read(b)

    if err != nil {
        fmt.Println("error: ", err)
        return
    }

    max := big.NewInt(1000)
    for cnt := 0; cnt < 10; cnt++ {
        i, _ := rand.Int(rand.Reader, max)
        println(i.Int64())
    }
}
