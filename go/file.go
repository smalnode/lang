package main

import (
    "bufio"
    "fmt"
    "os"
)

func main() {
    file, err := os.Open("file.go")
    if err != nil {
        println(err)
        return
    }
    defer file.Close()
    lineReader := bufio.NewReaderSize(file, 20)
    lineNumber := 1

    for line, isPrefix, err := lineReader.ReadLine();
        err == nil;
        line, isPrefix, err = lineReader.ReadLine() {
            lineNumber++
            os.Stdout.Write(line)
        if isPrefix {
            for {
                line, isPrefix, _ = lineReader.ReadLine();
                os.Stdout.Write(line)
                if !isPrefix {
                    break
                }
            }
        }
        fmt.Println()
    }

}
