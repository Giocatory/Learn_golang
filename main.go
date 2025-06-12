package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5}

	var arrCopy []int

	for _, v := range arr {
		arrCopy = append(arrCopy, v)
	}

	fmt.Println(arrCopy)
	fmt.Println(arr)

	arrCopy[0] = 100

	fmt.Println(arrCopy)
	fmt.Println(arr)
}
