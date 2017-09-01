package main

import "fmt"

func main() {
	p := []uint{30, 35, 15, 5, 10, 20, 25}
	mat(p[:])
}

func mat(p []uint) {
	n := len(p) - 1
	c := make([][]uint, n+1)
	ck := make([][]uint, n+1)
	for i := 0; i <= n; i++ {
		c[i] = make([]uint, n+1)
		ck[i] = make([]uint, n+1)
	}
	for i := n; i >= 1; i-- {
		for j := i + 1; j <= n; j++ {
			if i == 0 || j == 0 || i == j {
				c[i][j] = 0
			} else if j == i+1 {
				c[i][j] = p[i-1] * p[i] * p[j]
			} else {
				var mc uint = 1 << 62
				var tc uint
				for k := i; k < j; k++ {
					tc = c[i][k] + c[k+1][j] + p[i-1]*p[k]*p[j]
					if i == 4 && j == 6 {
						fmt.Printf("i = %d\tj = %d\tk = %d\tc = %d\n", i, j, k, tc)
					}
					if mc > tc {
						mc = tc
						ck[i][j] = uint(k)
					}
				}
				c[i][j] = mc
			}
		}
	}
	fmt.Println()
	for i := 0; i <= n; i++ {
		for j := 0; j <= n; j++ {
			fmt.Printf("%d\t", c[i][j])
		}
		fmt.Println()
		fmt.Println()
	}
	fmt.Println()
	fmt.Println()
	for i := 0; i <= n; i++ {
		for j := 0; j <= n; j++ {
			if j == i+1 {
				ck[i][j] = uint(i)
			}
			fmt.Printf("%d\t", ck[i][j])
		}
		fmt.Println()
		fmt.Println()
	}
	fmt.Printf("Minimum cost = %d\n", c[1][n])
	fmt.Print("Solution: ")
	solution(ck, 1, uint(n))
	fmt.Println()
}

func solution(s [][]uint, i uint, j uint) {
	if i == j {
		fmt.Print("A", i)
	} else {
		fmt.Print("(")
		solution(s, i, s[i][j])
		solution(s, s[i][j]+1, j)
		fmt.Print(")")
	}
}
