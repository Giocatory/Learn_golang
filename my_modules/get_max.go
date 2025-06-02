package my_modules

import "reflect"

func GetMax(literals ...any) (any, error) {
	if reflect.TypeOf(literals[0]).Kind() == reflect.Int {
		max := 0
		for _, literal := range literals {
			if literal.(int) > max {
				max = literal.(int)
			}
		}
		return max, nil
	} else if reflect.TypeOf(literals[0]).Kind() == reflect.Float64 {
		max := 0.0
		for _, literal := range literals {
			if literal.(float64) > max {
				max = literal.(float64)
			}
		}
		return max, nil
	} else if reflect.TypeOf(literals[0]).Kind() == reflect.String {
		max := ""
		for _, literal := range literals {
			if literal.(string) > max {
				max = literal.(string)
			}
		}
		return max, nil
	} else {
		return nil, nil
	}
}
