package main

import (
	"fmt"
)

type employer struct {
	firstname string
	lastname  string
	id        int
}

func main() {
	firstEmployer := employer{}
	secondEmployer := employer{
		firstname: "Tatyana",
		lastname:  "Derkunova",
		id:        1,
	}

	firstEmployer.firstname = "Mikhail"
	firstEmployer.lastname = "Derkunov"
	firstEmployer.id = 0

	fmt.Printf("First:\nid:\t%v;\nName:\t%v %v;\n", firstEmployer.id, firstEmployer.firstname, firstEmployer.lastname)
	fmt.Printf("Second:\nid:\t%v;\nName:\t%v %v;\n", secondEmployer.id, secondEmployer.firstname, secondEmployer.lastname)
}
