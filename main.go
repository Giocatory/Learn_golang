package main

import (
	"Learn_golang/my_modules"
	"fmt"
	"log"
	"reflect"
	"strconv"
)

func main() {
	var someVar any

	fmt.Print("Введите строку: ")

	input, err := my_modules.InputString()

	if err != nil {
		log.Fatal(err)
	}

	isDigitOnly := my_modules.IsDigit(input)

	if isDigitOnly {
		fmt.Print("The string contains only digits: \n")
		var err error
		someVar, err = strconv.ParseInt(input, 0, 0)
		if err != nil {
			fmt.Println("Error parsing integer:", err)
			return
		}
		fmt.Print(someVar, " ", reflect.TypeOf(someVar))
	} else {
		fmt.Println("The string does not contain only digits.")
		someVar = input
		fmt.Print(someVar, " ", reflect.TypeOf(someVar))
	}

	WindowClose()
}

func WindowClose() (string, error) {
	input, err := my_modules.InputString()

	if err != nil {
		log.Fatal(err)
	}

	return input, nil
}
