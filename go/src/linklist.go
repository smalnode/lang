package main

import "fmt"

type Node struct {
	Value int
	Next  *Node
}

type LinkList struct {
	Head *Node
	Tail *Node
}

func (l *LinkList) appendNode(n *Node) *LinkList {
	if l.Head == nil {
		l.Head = n
	}

	if l.Tail != nil {
		l.Tail.Next = n
	}

	l.Tail = n

	return l
}

func (l *LinkList) appendValue(v int) *LinkList {
	n := new(Node)
	n.Value = v
	l.appendNode(n)
	return l
}

func (l *LinkList) show() {
	fmt.Print("Link List: ")
	for n := l.Head; n != nil; n = n.Next {
		fmt.Printf("%d -> ", n.Value)
	}
	fmt.Println("nil")
}

func main() {
	l := new(LinkList)
	for i := 100; i > 0; i -= 13 {
		l.appendValue(i)
		l.show()
	}
}
