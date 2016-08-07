package main

import (
	"fmt";
	"strconv"
)

type Human struct {
	name string
	age int
	phone string
}

func (h Human) String() string {
	return "@" + h.name + " - " + strconv.Itoa(h.age) +
			" years - " + h.phone
}

func main() {
	h := Human{"Tong", 35, "222-201-130-33"}
	fmt.Println(h)
}
