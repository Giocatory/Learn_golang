package main

import (
	"Learn_golang/my_modules"
	"fmt"
)

func main() {
	stringResult, _ := my_modules.GetMax("a", "b", "c", "d", "e")
	intResult, _ := my_modules.GetMax(1, 2, 3, 4, 5)
	floatResult, _ := my_modules.GetMax(1.1, 2.2, 3.3, 4.4, 5.5)

	fmt.Println(stringResult)
	fmt.Println(intResult)
	fmt.Println(floatResult)
}
