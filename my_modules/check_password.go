package my_modules

const Template string = "ESC24er32!!!"

func CheckPassword(password string) bool {
	passRunes := []rune(password)

	if len(passRunes) < 6 {
		return false
	}

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
		case char == '*' || char == '/' || char == '-' || char == '+' || char == '!':
			hasSpecial = true
		default:
			// Недопустимый символ
			return false
		}
	}

	return hasLower && hasUpper && hasDigit && hasSpecial
}
