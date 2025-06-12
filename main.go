package main

import "Learn_golang/my_modules"

func main() {
	const number = 5
	someTarget := my_modules.MakePointer(number)

	*someTarget = 10

	println(*someTarget)
	println(number)
}
