package main

import "fmt"

func f() (ret int) {
	defer func() {
		ret++
	}()
	ret = 2
	return
}

func main() {
	i := f()
	fmt.Printf("%d\n", i)
}
