package main

import (
	"fmt"
	"net/http"
)

func sayhelloFunc(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("value:", v, "len(v) =", len(v))
	}
	fmt.Fprintf(w, "Hello, world!")
}

func main() {
	http.HandleFunc("/", sayhelloFunc)
	http.ListenAndServe(":9090", nil)
}
