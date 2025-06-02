package my_modules

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func ReadCSV(filename string) (map[string]map[string]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	result := make(map[string]map[string]string)
	scanner := bufio.NewScanner(file)
	lineNum := 0

	for scanner.Scan() {
		line := scanner.Text()
		lineNum++

		// Пропускаем заголовок
		if lineNum == 1 {
			continue
		}

		// Разбиваем строку на поля
		fields := strings.Split(line, ",")
		if len(fields) < 3 {
			return nil, fmt.Errorf("ошибка в строке %d: недостаточно полей", lineNum)
		}

		// Очищаем поля от пробелов
		id := strings.TrimSpace(fields[0])
		name := strings.TrimSpace(fields[1])
		age := strings.TrimSpace(fields[2])

		// Создаем вложенную карту для текущего ID
		result[id] = map[string]string{
			"name": name,
			"age":  age,
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return result, nil
}