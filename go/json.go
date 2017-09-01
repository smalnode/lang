package main

import (
    "fmt"
    "os"
    "io/ioutil"
    "encoding/json"
)

type Server struct {
    ServerName  string      `json:"serverName"`
    ServerIP    string      `json:"serverIP"`
}

type Servers struct {
    Servers     []Server    `json:"servers"`
}

func main() {
    file, err := os.Open("./conf.json")
    if err != nil {
        fmt.Printf("open file ./conf.json failed: \n%v\n", err)
        os.Exit(-1)
    }

    data, err := ioutil.ReadAll(file)
    if err != nil {
        fmt.Printf("read file ./conf.json failed: \n%v\n", err)
        os.Exit(-1)
    }

    var v Servers
    err = json.Unmarshal(data, &v)
    if err != nil {
        fmt.Printf("unmarshal file ./conf.json failed: \n%v\n", err)
        os.Exit(-1)
    }
    fmt.Printf("%v\n", v)

    output, err := json.MarshalIndent(&v, "", "    ")
    if err != nil {
        fmt.Printf("marshal struct v failed: \n%v\n", err)
        os.Exit(-1)
    }
    fmt.Printf("%v\n", string(output))

}
