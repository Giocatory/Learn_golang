package main

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	var someVar any

	fmt.Print("Введите строку: ")

	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	isDigitOnly := true
	for _, char := range input {
		if !unicode.IsDigit(char) {
			isDigitOnly = false
			break
		}
	}

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
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	return input, nil
}
