package main

import "fmt"

type Lang struct {
    Name string
    Year int
    URL string
    Next *Lang
}

func (l *Lang) Append(l2 *Lang) {
    l.Next = l2
}

func (l *Lang) AppendNew() {
    l.Next = new(Lang)
}

func initLang(lang_p *Lang) {
    lang_p = &Lang{}
}

func (l Lang) Print() {
    fmt.Printf("Name: %s\nYear: %d\nURL: %s\n", l.Name, l.Year, l.URL)
}

func main() {
    lang := Lang{ "Go", 2009, "http://golang.org", nil }
    lang.AppendNew()
    lang.Next.Print()
    var lang_p *Lang
    initLang(lang_p)
    lang_p.Print()
}
