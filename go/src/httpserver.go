package main

import (
	"fmt"
	"net/http"
)

func MyResponse(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	name := "world"
	if n := r.FormValue("name"); n != "" {
		name = n
	}
	w.Write([]byte(fmt.Sprintf("Hello, %s", name)))
}

func main() {
	http.HandleFunc("/", MyResponse)
	http.ListenAndServe(":10010", nil)
}
