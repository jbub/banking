package swift

import (
	"regexp"
	"strings"

	"github.com/jbub/banking/country"
)

const (
	lengthSwift8  = 8
	lengthSwift11 = 11
)

func validateLength(value string) error {
	length := len(value)
	if length != lengthSwift8 && length != lengthSwift11 {
		return ErrInvalidLength
	}
	return nil
}

func validateCase(value string) error {
	if value != strings.ToUpper(value) {
		return ErrInvalidCase
	}
	return nil
}

func validateBankCode(value string) error {
	code := extractBankCode(value)
	match, _ := regexp.MatchString("^[A-Z]+$", code)
	if !match {
		return ErrInvalidBankCode
	}
	return nil
}

func validateCountryCode(value string) error {
	code := extractCountryCode(value)
	match, _ := regexp.MatchString("^[A-Z]+$", code)
	if !match {
		return ErrInvalidCountryCode
	}

	exists := country.Exists(code)
	if !exists {
		return ErrCountryCodeNotPresent
	}

	return nil
}

func validateLocationCode(value string) error {
	code := extractLocationCode(value)
	match, _ := regexp.MatchString("^[A-Z0-9]+$", code)
	if !match {
		return ErrInvalidLocationCode
	}
	return nil
}

func validateBranchCode(value string) error {
	if hasBranchCode(value) {
		code := extractBranchCode(value)
		match, _ := regexp.MatchString("^[A-Z0-9]+$", code)
		if !match {
			return ErrInvalidBranchCode
		}
	}
	return nil
}

func extractBankCode(value string) string {
	return value[0:4]
}

func extractCountryCode(value string) string {
	return value[4:6]
}

func extractLocationCode(value string) string {
	return value[6:8]
}

func extractBranchCode(value string) string {
	return value[8:11]
}

func hasBranchCode(value string) bool {
	return len(value) == lengthSwift11
}
