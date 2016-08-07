package main

import (
	"fmt";
	"net/http";
	"strings";
	"log";
    "html/template"
)

func sayhelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println(r.Form)
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key: ", k)
		fmt.Println("val: ", strings.Join(v, ""))
	}
	fmt.Fprintf(w, "Hello astaxie!")
}

func login(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
    fmt.Println("method: ", r.Method)
    if r.Method == "GET" {
        t, _ := template.ParseFiles("login.html")
        t.Execute(w, nil)
    } else {
        fmt.Println("username: %s",
            template.HTMLEscapeString(r.Form["username"][0]))
        fmt.Println("password: %s",
            template.HTMLEscapeString(r.Form["password"][0]))
        template.HTMLEscape(w, []byte(r.Form.Get("username")))
    }
}

func main() {
	http.HandleFunc("/", sayhelloName)
	http.HandleFunc("/login", login)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("LintenAndServe: ", err)
	}
}
