package my_modules

import "fmt"

func CheckExam(grade int) (string, error) {
	if grade >= 0 && grade < 31 {
		return "Не Удовлетворительно!", nil
	}
	if grade >= 31 && grade < 61 {
		return "Удовлетворительно!", nil
	}
	if grade >= 61 && grade < 81 {
		return "Хорошо!", nil
	}
	if grade >= 81 && grade <= 100 {
		return "Отлично!", nil
	} else {
		return "", fmt.Errorf("оценка должна быть в пределах от 0 до 100, вы ввели: %d", grade)
	}
}
