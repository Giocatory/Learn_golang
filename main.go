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
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text: ")
	text, err := reader.ReadString('\n')

	if err != nil {
		log.Fatal(err)
	}

	text = strings.TrimSpace(text)

	result := my_modules.GetStringLength(text)

	fmt.Print("Length your text: ", result)
}
