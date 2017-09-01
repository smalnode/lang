package main

func inita(a []int) {
	for i, _ := range a {
		a[i] = 1
	}
}

func doublea(a []int) {
	for i, e := range a {
		a[i] = 2 * e
	}
}

func printa(a []int) {
	for i, e := range a {
		println("[", i, "] = ", e)
	}
}

func main() {
	a := make([]int, 100)
	inita(a)
	doublea(a)
	doublea(a)
	printa(a)
}
