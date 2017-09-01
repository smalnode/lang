package main

type T chan uint64

func M(f uint64) (in, out T) {
	in = make(T, 100)
	out = make(T, 100)
	go func(in, out T, f uint64) {
		for {
			out <- f * <-in
		}
	}(in, out, f)
	return in, out
}

func min(xs []uint64) uint64 {
	m := xs[0]
	for i := 1; i < len(xs); i++ {
		if xs[i] < m {
			m = xs[i]
		}
	}
	return m
}

func main() {
	F := []uint64{2, 3, 5, 7}
	var n = len(F)

	x := uint64(1)
	ins := make([]T, n)
	outs := make([]T, n)
	xs := make([]uint64, n)
	for i := 0; i < n; i++ {
		ins[i], outs[i] = M(F[i])
		xs[i] = x
	}

	for i := 0; i < 20; i++ {
		for i := 0; i < n; i++ {
			ins[i] <- x
		}

		for i := 0; i < n; i++ {
			if xs[i] == x {
				xs[i] = <-outs[i]
			}
		}

		x = min(xs)
		println(x)
	}
}
