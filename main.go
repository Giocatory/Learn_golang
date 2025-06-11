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
		Address: my_modules.Address{
			Street: "ул. Марджани",
			City:   "Казань",
			State:  "Россия",
		},
	}

	fmt.Print("1) ")
	user1.ShowSubscriberInfo()

	fmt.Print("2) ")
	user2 := my_modules.NewSubscriber("Egor", 5, false, "ул. Карима Тинчурина, Казань, Россия")
	user2.ShowSubscriberInfo()

	fmt.Print("3) ")
	employer := my_modules.NewEmployee("Yura", 100000, true, "ул. Карима Тинчурина, Казань, Россия")
	employer.ShowEmployeeInfo()

	fmt.Print("4) ")
	director := my_modules.NewEmployee("Mikhail", 10000000, true, "ул. Рублевская, Москва, Россия")
	fmt.Printf("Директор живет в городе %s :P\n", director.City)
}
