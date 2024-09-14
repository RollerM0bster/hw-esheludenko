package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(str string) (string, error) {
	_, err := strconv.Atoi(str)
	runes := []rune(str)
	if err != nil {
		builder := strings.Builder{}
		for i := 0; i < len(runes)-1; i++ {
			processErr := processSymbol(i, runes[i], runes[i+1], &builder, len(runes)-1)
			if processErr != nil {
				return "", processErr
			}
		}
		return builder.String(), nil
	}
	return "", ErrInvalidString
}

func processSymbol(index int, r1 rune, next rune, builder *strings.Builder, penultLen int) error {
	if index == 0 {
		if isDigit(r1) {
			return ErrInvalidString
		}
	}
	if unicode.IsLetter(r1) {
		if unicode.IsDigit(next) {
			digit, _ := strconv.Atoi(string(next))
			builder.WriteString(strings.Repeat(string(r1), digit))
		} else {
			builder.WriteString(string(r1))
		}
	} else {
		if isDigit(next) {
			return ErrInvalidString
		}
	}
	if index == penultLen-1 && unicode.IsLetter(next) {
		builder.WriteString(string(next))
	}
	return nil
}

func isDigit(r rune) bool {
	return r >= '0' && r <= '9'
}
