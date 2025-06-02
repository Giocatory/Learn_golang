package main

import (
	"Learn_golang/my_modules"
	"fmt"
)

func main() {
	number := my_modules.GetRandomNumber(10)
	fmt.Println(number)
}
