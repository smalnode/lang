/*
 * File: dir.go
 * 5/3/2014
 * Edited by Wangdeqin
 * List filenames in a specific path
 */

package main

import (
    "os"
    "errors"
    "fmt"
)

func GetFilenames(path string) ([]string, error) {
    dir, err := os.Open(path)
    checkErr(err)
    defer dir.Close()

    state, err := dir.Stat()
    checkErr(err)

    if !state.IsDir() {
        return nil, errors.New(fmt.Sprintf("%s is not a dir", path))
    }

    fis, err := dir.Readdir(-1)
    checkErr(err)

    filenames := make([]string, 0)

    for _, fi := range fis {
        filenames = append(filenames, fi.Name())
    }

    return filenames, nil
}

func checkErr(err error) {
    if err != nil {
        panic(err.Error())
    }
}

func usage(prog_name string) {
    fmt.Printf("Usage: %s <path>\n", prog_name)
    os.Exit(0)
}

func main() {
    if (len(os.Args) < 2) {
        usage(os.Args[0])
    }

    filenames, err := GetFilenames(os.Args[1])
    checkErr(err)

    for _, filename := range filenames {
        fmt.Println(filename[:len(filename) - 4])
    }
}
