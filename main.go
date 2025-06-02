package main

import (
	"Learn_golang/my_modules"
	"fmt"
)

func main() {
	//pass := "IPadventures*351"

	checked := my_modules.CheckPassword(my_modules.Template)

	fmt.Println(checked)
}
