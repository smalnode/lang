package main

import (
    "html/template"
    "os"
)

type UserProfile struct {
    UserName        string
}

func main() {
    t := template.New("Hello, {{.UserName}}!");
    t, _ = t.Parse("Hello, {{.UserName}}!\n");
    p := UserProfile{UserName: "smalacnt"}
    t.Execute(os.Stdout, p)
}
