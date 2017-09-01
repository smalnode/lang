package main

func main() {
	var b B
	b.hello()
	b.hi()
}

type Ia interface {
	hello()
}

type Ib interface {
	Ia
	hi()
}

type A struct {
}

func (a *A) hello() {
	println("hello!")
	for {
		break
	}
}

func (a *B) hi() {
	println("hi!")
}

func (b *B) hello() {
	println("overloaded hello!")
}

type B struct {
	A
}
