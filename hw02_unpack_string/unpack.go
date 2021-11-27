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
                // Если символ буква
                case unicode.IsLetter(c):
                        result += string(c)

                // Если не буква и не цифра
                case !unicode.IsLetter(c) && !unicode.IsDigit(c):
                        return ErrInvalidString.Error(), ErrInvalidString
                // Первый символ цифра
                case i == 0 && unicode.IsDigit(c):
                        return ErrInvalidString.Error(), ErrInvalidString
                // Цифра часть числа
                case unicode.IsDigit(rune(s[i+1])):
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
