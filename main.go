package main

import (
	"Learn_golang/my_modules"
	"fmt"
)

func main() {
	// Map с int-ключами
	m1 := map[int]string{1: "one", 2: "two"}
	fmt.Println(my_modules.MapContainsKey(m1, "1")) // true, nil (строка → int)

	// Map со string-ключами
	m2 := map[string]int{"1": 1, "2": 2}
	fmt.Println(my_modules.MapContainsKey(m2, 1)) // true, nil (int → строка)

	// Map с float-ключами
	m3 := map[float64]bool{1.5: true}
	fmt.Println(my_modules.MapContainsKey(m3, "1.5")) // true, nil (строка → float)

	// Неподдерживаемый тип
	m4 := map[bool]int{true: 1}
	fmt.Println(my_modules.MapContainsKey(m4, "true")) // false, error
}
