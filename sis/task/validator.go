package task

import (
	"fmt"
	"strings"
	"unicode"
	"unicode/utf8"
)

const MaxPasswordLength = 128
const MinPasswordLength = 8

type Processor interface {
	Process(input string) (Result, error)
}

type PasswordValidator struct{}

var commonPasswords = map[string]bool{
	"password":  true,
	"password1": true,
	"12345678":  true,
	"123456789": true,
	"qwerty123": true,
	"iloveyou":  true,
	"admin123":  true,
	"letmein1":  true,
	"abc12345":  true,
}

func ValidateInput(input string) error {
	if strings.TrimSpace(input) == "" {
		return ErrEmptyPassword
	}
	if !utf8.ValidString(input) {
		return fmt.Errorf("%w: password is not valid UTF-8", ErrInvalidInput)
	}
	if utf8.RuneCountInString(input) > MaxPasswordLength {
		return fmt.Errorf("%w: password is longer than %d characters", ErrInvalidInput, MaxPasswordLength)
	}
	return nil
}

func (PasswordValidator) Process(input string) (Result, error) {
	err := ValidateInput(input)
	if err != nil {
		return Result{}, err
	}

	hasUpper := false
	hasLower := false
	hasDigit := false
	hasSpecial := false
	hasSpace := false
	for _, r := range input {
		switch {
		case unicode.IsUpper(r):
			hasUpper = true
		case unicode.IsLower(r):
			hasLower = true
		case unicode.IsDigit(r):
			hasDigit = true
		case unicode.IsSpace(r):
			hasSpace = true
		default:
			hasSpecial = true
		}
	}
	isCommon := commonPasswords[strings.ToLower(input)]
	length := utf8.RuneCountInString(input)

	rules := []struct {
		name   string
		passed bool
	}{
		{"at least 8 characters", length >= MinPasswordLength},
		{"an uppercase letter", hasUpper},
		{"a lowercase letter", hasLower},
		{"a digit", hasDigit},
		{"a special character", hasSpecial},
		{"no spaces", !hasSpace},
		{"not a common password", !isCommon},
	}

	result := Result{Length: length}
	for _, rule := range rules {
		if rule.passed {
			result.Score++
		} else {
			result.Failed = append(result.Failed, rule.name)
		}
	}

	switch {
	case isCommon || result.Score <= 4:
		result.Strength = Weak
	case result.Score < len(rules):
		result.Strength = Medium
	default:
		result.Strength = Strong
	}
	return result, nil
}
