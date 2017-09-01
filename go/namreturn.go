package main

func namereturn() ( i int ) {
    i = 0
    for j := 0; j !=100; j++ {
        i++
    }
    return
}

func main() {
   println( namereturn() )
}
