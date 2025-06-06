package main

import (
	"Learn_golang/my_modules"
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	file, err := my_modules.ReadCSV("static/statistics.csv")

	if err != nil {
		log.Fatal(err)
	}

	for id, data := range file {
		fmt.Printf("ID: %s; Name: %s; Age: %s\n", id, data["name"], data["age"])
		fmt.Println()
	}

	fmt.Println(file["1"]["name"]) // Михаил

	WindowClose()
}

func WindowClose() (string, error) {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.Replace(input, "\n", "", -1)

	return input, nil
}
