package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

// ErrInvalidString is returned when the input string is invalid.
var ErrInvalidString = errors.New("invalid string")

// Unpack returns the unpacked string or an error if the input string is invalid.
func Unpack(s string) (string, error) {
	// создаем билдер для формирования результирующей строки
	var result strings.Builder
	// итерируем по строке
	for i := 0; i < len(s); i++ {
		// если текущий символ является цифрой
		if unicode.IsDigit(rune(s[i])) {
			// и если это первый символ в строке или предыдущий символ тоже цифра
			if i == 0 || unicode.IsDigit(rune(s[i-1])) {
				// возвращаем ошибку
				return "", ErrInvalidString
			}
		}
		// если следующий символ является цифрой
		if i+1 < len(s) && unicode.IsDigit(rune(s[i+1])) {
			// парсим цифру в число
			count, err := strconv.Atoi(string(s[i+1]))
			if err != nil {
				return "", err
			}
			// повторяем текущий символ count раз
			for j := 0; j < count; j++ {
				result.WriteByte(s[i])
			}
			// пропускаем цифру
			i++
		} else {
			// если следующий символ не является цифрой, добавляем текущий символ в результирующую строку
			result.WriteByte(s[i])
		}
	}
	// возвращаем результирующую строку
	return result.String(), nil
}
