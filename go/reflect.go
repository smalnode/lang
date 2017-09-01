package main

import "fmt"
import "reflect"

func main() {
	var x float64 = 3.4
	//v := reflect.ValueOf(x)

    /*
	fmt.Println("type: ", v.Type())
	fmt.Println("kind is float64: ", v.Kind() == reflect.Float64)
	fmt.Println("value: ", v.Float())
    */

	fmt.Println("value of x: ", x)
	p := reflect.ValueOf(&x)
	pv := p.Elem()
	pv.SetFloat(7.1)
	fmt.Println("value of x: ", x)
    f(&x)
	fmt.Println("value of x: ", x)
}

func f(x *float64) {
    p := reflect.ValueOf(x)
    pv := p.Elem()
    pv.SetFloat(16.0)
}
