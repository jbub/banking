package swift

import (
	"github.com/jbub/banking/country"
)

const (
	// lengthSwift8 represents length of type Swift8 swift codes.
	lengthSwift8 = 8

	// lengthSwift11 represents length of type Swift11 swift codes.
	lengthSwift11 = 11
)

func validateLength(value string) error {
	if l := len(value); l != lengthSwift8 && l != lengthSwift11 {
		return ErrInvalidLength
	}
	return nil
}

func validateCase(value string) error {
	for _, r := range value {
		if 'a' <= r && r <= 'z' {
			return ErrInvalidCase
		}
	}
	return nil
}

func validateBankCode(value string) error {
	if bankCode := extractBankCode(value); !validateAlpha(bankCode) {
		return ErrInvalidBankCode
	}
	return nil
}

func validateCountryCode(value string) error {
	code := extractCountryCode(value)
	if !validateAlpha(code) {
		return ErrInvalidCountryCode
	}

	if !country.Exists(code) {
		return ErrCountryCodeNotPresent
	}
	return nil
}

func validateLocationCode(value string) error {
	if code := extractLocationCode(value); !validateAlphaNum(code) {
		return ErrInvalidLocationCode
	}
	return nil
}

func validateBranchCode(value string) error {
	if !hasBranchCode(value) {
		return nil
	}

	if code := extractBranchCode(value); !validateAlphaNum(code) {
		return ErrInvalidBranchCode
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

func validateAlpha(s string) bool {
	return validateString(s, func(r rune) bool {
		return 'A' <= r && r <= 'Z'
	})
}

func validateAlphaNum(s string) bool {
	return validateString(s, func(r rune) bool {
		return 'A' <= r && r <= 'Z' || '0' <= r && r <= '9'
	})
}

func validateString(s string, validator func(rune) bool) bool {
	if s == "" {
		return false
	}
	for _, r := range s {
		if !validator(r) {
			return false
		}
	}
	return true
}
