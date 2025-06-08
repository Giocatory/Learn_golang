package main

import (
	"fmt"
)

type person struct {
	name string
	age  int
	pet  string
}

func main() {
	user1 := person{
		name: "Mikhail",
		age:  37,
		pet:  "Cat",
	}

	fmt.Println(user1.name)
	fmt.Println(user1.age)
	fmt.Println(user1.pet)
}
