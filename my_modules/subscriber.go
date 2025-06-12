package my_modules

import (
	"fmt"
	"strings"
)

type Subscriber struct {
	Name   string
	Rate   float64
	Active bool
	Address
}

type Employee struct {
	Name   string
	Salary float64
	Active bool
	Address
}

type Address struct {
	Street string
	City   string
	State  string
}

func (s *Subscriber) ToString() {
	fmt.Println("Подписчик:")
	fmt.Printf("Имя: %10s\n", s.Name)
	fmt.Printf("Рейтинг: %10.2f\n", s.Rate)
	fmt.Printf("Активность: %10t\n", s.Active)
	fmt.Print("Адрес:\t")
	fmt.Printf("%s, %s, %s\n", s.Street, s.City, s.State)
	fmt.Println("========================================================================")
}

func (s *Employee) ToString() {
	fmt.Println("Сотрудник:")
	fmt.Printf("Имя: %10s\n", s.Name)
	fmt.Printf("Зарплата: %10.2f\n", s.Salary)
	fmt.Printf("Активность: %10t\n", s.Active)
	fmt.Print("Адрес:\t")
	fmt.Printf("%s, %s, %s\n", s.Street, s.City, s.State)
	fmt.Println("========================================================================")
}

func NewSubscriber(name string, rate float64, active bool, address string) Subscriber {
	ad := strings.Split(address, ",")
	s := Subscriber{
		Name:   name,
		Rate:   rate,
		Active: active,
		Address: Address{
			Street: ad[0],
			City:   ad[1],
			State:  ad[2],
		},
	}
	return s
}

func NewEmployee(name string, salary float64, active bool, address string) Employee {
	ad := strings.Split(address, ",")
	s := Employee{
		Name:   name,
		Salary: salary,
		Active: active,
		Address: Address{
			Street: ad[0],
			City:   ad[1],
			State:  ad[2],
		},
	}
	return s
}

/*
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
*/
