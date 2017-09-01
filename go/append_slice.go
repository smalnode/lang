package main

import "fmt"

func main() {
	x := []float32{1, 2, 3}
	y := []float32{4, 5, 6}
	x = append(x, y...)
	fmt.Println(x)
}
