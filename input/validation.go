package input

import (
	"errors"
	"net/mail"
	"net/url"
	"regexp"
	"strconv"
)

// ValidationRule is an alias for string representing a validation rule.
type ValidationRule string

// ValidateFunc is a function that returns an error if the input is invalid.
type ValidateFunc func(string) error

const (
	// Alphanumeric is the string representing the alphanumeric validation rule.
	Alphanumeric ValidationRule = "alphanumeric"
	// DigitsOnly is the string representing the alphanumeric validation rule.
	DigitsOnly ValidationRule = "digits"
	// Integers is the string representing the alphanumeric validation rule.
	Integers ValidationRule = "integers"
	// Floats is the string representing the alphanumeric validation rule.
	Floats ValidationRule = "floats"
	// EmailAddress is the string representing the alphanumeric validation rule.
	EmailAddress ValidationRule = "email"
	// URL is the string representing the alphanumeric validation rule.
	URL ValidationRule = "url"
)

var ruleErrorMsgMap = map[ValidationRule]string{
	Alphanumeric: "it must be an alphanumeric value",
	DigitsOnly:   "it must be a digits only value",
	Integers:     "it must be an integer value",
	Floats:       "it must be a float value",
	EmailAddress: "it must be a valid email address",
	URL:          "it must be a valid URL",
}

// ===============================================================

// ValidateAlphanumeric is the validation function to validate an alphanumeric string.
func ValidateAlphanumeric(s string) error {
	return validate(s, Alphanumeric)
}

// ValidateDigits is the validation function to validate a digits only string.
func ValidateDigits(s string) error {
	return validate(s, DigitsOnly)
}

// ValidateInteger is the validation function to validate an integer string.
func ValidateInteger(s string) error {
	return validate(s, Integers)
}

// ValidateFloat is the validation function to validate a float string.
func ValidateFloat(s string) error {
	return validate(s, Floats)
}

// ValidateEmail is the validation function to validate an email address string.
func ValidateEmail(s string) error {
	return validate(s, EmailAddress)
}

// ValidateURL is the validation function to validate an URL string.
func ValidateURL(s string) error {
	return validate(s, URL)
}

// ===============================================================

func alphanumericMatcher(s string) bool {
	re := regexp.MustCompile("^[a-zA-Z0-9 ]*$")
	return re.MatchString(s)
}
func digitsOnlyMatcher(s string) bool {
	re := regexp.MustCompile(`\d`)
	return re.MatchString(s)
}

func integersMatch(s string) bool {
	if _, err := strconv.ParseInt(s, 10, 64); err == nil {
		return true
	}
	return false
}

func floatMatcher(s string) bool {
	if _, err := strconv.ParseInt(s, 10, 64); err == nil {
		return true
	}
	return false
}

func emailMatcher(s string) bool {
	if _, err := mail.ParseAddress(s); err == nil {
		return true
	}
	return false
}

func urlMarcher(s string) bool {
	if _, err := url.ParseRequestURI(s); err == nil {
		return true
	}
	return false
}

// ===============================================================

func validate(field string, rule ValidationRule) error {
	errorMsg := ""

	switch rule {
	case Alphanumeric:
		if !alphanumericMatcher(field) {
			errorMsg = ruleErrorMsgMap[rule]
		}
	case DigitsOnly:
		if !digitsOnlyMatcher(field) {
			errorMsg = ruleErrorMsgMap[rule]
		}
	case Integers:
		if !integersMatch(field) {
			errorMsg = ruleErrorMsgMap[rule]
		}
	case Floats:
		if !floatMatcher(field) {
			errorMsg = ruleErrorMsgMap[rule]
		}
	case EmailAddress:
		if !emailMatcher(field) {
			errorMsg = ruleErrorMsgMap[rule]
		}
	case URL:
		if !urlMarcher(field) {
			errorMsg = ruleErrorMsgMap[rule]
		}
	default:
		errorMsg = "unknown validation rule"
	}

	if !isEmpty(errorMsg) {
		return errors.New(errorMsg)
	}
	return nil
}
