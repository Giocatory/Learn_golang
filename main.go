package main

import (
	"Learn_golang/my_modules"
	"fmt"
)

func main() {
	user1 := my_modules.Subscriber{
		Name:   "Marsel",
		Rate:   4.99,
		Active: true,
		HomeAddress: my_modules.Address{
			Street: "ул. Марджани",
			City:   "Казань",
			State:  "Россия",
		},
	}

	fmt.Print("1) ")
	my_modules.ShowSubscriberInfo(user1)

	fmt.Print("2) ")
	user2 := my_modules.NewSubscriber("Egor", 5, false, "ул. Карима Тинчурина, Казань, Россия")
	my_modules.ShowSubscriberInfo(user2)

	fmt.Print("3) ")
	employer := my_modules.NewEmployee("Yura", 100000, true, "ул. Карима Тинчурина, Казань, Россия")
	my_modules.ShowEmployeeInfo(employer)
}
