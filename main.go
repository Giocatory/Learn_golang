package main

import (
	"fmt"
	"reflect"
)

func main() {
	a := []string{"a", "b", "c", "d", "e", "f"}

	b := make([]string, len(a[:4]))

	copy(b, a[:4])

	b[0] = "abc"

	fmt.Println(a, " ", reflect.TypeOf(a))
	fmt.Println(b, " ", reflect.TypeOf(b))

	// WindowClose()
}

// func WindowClose() (string, error) {
// 	input, err := my_modules.InputString()

// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	return input, nil
// }
