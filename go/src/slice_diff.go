package main

func main() {
    m := []strings{"A", "B", "E", "F"};
    n := []strings{"A", "B", "C", "D", "E", "F", "G"}
    a := Slice_diff(m, n)
    if a == nil {
        println(n, " contain ", m)
    }
}
