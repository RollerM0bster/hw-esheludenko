package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(str string) (string, error) {
	_, err := strconv.Atoi(str)
	runes := []rune(str)
	if err != nil {
		builder := strings.Builder{}
		if len(runes) > 0 && isDigit(runes[0]) {
			return "", ErrInvalidString
		}
		for i := 0; i <= len(runes)-2; i++ {
			processErr := processSymbol(runes[i], runes[i+1], &builder)
			if processErr != nil {
				return "", processErr
			}
		}
		if len(runes) >= 2 && !isDigit(runes[len(runes)-1]) {
			builder.WriteString(string(runes[len(runes)-1]))
		}
		return builder.String(), nil
	}
	return "", ErrInvalidString
}

func processSymbol(r1 rune, next rune, builder *strings.Builder) error {
	if !isDigit(r1) {
		if isDigit(next) {
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
	return nil
}

func isDigit(r rune) bool {
	return r >= '0' && r <= '9'
}
