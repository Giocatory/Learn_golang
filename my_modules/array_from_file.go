package my_modules

import (
	"bufio"
	"log"
	"os"
)

func GetArrayFromFile(filename string) []string {
	var fileText []string

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	if scanner.Err() != nil {
		log.Fatal(scanner.Err())
	}

	for scanner.Scan() {
		fileText = append(fileText, scanner.Text())
	}

	defer file.Close()

	return fileText
}
