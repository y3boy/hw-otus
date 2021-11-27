package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(s string) (string, error) {
	var result string
	for i, c := range s {
		n, _ := strconv.Atoi(string(c))
		switch {
		case !(unicode.IsDigit(c)): // Если символ буква
			result += string(c)
		case i == 0 && unicode.IsDigit(c): // Первый символ цифра
			return ErrInvalidString.Error(), ErrInvalidString
		case unicode.IsDigit(rune(s[i+1])): // Цифра часть числа
			return ErrInvalidString.Error(), ErrInvalidString
		default: // Если цифра
			if n == 0 {
				result = result[:len(result)-1]
			} else {
				result += strings.Repeat(string(s[i-1]), n-1)
			}
		}
	}
	return result, nil
}
