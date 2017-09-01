package main

import "fmt"

func main() {
	var minimum int64 = -(1 << 63)
	var maximum int64 = minimum - 1
	fmt.Printf("minimum int64 = %d\nmaximum int64 = %d\n", minimum, maximum)
}
