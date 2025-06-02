package my_modules

func CheckPassword(password string) bool {
	// Преобразуем пароль в руны для корректной обработки
	passRunes := []rune(password)

	if len(passRunes) < 6 {
		return false
	}

	// Флаги для проверки наличия обязательных категорий
	hasLower := false   // Строчные буквы
	hasUpper := false   // Заглавные буквы
	hasDigit := false   // Цифры
	hasSpecial := false // Спецсимволы (*, /, -, +)

	for _, char := range passRunes {
		switch {
		case 'a' <= char && char <= 'z':
			hasLower = true
		case 'A' <= char && char <= 'Z':
			hasUpper = true
		case '0' <= char && char <= '9':
			hasDigit = true
		case char == '*' || char == '/' || char == '-' || char == '+':
			hasSpecial = true
		default:
			// Недопустимый символ
			return false
		}
	}

	// Проверяем наличие всех требуемых категорий
	return hasLower && hasUpper && hasDigit && hasSpecial
}
