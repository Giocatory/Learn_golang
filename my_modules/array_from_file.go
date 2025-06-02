package my_modules

import (
	"bufio"
	"os"
)

func GetArrayFromFile(filename string) ([]string, error) {
	var fileText []string

	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(file)
	if scanner.Err() != nil {
		return nil, scanner.Err()
	}

	for scanner.Scan() {
		fileText = append(fileText, scanner.Text())
	}

	defer file.Close()

	return fileText, nil
}
