package my_modules

import (
	"fmt"
	"reflect"
	"strconv"
)

// Проверяет, является ли тип числовым
func isNumericKind(k reflect.Kind) bool {
	switch k {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64,
		reflect.Float32, reflect.Float64:
		return true
	}
	return false
}

// Универсальная проверка наличия ключа в map с поддержкой преобразований
func MapContainsKey(in interface{}, key any) (bool, error) {
	v := reflect.ValueOf(in)
	if v.Kind() != reflect.Map {
		return false, fmt.Errorf("input is not a map")
	}

	keyType := v.Type().Key()
	keyVal := reflect.ValueOf(key)

	// Прямое совпадение типов
	if keyVal.Type().ConvertibleTo(keyType) {
		convertedKey := keyVal.Convert(keyType)
		return v.MapIndex(convertedKey).IsValid(), nil
	}

	// Преобразование: строка → число
	if isNumericKind(keyType.Kind()) && keyVal.Kind() == reflect.String {
		strKey := keyVal.String()
		switch keyType.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			intVal, err := strconv.ParseInt(strKey, 10, 64)
			if err != nil {
				return false, fmt.Errorf("cannot convert string to int: %w", err)
			}
			convertedKey := reflect.ValueOf(intVal).Convert(keyType)
			return v.MapIndex(convertedKey).IsValid(), nil

		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			intVal, err := strconv.ParseUint(strKey, 10, 64)
			if err != nil {
				return false, fmt.Errorf("cannot convert string to uint: %w", err)
			}
			convertedKey := reflect.ValueOf(intVal).Convert(keyType)
			return v.MapIndex(convertedKey).IsValid(), nil

		case reflect.Float32, reflect.Float64:
			floatVal, err := strconv.ParseFloat(strKey, 64)
			if err != nil {
				return false, fmt.Errorf("cannot convert string to float: %w", err)
			}
			convertedKey := reflect.ValueOf(floatVal).Convert(keyType)
			return v.MapIndex(convertedKey).IsValid(), nil
		}
	}

	// Преобразование: число → строка
	if keyType.Kind() == reflect.String && isNumericKind(keyVal.Kind()) {
		var strKey string
		switch keyVal.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			strKey = strconv.FormatInt(keyVal.Int(), 10)
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			strKey = strconv.FormatUint(keyVal.Uint(), 10)
		case reflect.Float32, reflect.Float64:
			strKey = strconv.FormatFloat(keyVal.Float(), 'f', -1, 64)
		default:
			return false, fmt.Errorf("unsupported numeric kind: %v", keyVal.Kind())
		}
		convertedKey := reflect.ValueOf(strKey).Convert(keyType)
		return v.MapIndex(convertedKey).IsValid(), nil
	}

	return false, fmt.Errorf("key type %v not convertible to %v", keyVal.Type(), keyType)
}
