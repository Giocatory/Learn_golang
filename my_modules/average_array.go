package my_modules

func GetAverageIntNumbers(numbers ...int) int {
	total := 0

	for _, number := range numbers {
		total += number
	}

	return total / len(numbers)
}

func GetAverageFloatNumbers(numbers ...float64) float64 {
	total := 0.0

	for _, number := range numbers {
		total += number
	}

	return total / float64(len(numbers))
}
