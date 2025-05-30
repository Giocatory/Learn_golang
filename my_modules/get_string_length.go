package my_modules

func GetStringLength(text string) int {
	if text == "" {
		return 0
	}
	temp := []rune(text)
	return len(temp)
}
