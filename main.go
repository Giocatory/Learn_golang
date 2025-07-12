package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	word := "абв"
	letter, n := utf8.DecodeRuneInString(word)
	fmt.Printf("Первый символ: %c\tБайт: %v\n", letter, n)
	fmt.Printf("Длина строки: %d\tБайт: %v\n", utf8.RuneCountInString(word), len(word))
}
