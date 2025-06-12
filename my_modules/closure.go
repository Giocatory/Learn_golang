package my_modules

import "math"

func Variables(a, b, c float64) func(float64) float64 {
	f := func(x float64) float64 {
		return a*math.Pow(x, 2) + b*x + c
	}
	return f
}
