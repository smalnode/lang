package main

import "math"
import "fmt"

func main() {
	a := math.NaN()
	fmt.Println("NaN > 0 ? ", a > 0, " NaN <= 0 ? ", a <= 0)
	fmt.Println("math.NaN() = ", math.NaN())
}
