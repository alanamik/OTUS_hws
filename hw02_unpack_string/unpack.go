package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(str string) (string, error) {
	if str == "" || str == " " {
		return "", nil
	}
	runes := []rune(str)
	if runes[0] < 58 && runes[0] > 47 {
		return "", ErrInvalidString
	}
	var sb strings.Builder
	var replaceStr string

	for i, char := range runes {
		replaceStr = ""
		if char == 48 || (i+1 != len(runes) && runes[i+1] == 48) {
			replaceStr = "n"
		}
		if char < 58 && char > 48 {
			if i+1 != len(runes) && runes[i+1] < 58 && runes[i+1] > 47 {
				return "", ErrInvalidString
			}
			count, _ := strconv.Atoi(string(char))
			replaceStr = strings.Repeat(string(runes[i-1]), count-1)
			sb.WriteString(replaceStr)
		}
		if replaceStr == "" {
			sb.WriteRune(char)
		}
	}
	return sb.String(), nil
}
