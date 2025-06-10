package main

import (
	"fmt"
)

type subscriber struct {
	name   string
	rate   float64
	active bool
}

func showInfo(s subscriber) {
	fmt.Println("Подписчик:")
	fmt.Printf("Имя: %10s\n", s.name)
	fmt.Printf("Рейтинг: %10.2f\n", s.rate)
	fmt.Printf("Активность: %10t\n", s.active)
}

func newSubscriber(name string, rate float64, active bool) subscriber {
	s := subscriber{
		name:   name,
		rate:   rate,
		active: active,
	}
	return s
}

func main() {
	user1 := subscriber{
		name:   "Alyona",
		rate:   4.99,
		active: true,
	}

	fmt.Print("1) ")
	showInfo(user1)

	user2 := newSubscriber("Boris", 5, false)
	fmt.Print("2) ")
	showInfo(user2)
}
