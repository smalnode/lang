package main

import "fmt"

type T int

func (T) m() {
	fmt.Println("m()")
}

func main() {
	var x **T
	var y *T
	var z T = 100
	y = &z
	x = &y
	(*x).m()
	y.m()
	(**x).m()
	z.m()
}
