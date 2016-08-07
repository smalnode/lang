package main

import "fmt"

//import "reflect"

func main() {
	const N = 10
	var a [N][N]int
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			a[i][j] = N*i + j
		}
	}

	for _, s := range a {
		println("len(s) = ", len(s))
		for _, e := range s {
			fmt.Println(e)
		}
	}

}
