package main

import "fmt"

func main() {
	a := make([]int, 10)
	for i := 0; i != len(a); i++ {
		a[i] = 10 - i
	}
	showSlice(a)
	qsort(a)
	showSlice(a)
}

func showSlice(a []int) {
	fmt.Println("slice: ")
	for _, v := range a {
		fmt.Printf("%d\t", v)
	}
	fmt.Println()
}

func qsort(a []int) {
	sort(a, 0, len(a))
}

// s start pos
// e end pos + 1
func sort(a []int, s int, e int) {
	if s >= e {
		return
	}
	p := a[s]
	i := s
	j := s + 1
	for j < e {
		if a[j] < p {
			a[i+1], a[j] = a[j], a[i+1]
			i++
			j++
		} else {
			j++
		}
	}

	a[s], a[i] = a[i], a[s]
	sort(a, s, i)
	sort(a, i+1, e)
}
