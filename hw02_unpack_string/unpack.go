package hw02unpackstring

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func Unpack(s string) (string, error) {
	var builder strings.Builder
	var previousRune rune

	for _, r := range s {
		if unicode.IsDigit(r) {
			// If the current rune is a digit, we expect the previous rune
			// to be repeated the number of times indicated by the digit.
			if previousRune == 0 {
				// If the previous rune is not set, it means the string
				// is not correctly formatted, so we return an error.
				return "", fmt.Errorf("incorrect string: %s", s)
			}

			// We convert the digit to a number and repeat the previous
			// rune that number of times.
			num, err := strconv.Atoi(string(r))
			if err != nil {
				return "", err
			}
			builder.WriteString(strings.Repeat(string(previousRune), num))
		} else {
			// If the current rune is not a digit, we simply append it
			// to the string being built.
			builder.WriteRune(r)
			previousRune = r
		}
	}

	return builder.String(), nil
}
