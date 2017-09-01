package main
import "fmt"

func main(){
	i := 0
	PrintI : fmt.Printf("%d\n", i)
	i++
	if i != 10 {
		goto PrintI
	}
}
