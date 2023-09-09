package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
	"unicode/utf8"
)

// ErrInvalidString is returned when the input string is invalid.
var ErrInvalidString = errors.New("invalid string")

// Unpack returns the unpacked string or an error if the input string is invalid.
func Unpack(s string) (string, error) {
	// создаем билдер для формирования результирующей строки
	var result strings.Builder
	var prev rune
	var prevIsDigit bool

	// проходим по всем символам входной строки
	for idx := 0; idx < len(s); {
		curR, size := utf8.DecodeRuneInString(s[idx:])    // получаем текущий символ и его размер
		nextR, _ := utf8.DecodeRuneInString(s[idx+size:]) // получаем следующий символ и его размер

		// если текущий символ - буква
		if unicode.IsLetter(curR) {
			// если следующий символ - 0, то ничего не делаем
			if nextR == '0' {
				idx += size
				prev = curR
				prevIsDigit = unicode.IsDigit(curR)
				continue
			}

			result.WriteString(string(curR))
			prevIsDigit = false
			prev = curR
		}

		// если текущий символ - цифра
		if unicode.IsDigit(curR) {
			// если текущий символ - 0 и предыдущий символ - не цифра, то ничего не делаем
			if curR == '0' && !prevIsDigit {
				prev = curR
				prevIsDigit = unicode.IsDigit(curR)
				idx += size
				continue
			}

			// если это первый символ в строке или предыдущий символ - цифра
			if prev == 0 || prevIsDigit {
				// возвращаем ошибку
				return "", ErrInvalidString
			}

			// запоминаем, что предыдущий символ - цифра
			prevIsDigit = true

			// добавляем в результирующую строку повторяющиеся символы
			repeatCount, _ := strconv.Atoi(string(curR))
			result.WriteString(strings.Repeat(string(prev), repeatCount-1))
		}

		idx += size // переходим к следующему символу
	}

	// возвращаем результирующую строку
	return result.String(), nil
}
