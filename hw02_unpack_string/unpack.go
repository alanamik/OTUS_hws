package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(str string) (string, error) {
	if len(str) <= 1 {
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
		if i+1 != len(runes) && runes[i+1] == 48 {
			replaceStr = "no"
		}
		if i+1 != len(runes) && runes[i+1] < 58 && runes[i+1] > 48 {
			if i+2 != len(runes) && runes[i+2] < 58 && runes[i+2] > 47 {
				return "", ErrInvalidString
			}
			count, _ := strconv.Atoi(string(runes[i+1]))
			replaceStr = strings.Repeat(string(char), count)
			sb.WriteString(replaceStr)
		}
		if replaceStr == "" && (char < 48 || char >= 58) {
			_, _ = sb.WriteRune(char)
		}
	}
	return sb.String(), nil
}
