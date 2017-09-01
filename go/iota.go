package main

import "fmt"

func main() {
	const (
		A int = 1 << (10 * iota)
		B
		C
		D
		E
		F
	)

	fmt.Printf("%d\t%d\t%d\t%d\t%d\t%d\t", A, B, C, D, E, F)
}
