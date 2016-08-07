package main

import "fmt"

const (
	LU = 1 << iota
	LEFT
	UP
)

func main() {
	x := []byte{'1', '0', '0', '1', '0', '1', '0', '1'}
	y := []byte{'0', '1', '0', '1', '1', '0', '1', '1', '0'}
	printSeq(x)
	printSeq(y)
	printSeq(lcs(x, y))
}

func lcs(x, y []byte) (z []byte) {
	m := len(x)
	n := len(y)
	c := make([][]byte, m+1)
	s := make([][]byte, m+1)
	for i := 0; i <= m; i++ {
		c[i] = make([]byte, n+1)
		s[i] = make([]byte, n+1)
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			ii := i - 1
			jj := j - 1
			if x[ii] == y[jj] {
				c[i][j] = c[i-1][j-1] + 1
				s[i][j] = LU
			} else {
				if c[i][j-1] > c[i-1][j] {
					c[i][j] = c[i][j-1]
					s[i][j] = UP
				} else {
					c[i][j] = c[i-1][j]
					s[i][j] = LEFT
				}
			}
		}
	}
	z = make([]byte, c[m][n])
	solution(x, z, s)
	return
}

func solution(x, z []byte, c [][]byte) {
	i := len(c) - 1
	j := len(c[1]) - 1
	k := len(z) - 1
	for k >= 0 {
		switch c[i][j] {
		case LU:
			z[k] = x[i-1]
			k--
			i--
			j--
		case LEFT:
			i--
		case UP:
			j--
		}
	}
}

func printSeq(s []byte) {
	fmt.Print("[")
	for _, v := range s {
		fmt.Printf("%c, ", v)
	}
	fmt.Print("\b\b]\n")
}
