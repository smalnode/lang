package main

import (
    "fmt"
    "os"
    "net"
    "time"
    "strconv"
)

func checkErr(err error) {
    if err != nil {
        fmt.Printf("error: %s\n", err)
        os.Exit(1)
    }
}

func main() {
    service := ":9090"
    tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
    checkErr(err)

    listener, err := net.ListenTCP("tcp", tcpAddr)
    checkErr(err)

    for {
        conn, err := listener.Accept()
        if err != nil {
            continue
        }
        go handleConn(conn)
    }

    os.Exit(0)
}

func handleConn(conn net.Conn) {
    conn.SetReadDeadline(time.Now().Add(2 * time.Second))
    defer conn.Close()

    request := make([]byte, 128)

    for {
        readLen, err := conn.Read(request)

        if err != nil {
            break
        }

        if readLen == 0 {
            break   // conn closed
        } else if string(request[:readLen]) == "timestamp" {
            daytime := strconv.FormatInt(time.Now().Unix(), 10)
            conn.Write([]byte(daytime))
        } else {
            daytime := time.Now().String()
            conn.Write([]byte(daytime))
        }

        request = make([]byte, 128)
    }
}
