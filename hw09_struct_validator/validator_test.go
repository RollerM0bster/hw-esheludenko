package hw09structvalidator

import (
	"errors"
	"fmt"
	"reflect"
	"testing"
)

type (
	User struct {
		ID     string `json:"id" validate:"len:36"`
		Name   string
		Age    int      `validate:"min:18|max:50"`
		Email  string   `validate:"regexp:^[\\w\\.-]+@[\\w\\.-]+\\.[a-z]{2,}$"`
		Phones []string `validate:"len:11"`
	}

	App struct {
		Version string `validate:"len:5"`
	}

	Response struct {
		Code int    `validate:"in:200,404,500"`
		Body string `json:"omitempty"`
	}

	StructWithBadTag struct {
		BadField string `validate:"len::5"`
	}
	StructWithUnsupportedField struct {
		BadField float64 `validate:"min:1|max:2"`
	}
)

func TestBadStructValidate(t *testing.T) {
	bs := StructWithBadTag{BadField: "1"}
	bs2 := StructWithUnsupportedField{BadField: 1.1}
	result := Validate(bs)
	if reflect.TypeOf(result) == reflect.TypeOf(&ValidationError{}) {
		t.Errorf("Expected error, got ValidationError struct")
	}
	result2 := Validate(bs2)
	if reflect.TypeOf(result2) == reflect.TypeOf(&ValidationError{}) {
		t.Errorf("Expected error, got ValidationError struct")
	}
}

func TestValidate(t *testing.T) {
	tests := []struct {
		name        string
		input       interface{}
		expectedErr error
	}{
		{
			name: "Valid User",
			input: User{
				ID:     "12345678-1234-5678-1234-567812345678",
				Name:   "John Doe",
				Age:    25,
				Email:  "john.doe@example.com",
				Phones: []string{"88005553535", "12345678901"},
			},
			expectedErr: nil,
		},
		{
			name: "Invalid User ID length",
			input: User{
				ID:     "12345",
				Name:   "John Doe",
				Age:    25,
				Email:  "john.doe@example.com",
				Phones: []string{"88005553535", "12345678901"},
			},
			expectedErr: ValidationErrors{
				{Field: "ID", Err: fmt.Errorf("field ID must be 36 characters, but got 5")},
			},
		},
		{
			name: "Invalid User Age below min",
			input: User{
				ID:     "12345678-1234-5678-1234-567812345678",
				Name:   "John Doe",
				Age:    17,
				Email:  "john.doe@example.com",
				Phones: []string{"88005553535", "12345678901"},
			},
			expectedErr: ValidationErrors{
				{Field: "Age", Err: fmt.Errorf("field Age must be at least 18, but got 17")},
			},
		},
		{
			name: "Invalid User Email format",
			input: User{
				ID:     "12345678-1234-5678-1234-567812345678",
				Name:   "John Doe",
				Age:    25,
				Email:  "john.doe",
				Phones: []string{"88005553535", "12345678901"},
			},
			expectedErr: ValidationErrors{
				{Field: "Email", Err: fmt.Errorf("field Email must match regexp: ^[\\w\\.-]+@[\\w\\.-]+\\.[a-z]{2,}$")},
			},
		},
		{
			name: "Invalid Phone numbers length",
			input: User{
				ID:     "12345678-1234-5678-1234-567812345678",
				Name:   "John Doe",
				Age:    25,
				Email:  "john.doe@example.com",
				Phones: []string{"12345", "12345678901"},
			},
			expectedErr: ValidationErrors{
				{Field: "Phones", Err: fmt.Errorf("field Phones must be 11 characters, but got 5")},
			},
		},
		{
			name: "Valid App",
			input: App{
				Version: "1.0.0",
			},
			expectedErr: nil,
		},
		{
			name: "Invalid App Version length",
			input: App{
				Version: "1.0",
			},
			expectedErr: ValidationErrors{
				{Field: "Version", Err: fmt.Errorf("field Version must be 5 characters, but got 3")},
			},
		},
		{
			name: "Valid Response",
			input: Response{
				Code: 200,
				Body: "OK",
			},
			expectedErr: nil,
		},
		{
			name: "Invalid Response Code",
			input: Response{
				Code: 403,
				Body: "Forbidden",
			},
			expectedErr: ValidationErrors{
				{Field: "Code", Err: fmt.Errorf("field Code must contain at least one (200,404,500)")},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := Validate(tt.input)

			if tt.expectedErr == nil {
				if err != nil {
					checkNoError(t, err)
				}
			} else {
				checkError(t, err, tt.expectedErr)
			}
		})
	}
}

func checkNoError(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Errorf("Expected no error, got: %v", err)
	}
}

func checkError(t *testing.T, err error, expectedErr error) {
	t.Helper()
	var validationErrors ValidationErrors
	if !errors.As(err, &validationErrors) {
		t.Errorf("Expected ValidationErrors, got: %v", err)
		return
	}

	var expectedErrors ValidationErrors
	errors.As(expectedErr, &expectedErrors)
	for _, expected := range expectedErrors {
		found := false
		for _, actual := range validationErrors {
			if expected.Field == actual.Field && expected.Err.Error() == actual.Err.Error() {
				found = true
				break
			}
		}
		if !found {
			t.Errorf("Expected error for field %s: %s, but not found", expected.Field, expected.Err)
		}
	}
}
