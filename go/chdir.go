/*
 * File: chdir.go
 * 5/3/2014
 * Edited by Wangdeqin
 * Test Chdir in pkg os
 */

package main

import (
    "os"
    "fmt"
)

func usage(prog_name string) {
    fmt.Printf("Usage: %s <path>\n", prog_name)
    os.Exit(1)
}

func checkErr(err error) {
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error: %s\n", err.Error())
        os.Exit(1)
    }
}

func main() {
    if len(os.Args) < 2 {
        usage(os.Args[0])
    }

    err := os.Chdir(os.Args[1])
    checkErr(err)

    mark, err := os.Create("Mark.txt")
    defer mark.Close()
    checkErr(err)

    mark.WriteString("mark\n")
}
