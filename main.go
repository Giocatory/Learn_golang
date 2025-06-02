package main

import (
	"Learn_golang/my_modules"
	"fmt"
)

func main() {
	resultInt := my_modules.GetAverageIntArray(1, 2, 3, 4, 5)
	resultFloat := my_modules.GetAverageFloatArray(1.0, 2.0, 3.0, 4.0, 5.0)

	fmt.Println(resultInt, resultFloat)
}
