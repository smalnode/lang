package main

import (
	"fmt"
)

func main() {
	c := [...]int{1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Printf("%%v%v\n", c)
	fmt.Printf("%%q%q\n", c)
	fmt.Printf("%%s%s\n", c)

	var a, b int
	fmt.Scanf("%d", &a)
	fmt.Scanf("%d", &b)
	fmt.Printf("a = %d b = %d a * b = %d\n", a, b, a*b)
}
