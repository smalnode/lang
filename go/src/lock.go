package main

import "sync"
import "fmt"
import "time"

func main() {
    var global_sum int
    global_sum = 0

    var mutex sync.Mutex

    go func() {
        mutex.Lock()
        global_sum += 100
        mutex.Unlock()
    }()

    go func() {
        mutex.Lock()
        global_sum += 100
        mutex.Unlock()
    }()

    mutex.Lock()
    global_sum += 50
    mutex.Unlock()

    time.Sleep(1e9)

    fmt.Printf("global_sum = %d\n", global_sum)
}
