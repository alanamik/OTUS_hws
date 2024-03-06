package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(str string) (string, error) {
	if len(str) <= 1 {
		return "", nil
	}
	runes := []rune(str)
	if unicode.IsDigit(runes[0]) {
		return "", ErrInvalidString
	}
	var sb strings.Builder
	var replaceStr string

	for i, char := range runes {
		if unicode.IsDigit(char) {
			if i+1 != len(runes) && unicode.IsDigit(runes[i+1]) {
				return "", ErrInvalidString
			}
			count, _ := strconv.Atoi(string(char))
			if count != 0 {
				replaceStr = strings.Repeat(string(runes[i-1]), count-1)
				sb.WriteString(replaceStr)
			} else {
				s := sb.String()
				bef, _ := strings.CutSuffix(s, string(runes[i-1]))
				sb.Reset()
				sb.WriteString(bef)
			}
		} else {
			sb.WriteRune(char)
		}
	}
	return sb.String(), nil
}
