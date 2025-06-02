package main

import (
	"Learn_golang/my_modules"
	"fmt"
)

func main() {
	number, err := my_modules.GetRandomNumber(10)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(number)
}
