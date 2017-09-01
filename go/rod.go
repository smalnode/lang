package main

import "fmt"

func main() {
	p := []int{0, 0, 5, 8, 9, 10, 17, 17, 20, 24, 30}
	var l int
	fmt.Println("Input a positive integer: ")
	fmt.Scanf("%d", &l)
	if l <= 0 {
		return
	}
	solve(l, p[:])
}

func solve(l int, p []int) {
	r := make([]int, l+1)
	s := make([]int, l+1)
	r[0] = p[0]
	for j := 1; j <= l; j++ {
		max := -(1 << 63)
		for i := 1; i <= j && i < len(p); i++ {
			if max < p[i]+r[j-i] {
				max = p[i] + r[j-i]
				s[j] = i
			}
		}
		r[j] = max
	}
	fmt.Println(r[len(r)-1])
	solution(l, s)
}

func solution(l int, s []int) {
	fmt.Print("Solution: ")
	for l > 0 {
		fmt.Printf("%d\t", s[l])
		l = l - s[l]
	}
	fmt.Println()
}
