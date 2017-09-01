package main

// int multiply(int lhs, int rhs) {
//     return lhs * rhs;
// }
import "C"
import "fmt"

func main() {
	fmt.Println(C.multiply(6, 7))
}
