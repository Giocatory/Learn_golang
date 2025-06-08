package main

import (
	"fmt"
)

type employer struct {
	firstName string
	lastName  string
	id        int
}

func main() {
	firstEmployer := employer{}
	secondEmployer := employer{
		firstName: "Tatyana",
		lastName:  "Derkunova",
		id:        1,
	}

	firstEmployer.firstName = "Mikhail"
	firstEmployer.lastName = "Derkunov"
	firstEmployer.id = 0

	fmt.Printf("First:\nid:\t%v;\nName:\t%v %v;\n", firstEmployer.id, firstEmployer.firstName, firstEmployer.lastName)
	fmt.Printf("Second:\nid:\t%v;\nName:\t%v %v;\n", secondEmployer.id, secondEmployer.firstName, secondEmployer.lastName)
}
