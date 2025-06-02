package main

import (
	"Learn_golang/my_modules"
	"fmt"
)

func main() {
	textFromFile := my_modules.GetArrayFromFile("static/poem.txt")

	for _, line := range textFromFile {
		fmt.Println(line)
	}
}
