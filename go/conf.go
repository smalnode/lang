package main

import (
    "fmt"
    "encoding/xml"
    "os"
    "io/ioutil"
)

type H1 struct {
    XMLName         xml.Name        `xml:"h1"`
    H2              H2              `xml:"h2"`
}

type H2 struct {
    XMLName         xml.Name        `xml:"h2"`
    H4s             []H4            `xml:"h4"`
    H6s             []float64       `xml:"h5>h6"`
    Version         string          `xml:"version,attr"`
}

type H4 struct {
    XMLName         xml.Name        `xml:"h4"`
    Value           string          `xml:",innerxml"`
    Version         string          `xml:"class,attr"`
}

func main() {
    file, err := os.Open("./conf.xml")
    if err != nil {
        fmt.Fprintf(os.Stderr, "error: %v\n", err)
        os.Exit(-1)
    }
    defer func() {
        file.Close()
    } ()

    v := H1{}

    data, err := ioutil.ReadAll(file)
    if err != nil {
        fmt.Fprintf(os.Stderr, "error: %v\n", err)
        os.Exit(-1)
    }

    err = xml.Unmarshal(data, &v)
    if err != nil {
        fmt.Fprintf(os.Stderr, "error: %v\n", err)
        os.Exit(-1)
    }

    fmt.Printf("%v\n", v)
}
