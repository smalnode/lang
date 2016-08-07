package main

type A interface{}

func main() {
	switch {
	case 3 < 4:
		println("3<4")
	case 1 < 2:
		println("1<2")
	}

	var i int64 = 10

	var a A = i

	switch a.(type) {
	case int64:
		v, _ := a.(int64)
		println(v)
	}
}
