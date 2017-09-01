package main

import "fmt"

//import "sort"

type Sequence []int

func (s Sequence) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s Sequence) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s Sequence) Len() int {
	return len(s)
}

func main() {
	a := make([]int, 10)
	for i, _ := range a {
		a[i] = 10 - i
	}
	s := a
	//sort.Sort(s)
	//fmt.Println(s.Len())
	fmt.Println(s.Less(1, 2))
}
