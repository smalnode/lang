package main

import (
    "fmt"
)

func main() {
    defer func() {
        for i := 0; i < 2; i++ {
            if x := recover(); x != nil {
                switch x.(type) {
                case string:
                    v, _ := x.(string)
                    fmt.Printf("recover[string]: %s\n", v)
                case int:
                    v, _ := x.(int)
                    fmt.Printf("recover[int]: %d\n", v)
                case float64:
                    v, _ := x.(float64)
                    fmt.Printf("recover[float64]: %f\n", v)
                case error:
                    v, _ := x.(error)
                    fmt.Printf("recover[eorror]: %s\n", v)
                default:
                    fmt.Printf("recover[unknown type]: %v\n", x)
                }
            }
        }
    } ()
    x := 1
    y := 0

    z := x / y
    println("z =", z)

}
