package main

import "fmt"

func f() (ret int) {
	defer func() {
		ret++
	}()
	return
}

func main() {
	i := f()
	fmt.Printf("%d\n", i)
}
