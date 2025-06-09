package main

import (
	"fmt"
	"strings"
)

func main() {
	num := "five"

	num = strings.Replace(num, "e", "i", -1)

	fmt.Println(num)
}
