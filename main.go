package main

import (
	"fmt"
)

func main() {
	someArr := []string{"one", "second", "long string"}

loop:
	for _, v := range someArr {
		temp := len(v)
		switch {
		case temp > 0 && temp <= 5:
			fmt.Println("short word")
		case temp > 5 && temp <= 7:
			fmt.Println("middle word")
			break loop
		default:
			fmt.Println("default value")
		}
	}
}
