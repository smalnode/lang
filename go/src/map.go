package main

func main() {
    m := map[string]int{}
    m["wdeqin"] = 0
    m["trxvel"] = 1
    m["trcvel"] = 2

    for a, b := range m {
        println("a = ", a, ", b = ", b)
    }
}
