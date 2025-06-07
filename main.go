package main

import (
	"fmt"
)

func main() {
	m := map[string]int{
		"hello": 5,
		"world": 0,
	}

	v, ok := m["hello"]
	fmt.Println(v, ok)

	// WindowClose()
}

// func WindowClose() (string, error) {
// 	input, err := my_modules.InputString()

// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	return input, nil
// }
