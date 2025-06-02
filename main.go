package main

import "fmt"

func main() {
	a := 5

	b := &a

	*b = 13

	fmt.Println(a)
}
