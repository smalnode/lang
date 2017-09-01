package main

type IGreeting interface {
	sayHello()
}

type Bird struct{}

func (b *Bird) sayHello() {
	println("Hi, I'm flying!")
}

func sayHello() {
	println("Hi, I'm flying!")
}

func main() {
	var b Bird
	b.sayHello()
	(&b).sayHello()
}
