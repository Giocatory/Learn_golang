package main

import (
	"fmt"
)

func main() {
	num := "five"

	fmt.Println(num)

	changeVal(&num)

	fmt.Println(num)
}

func changeVal(value *string) string {
	*value = "new value"
	return *value
}
