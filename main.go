package main

import (
	"fmt"
)

func main() {
	myArray := [...]int{1, 2, 3, 4, 5}

	for index, value := range myArray {
		fmt.Println(index, value)
	}
}
