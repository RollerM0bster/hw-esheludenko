package hw02unpackstring

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestUnpack(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{input: "a4bc2d5e", expected: "aaaabccddddde"},
		{input: "abccd", expected: "abccd"},
		{input: "", expected: ""},
		{input: "aaa0b", expected: "aab"},
		{input: "d\n5abc", expected: "d\n\n\n\n\nabc"},
		// uncomment if task with asterisk completed
		// {input: `qwe\4\5`, expected: `qwe45`},
		// {input: `qwe\45`, expected: `qwe44444`},
		// {input: `qwe\\5`, expected: `qwe\\\\\`},
		// {input: `qwe\\\3`, expected: `qwe\3`},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.input, func(t *testing.T) {
			result, err := Unpack(tc.input)
			require.NoError(t, err)
			require.Equal(t, tc.expected, result)
		})
	}
}

func TestUnpackInvalidString(t *testing.T) {
	invalidStrings := []string{"3abc", "45", "aaa10b"}
	for _, tc := range invalidStrings {
		tc := tc
		t.Run(tc, func(t *testing.T) {
			_, err := Unpack(tc)
			require.Truef(t, errors.Is(err, ErrInvalidString), "actual error %q", err)
		})
	}
}

func TestUnpackDifferentStrings(t *testing.T) {
	differentStrings := []struct {
		input    string
		expected string
	}{
		{input: "a4bc2d5e", expected: "aaaabccddddde"},
		{input: "3abc", expected: "invalid string"},
		{input: "aaa10b", expected: "invalid string"},
		{input: "abccd", expected: "abccd"},
		{input: "d\n5abc", expected: "d\n\n\n\n\nabc"},
	}
	for _, tc := range differentStrings {
		tc := tc
		t.Run(tc.input, func(t *testing.T) {
			str, err := Unpack(tc.input)
			if err != nil {
				fmt.Println(tc.input, err)
				require.Truef(t, errors.Is(err, ErrInvalidString), "actual error %q", err)
			} else {
				fmt.Println(tc.input, str)
				require.Equal(t, tc.expected, str)
				fmt.Println("Result: " + str)
			}
		})
	}
}
