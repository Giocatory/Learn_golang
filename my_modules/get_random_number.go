package my_modules

import (
	"errors"
	"math/rand"
	"time"
)

func GetRandomNumber(number int) (int, error) {
	if number < 0 {
		return 0, errors.New("number is negative")
	}
	seconds := time.Now().Unix() // Получаем текущую дату и время в целом числе
	rand.Seed(seconds)           // Функция генератор случайного числа
	target := rand.Intn(number) + 1
	return target, nil
}
