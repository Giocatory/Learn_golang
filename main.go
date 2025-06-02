package main

import (
	"Learn_golang/my_modules"
	"fmt"
	"log"
)

func main() {
	textFromFile, err := my_modules.GetArrayFromFile("static/poem.txt")

	if err != nil {
		log.Fatal(err)
		return
	}

	for _, line := range textFromFile {
		fmt.Println(line)
	}
}
