package main

import (
    "os"
    "fmt"
    "net"
)

func main() {
    if len(os.Args) != 2 {
        fmt.Printf("Usage: %s host:port\n", os.Args[0])
        return
    }

    server := os.Args[1]
    tcpAddr, err := net.ResolveTCPAddr("tcp4", server)
    checkErr(err)

    conn, err := net.DialTCP("tcp", nil, tcpAddr)
    checkErr(err)

    _, err = conn.Write([]byte("timestamp"))
    checkErr(err)

    response := make([]byte, 128)

    for {
        read_len, err := conn.Read(response)
        if read_len == 0 || err != nil {
            break
        }
        fmt.Println(string(response[:read_len]))
    }

    os.Exit(0)
}

func checkErr(err error) {
    if err != nil {
        fmt.Printf("Error: %s\n", err)
        os.Exit(1)
    }
}
