package main

import (
	"Learn_golang/my_modules"
	"fmt"
)

func main() {
	pass := "IPsequence*351"

	checked := my_modules.CheckPassword(pass)

	fmt.Println(checked)
}
