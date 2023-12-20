package hw02unpackstring

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// DONE

var ErrInvalidString = errors.New("invalid string")

func Unpack(str string) (string, error) {
	if str == "" || str == " " {
		return "", nil
	}

	if rune(str[0]) < 58 && rune(str[0]) > 47 {
		return "", ErrInvalidString
	}
	var sb strings.Builder
	var replaceStr string

	for i, char := range str {
		replaceStr = ""
		if char == 48 {
			s := sb.String()
			s = s[:len(s)-1]
			sb.Reset()
			sb.WriteString(s)

			fmt.Println(sb.String())
		}
		if str[i] < 58 && str[i] > 48 {
			if i+1 != len(str) && str[i+1] < 58 && str[i+1] > 47 {
				return "", ErrInvalidString
			}
			count, err := strconv.Atoi(string(char))
			if err != nil {
				return "", errors.New("can't converted char to int")
			}
			replaceStr = strings.Repeat(string(str[i-1]), count-1)
		}

		if replaceStr != "" {
			sb.WriteString(replaceStr)
		} else if char != 48 {
			sb.WriteRune(char)
		}
	}

	return sb.String(), nil
}
