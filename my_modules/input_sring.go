package my_modules

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func InputString() (string, error) {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')

	if err != nil {
		return "", fmt.Errorf("%v", err.Error())
	}

	input = strings.TrimSpace(input)

	return input, nil
}
