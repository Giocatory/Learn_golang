package main

import "fmt"

func main() {
	amount := 6
	changeNumber(&amount)
	fmt.Println(amount) // 12
}

func changeNumber(amount *int) {
	*amount *= 2
}
