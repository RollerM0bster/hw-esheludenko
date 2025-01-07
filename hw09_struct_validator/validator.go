package hw09structvalidator

import (
	"errors"
	"fmt"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

type ValidationError struct {
	Field string
	Err   error
}

var (
	ErrInvalidTag    = errors.New("invalid tag")
	ErrInvalidRegExp = errors.New("invalid regular expression")
)

type ValidationErrors []ValidationError

func (v ValidationErrors) Error() string {
	var res strings.Builder
	for _, err := range v {
		res.WriteString(fmt.Sprintf("Field %s: %s\n", err.Field, err.Err))
	}
	return res.String()
}

func Validate(v interface{}) error {
	var validationErrors ValidationErrors
	value := reflect.ValueOf(v)
	if value.Kind() == reflect.Ptr {
		value = value.Elem()
	}
	if value.Kind() != reflect.Struct {
		return errors.New("value must be a struct")
	}
	tp := value.Type()
	for i := 0; i < value.NumField(); i++ {
		field := tp.Field(i)
		fieldValue := value.Field(i)
		tag := field.Tag.Get("validate")
		if tag == "" {
			continue
		}
		rules := strings.Split(tag, "|")
		for _, rule := range rules {
			if err := validateField(field.Name, rule, fieldValue); err != nil {
				validationErrors = append(validationErrors, ValidationError{Field: field.Name, Err: err})
			}
		}
	}
	if len(validationErrors) > 0 {
		return validationErrors
	}
	return nil
}

func validateField(fieldName, rule string, value reflect.Value) error {
	switch value.Kind() {
	case reflect.String:
		return validateString(fieldName, value.String(), rule)
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return validateInt(fieldName, rule, int(value.Int()))
	case reflect.Slice:
		for i := 0; i < value.Len(); i++ {
			elem := value.Index(i)
			if err := validateField(fieldName, rule, elem); err != nil {
				return err
			}
		}
	case reflect.Invalid, reflect.Bool, reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64,
		reflect.Uintptr, reflect.Float32, reflect.Float64, reflect.Complex64, reflect.Complex128,
		reflect.Array, reflect.Chan, reflect.Func, reflect.Interface, reflect.Map,
		reflect.Pointer, reflect.Struct, reflect.UnsafePointer:
		return fmt.Errorf("unsupported type %s for field %s", value.Kind(), fieldName)
	default:
		return fmt.Errorf("unsupported type %s for field %s", value.Kind(), fieldName)
	}
	return nil
}

func validateString(field, str, rule string) error {
	ruleParts := strings.Split(rule, ":")
	if len(ruleParts) != 2 {
		return fmt.Errorf("%w: %s", ErrInvalidTag, rule)
	}
	switch ruleParts[0] {
	case "len":
		length, err := strconv.Atoi(ruleParts[1])
		if err != nil {
			return fmt.Errorf("field %s, %w: %w", field, ErrInvalidTag, err)
		}
		if len(str) != length {
			return fmt.Errorf("field %s must be %d characters, but got %d", field, length, len(str))
		}
	case "regexp":
		reg, err := regexp.Compile(ruleParts[1])
		if err != nil {
			return fmt.Errorf("field %s, %w: %w", field, ErrInvalidRegExp, err)
		}
		if !reg.MatchString(str) {
			return fmt.Errorf("field %s must match regexp: %s", field, ruleParts[1])
		}
	case "in":
		vars := strings.Split(ruleParts[1], ",")
		for _, v := range vars {
			if str == v {
				return nil
			}
		}
		return fmt.Errorf("field %s must contain at least one (%s)", field, ruleParts[1])
	}
	return nil
}

func validateInt(fieldName, rule string, num int) error {
	ruleParts := strings.Split(rule, ":")
	if len(ruleParts) != 2 {
		return fmt.Errorf("%w: %s", ErrInvalidTag, rule)
	}
	switch ruleParts[0] {
	case "min":
		min, err := strconv.Atoi(ruleParts[1])
		if err != nil {
			return fmt.Errorf("field %s %w: %w", fieldName, ErrInvalidTag, err)
		}
		if num < min {
			return fmt.Errorf("field %s must be at least %d, but got %d", fieldName, min, num)
		}
	case "max":
		max, err := strconv.Atoi(ruleParts[1])
		if err != nil {
			return fmt.Errorf("field %s %w: %w", fieldName, ErrInvalidTag, err)
		}
		if num > max {
			return fmt.Errorf("field %s must be at most %d, but got %d", fieldName, max, num)
		}
	case "in":
		vars := strings.Split(ruleParts[1], ",")
		for _, v := range vars {
			v, err := strconv.Atoi(v)
			if err != nil {
				return fmt.Errorf("field %s %w: %w", fieldName, ErrInvalidTag, err)
			}
			if num == v {
				return nil
			}
		}
		return fmt.Errorf("field %s must contain at least one (%s)", fieldName, ruleParts[1])
	}
	return nil
}
