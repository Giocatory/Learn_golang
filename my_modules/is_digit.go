package my_modules

import "unicode"

func IsDigit(str string) bool {
	isDigitOnly := true
	for _, char := range str {
		if !unicode.IsDigit(char) {
			isDigitOnly = false
			break
		}
	}

	return isDigitOnly
}
