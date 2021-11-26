package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(s string) (string, error) {
	var result string
	for i, c := range s {
		n, err := strconv.Atoi(string(c))
		if i != 0 {
			if err != nil {
				if n != 0 {
					result += strings.Repeat(string(s[i-1]), n-1)
				} else {
					result = result[:i-1]
				}
			} else {
				result += string(s[i])
			}
		} else {
			if err != nil {
				result += string(s[i])
			} else {
				return ErrInvalidString.Error(), ErrInvalidString
			}
		}
	}
	return result, nil
}
