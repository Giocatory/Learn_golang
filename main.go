package main

import (
	md "Learn_golang/my_modules"
	"fmt"
)

func main() {
	var intStack md.Stack[int]

	intStack.Push(10)
	intStack.Push(20)
	intStack.Push(30)
	v, ok := intStack.Pop()

	fmt.Println(v, ok)
	intStack.Print()
}
