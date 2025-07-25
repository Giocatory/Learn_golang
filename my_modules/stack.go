package my_modules

import "fmt"

type Stack[T any] struct {
	vals []T
}

func (s *Stack[T]) Push(val T) {
	s.vals = append(s.vals, val)
}

func (s *Stack[T]) Pop() (T, bool) {
	if len(s.vals) == 0 {
		var zero T
		return zero, false
	}
	top := s.vals[len(s.vals)-1]
	s.vals = s.vals[:len(s.vals)-1]
	return top, true
}

func (s *Stack[T]) Print() {
	fmt.Println(s.vals)
}

// Перепроверка
// package main

// import (
// 	"Learn_golang/my_modules"
// 	"fmt"
// )

// func main() {
// 	var intStack my_modules.Stack[int]

// 	intStack.Push(10)
// 	intStack.Push(20)
// 	intStack.Push(30)
// 	v, ok := intStack.Pop()

// 	fmt.Println(v, ok)
// 	intStack.Print()
// }
