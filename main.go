package main

import (
	"Learn_golang/my_modules"
	"fmt"
	"log"
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
}
