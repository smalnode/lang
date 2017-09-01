package main

import "fmt"
import "math/rand"

func main() {
	n := 30
	nn := int32(n)
	p := make([]int, n)
	for i := 0; i < n; i++ {
		p[i] = int(rand.Int31n(nn))
	}
	fmt.Println(p)
	fmt.Println(lmis(p))
}

func lmis(p []int) []int {
	// c[i] = len of LIS end with p[i]
	// t[i] = the processor of p[i] in LIS end with p[i]
	n := len(p)
	c := make([]int, n)
	t := make([]int, n)

	for i := 0; i < n; i++ {
		c[i] = 1
		t[i] = i
	}

	for i := 1; i < n; i++ {
		for j := 0; j < i; j++ {
			if p[j] <= p[i] && c[i] < c[j]+1 {
				c[i] = c[j] + 1
				t[i] = j
			}
		}
	}

	max := c[0]
	maxi := 0

	for i := 1; i < n; i++ {
		if max < c[i] {
			max = c[i]
			maxi = i
		}
	}

	s := make([]int, max)
	k := max - 1

	for k >= 0 {
		s[k] = p[maxi]
		k--
		maxi = t[maxi]
	}

	return s

}
