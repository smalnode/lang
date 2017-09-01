package main

import "fmt"

func main() {
	var array [10]int
	for i := 0; i != 10; i++ {
		array[i] = i + 1
		fmt.Printf("array[%d] = %d\n", i, array[i])
	}
	slice := array[2:]
	slice = append(slice, slice...)
	copy(slice, array[0:])
	for _, element := range slice {
		fmt.Printf("%d\n", element)
	}
}
