package main

func main() {
    var i interface{}
    i = 10
    v, ok := i.(int)

    println("type of i is int:", ok, "\nvalue(int) of i is:", v)
}
