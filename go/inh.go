package main

import "fmt"

type Human struct {
	name string
	age int
	phone string
}

type Student struct {
	Human
	school string
	loan float32
}

type Employee struct {
	Human
	company string
	money float32
}

func (h Human) Sing (lyrics string) {
	fmt.Println("La la, la la la, la la la la... ", lyrics)
}

func (h Human) Guzzle (beerStein string) {
	fmt.Println("Guzzle Guzzle Guzzle... ", beerStein)
}

func (h Human) SayHi() {
	fmt.Printf("Hi, I am %s, you can call me on %s.\n", h.name, h.phone);
}

func (e Employee) SayHi() {
	fmt.Printf("Hi, I'm %s. I work at %s, call me on %s.\n", e.name,
				e.company, e.phone)
}

func (e *Employee) SpendSalary(amount float32) {
	e.money -= amount
}

func (s *Student) BorrowMoney(amount float32) {
	s.loan += amount
}

type Men interface {
	SayHi()
	Sing(lyrics string)
	Guzzle(beerStein string)
}

type  YoungChap interface {
	SayHi()
	Sing(lyrics string)
	BorrowMoney(amount float32)
}

type ElderlyGent interface {
	SayHi()
	Sing(lyrics string)
	SpendSalary(amount float32)
}
func main() {
	mark := Student{Human{"Mark", 25, "222-222-YYYY"}, "SCUT", 0}
	sam := Employee{Human{"Sam", 30, "222-222-XXXX"}, "Google", 1000}
	tom := Employee{Human{"Tom", 30, "332-222-XXXX"}, "Golang Inc.", 200}

	var i Men

	i = mark
	i.Sing("Blaze...")

	i = tom
	i.Guzzle("Gua Gua Gua...")

	x := make([]Men, 3)
	x[0], x[1], x[2] = mark, sam, tom

	for _, v := range x {
		v.SayHi()
	}
}
