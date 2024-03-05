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
		if char == 48 {
			s := sb.String()
			if runes[i-1] > 127 { // если предыдущая руна имеет другую кодировку
				bef, _, _ := strings.Cut(s, string(runes[i-1]))
				sb.Reset()
				sb.WriteString(bef)
			} else {
				s = s[:len(s)-1]
				sb.Reset()
				sb.WriteString(s)
			}
		}
		if char < 58 && char > 48 {
			if i+1 != len(runes) && runes[i+1] < 58 && runes[i+1] > 47 {
				return "", ErrInvalidString
			}
			count, _ := strconv.Atoi(string(char))
			replaceStr = strings.Repeat(string(runes[i-1]), count-1)
		}
		if replaceStr != "" {
			sb.WriteString(replaceStr)
		} else if char != 48 {
			sb.WriteRune(char)
		}
	}

	return sb.String(), nil
}
