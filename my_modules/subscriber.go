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

func (s Subscriber) ShowSubscriberInfo() {
	fmt.Println("Подписчик:")
	fmt.Printf("Имя: %10s\n", s.Name)
	fmt.Printf("Рейтинг: %10.2f\n", s.Rate)
	fmt.Printf("Активность: %10t\n", s.Active)
	fmt.Print("Адрес:\t")
	fmt.Printf("%s, %s, %s\n", s.Street, s.City, s.State)
	fmt.Println("========================================================================")
}

func (s Employee) ShowEmployeeInfo() {
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
