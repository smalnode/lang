package main

import (
    "fmt"
    "os"
    "net"
)

func checkErr(err error) {
    if err != nil {
        fmt.Fprintf(os.Stderr, "error: %v\n", err)
        os.Exit(1)
    }
}
func main() {
    if len(os.Args) != 2 {
        fmt.Fprintf(os.Stdout, "Usage: %s host:port\n", os.Args[0])
        return
    }

    addr, err := net.ResolveUDPAddr("udp4", os.Args[1])
    checkErr(err)

    conn, err := net.DialUDP("udp", nil, addr)
    checkErr(err)

    _, err = conn.Write([]byte("UDP request"))
    checkErr(err)

    response := make([]byte, 512)
    read_len, addr, err := conn.ReadFromUDP(response)
    checkErr(err)

    fmt.Println(string(response[:read_len]))
}
