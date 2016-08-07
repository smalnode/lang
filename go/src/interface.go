package main

import "fmt"

type Int64 int

type Show interface {
	show()
}

func (v Int64) show() {
	fmt.Println(v)
}

func (v Int64) less100() bool {
	return v < 100
}

func print2(v Show) {
	v.show()
}

func main() {
	var a Int64 = 100
	print2(a)
	fmt.Println(a.less100())
}
