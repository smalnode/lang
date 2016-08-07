package main

import (
    "net/http"
    "net/rpc"
    "rpcArg"
    "fmt"
    "os"
)

func main() {
    arith := new(rpcArg.Arith)
    rpc.Register(arith)
    rpc.HandleHTTP()
    err := http.ListenAndServe(":9090", nil)
    if  err != nil {
        fmt.Fprintf(os.Stderr, "error: %v\n", err)
        os.Exit(1)
    }
}
