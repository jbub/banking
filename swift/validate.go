package swift

import (
	"regexp"
	"strings"

	"github.com/jbub/banking/country"
)

const (
	// lengthSwift8 represents length of type Swift8 swift codes
	lengthSwift8 = 8

	// lengthSwift11 represents length of type Swift11 swift codes
	lengthSwift11 = 11
)

var (
	// regexBankCode holds Regexp for matching bank codes
	regexBankCode = regexp.MustCompile("^[A-Z]+$")

	// regexCountryCode holds Regexp for matching country codes
	regexCountryCode = regexp.MustCompile("^[A-Z]+$")

	// regexLocationCode holds Regexp for matching location codes
	regexLocationCode = regexp.MustCompile("^[A-Z0-9]+$")

	// regexBranchCode holds Regexp for matching location codes
	regexBranchCode = regexp.MustCompile("^[A-Z0-9]+$")
)

func validateLength(value string) error {
	if l := len(value); l != lengthSwift8 && l != lengthSwift11 {
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
	if code := extractBankCode(value); !regexBankCode.MatchString(code) {
		return ErrInvalidBankCode
	}
	return nil
}

func validateCountryCode(value string) error {
	code := extractCountryCode(value)
	if !regexCountryCode.MatchString(code) {
		return ErrInvalidCountryCode
	}

	if !country.Exists(code) {
		return ErrCountryCodeNotPresent
	}

	return nil
}

func validateLocationCode(value string) error {
	if code := extractLocationCode(value); !regexLocationCode.MatchString(code) {
		return ErrInvalidLocationCode
	}
	return nil
}

func validateBranchCode(value string) error {
	if hasBranchCode(value) {
		if code := extractBranchCode(value); !regexBranchCode.MatchString(code) {
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
