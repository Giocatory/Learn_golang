package main

import (
	"Learn_golang/my_modules"
	"fmt"
	"log"
	"slices"
)

func main() {
	a := []string{"a", "b", "c", "d", "e", "f"}
	b := a[:2]
	c := a[len(a)-3:]
	d := a[:]

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)

	res := slices.Equal(a, d)

	fmt.Println(res)

	WindowClose()
}

func WindowClose() (string, error) {
	input, err := my_modules.InputString()

	if err != nil {
		log.Fatal(err)
	}

	return input, nil
}
