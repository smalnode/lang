package main

func chgArray(array []int) {
    for i := 0; i < len(array); i++ {
        array[i] *= 2
    }
}

func main() {
    var array []int
    array = make([]int, 100)
    for i := 99; i >= 0; i-- {
        array[i] = 100 - i
    }

    chgArray(array[25:75])

    for i := 0; i < 100; i++ {
        println(array[i])
    }
    
}
