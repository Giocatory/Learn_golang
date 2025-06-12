package main

import (
	"Learn_golang/my_modules"
	"fmt"
)

func main() {
	expression := my_modules.Variables(2, 2, 2)
	calc := expression(3)

	fmt.Println(calc)
}
