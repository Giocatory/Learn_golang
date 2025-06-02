package main

import (
	"Learn_golang/my_modules"
	"fmt"
	"reflect"
)

func main() {
	stringResult, _ := my_modules.GetMax("a", "b", "c", "d", "e")
	intResult, _ := my_modules.GetMax(1, 2, 3, 4, 5)
	floatResult, _ := my_modules.GetMax(1.1, 2.2, 3.3, 4.4, 5.5)

	fmt.Println(stringResult, reflect.TypeOf(stringResult)) // e string
	fmt.Println(intResult, reflect.TypeOf(intResult))       // 5 int
	fmt.Println(floatResult, reflect.TypeOf(floatResult))   // 5.5 float64
}
