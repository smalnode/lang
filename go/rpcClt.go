package main

import (
    "net/rpc"
    "fmt"
    "log"
    "rpcArg"
)

func main() {
    client, err := rpc.DialHTTP("tcp", "localhost:9090")
    if err != nil {
        log.Fatal("dialing: ", err)
    }

    args := rpcArg.Args{7, 8}
    var reply int
    err = client.Call("Arith.Multiply", args, &reply)
    if err != nil {
        log.Fatal("arith error: ", err)
    }

    fmt.Printf("Arith: %d * %d = %d\n", args.A, args.B, reply)

}
