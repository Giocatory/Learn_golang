package main

import "fmt"

func main() {
	name := "Михаил"
	var firstLatter rune = []rune(name)[0]
	fmt.Println(string(firstLatter))
}
