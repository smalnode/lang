package main

import "fmt"

type A struct {
	v    int
	name string
}

func main() {
	a := make([]int, 3)
	b := a[:1]

	var c [3]int

	b1 := b[0]
	showArray(b)
	chgArray(b)
	showArray(b)
	showArray(a)
	b2 := b[0]
	if b1 != b2 {
		fmt.Println("Slice was passed as reference type")
	} else {
		fmt.Println("Slice was passed as element type")
	}

	c1 := c[0]
	showArray2(c)
	chgArray2(c)
	showArray2(c)
	c2 := c[0]
	if c1 != c2 {
		fmt.Println("Array was passed as reference type")
	} else {
		fmt.Println("Array was passed as element type")
	}

	var s A
	s.v = 0
	p := s.v
	chgStruct(s)
	if s.v != p {
		fmt.Println("Struct was passed as referenced type")
	} else {
		fmt.Println("Struct was passed as element type")
	}

	fmt.Println(newA().v)
}

func showArray(a []int) {
	fmt.Println("array:")
	for i := 0; i != len(a); i++ {
		fmt.Println(a[i])
	}
	fmt.Println("array end")
}

func chgArray(a []int) {
	for i := 0; i != len(a); i++ {
		a[i] += 1
	}
}

func showArray2(a [3]int) {
	fmt.Println("array:")
	for i := 0; i != 3; i++ {
		fmt.Println(a[i])
	}
	fmt.Println("array end")
}

func chgArray2(a [3]int) {
	for i := 0; i != 3; i++ {
		a[i] += 1
	}
}

func chgStruct(s A) {
	s.v += 1
}

func newA() *A {
	return &A{v: 42, name: "answer of everything"}
}
