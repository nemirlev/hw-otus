package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
	"unicode/utf8"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(s string) (string, error) {
	var result strings.Builder
	var prevR rune
	var isPrevCharDigit, isPrevRuneEscaped bool

	for idx := 0; idx < len(s); {
		curR, size := utf8.DecodeRuneInString(s[idx:])    // получаем текущий символ и его размер
		nextR, _ := utf8.DecodeRuneInString(s[idx+size:]) // получаем следующий символ

		if curR == '\\' {
			if isPrevRuneEscaped {
				isPrevRuneEscaped = false
				result.WriteString(string(curR))
			} else {
				isPrevRuneEscaped = true
			}

			prevR = curR
			idx += size
			continue
		}

		if isPrevRuneEscaped {
			result.WriteString(string(curR))
			isPrevRuneEscaped = false
			prevR = curR
			idx += size
			continue
		}

		if unicode.IsLetter(curR) {
			if isPrevRuneEscaped {
				repeatCount, _ := strconv.Atoi(string(curR))
				result.WriteString(strings.Repeat(string(prevR), repeatCount-1))
			}

			if nextR == '0' {
				idx += size
				continue
			}

			result.WriteString(string(curR))
			isPrevCharDigit = false

			prevR = curR
		}

		if unicode.IsDigit(curR) {
			if curR == '0' && !isPrevCharDigit {
				prevR = curR
				idx += size
				continue
			}

			if isPrevRuneEscaped {
				isPrevRuneEscaped = false
				prevR = curR
				result.WriteString(string(curR))
				idx += size
				continue
			}

			if prevR == 0 || isPrevCharDigit {
				return "", ErrInvalidString
			}

			isPrevCharDigit = true

			repeatCount, _ := strconv.Atoi(string(curR))
			result.WriteString(strings.Repeat(string(prevR), repeatCount-1))
		}

		idx += size
	}

	return result.String(), nil
}
