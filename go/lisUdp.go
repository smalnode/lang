package main

import (
    "os"
    "net"
    "fmt"
    "time"
)

func checkErr(err error) {
    if err != nil {
        fmt.Printf("Error: %v\n", err)
        os.Exit(1)
    }
}

func main() {
    service := ":9090"
    addr, err := net.ResolveUDPAddr("udp", service)
    checkErr(err)

    conn, err := net.ListenUDP("udp", addr)
    checkErr(err)

    for {
        handleClient(conn)
    }
}

func handleClient(conn *net.UDPConn) {
    buffer := make([]byte, 512)

    _, addr, err := conn.ReadFromUDP(buffer)
    checkErr(err)

    daytime := time.Now().String()
    conn.WriteToUDP([]byte(daytime), addr)
}
