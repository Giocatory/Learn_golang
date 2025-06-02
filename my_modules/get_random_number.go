package my_modules

import (
	"math/rand"
	"time"
)

func GetRandomNumber(number int) int {
	seconds := time.Now().Unix() // Получаем текущую дату и время в целом числе
	rand.Seed(seconds)           // Функция генератор случайного числа
	target := rand.Intn(number) + 1
	return target
}
