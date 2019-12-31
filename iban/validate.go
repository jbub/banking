package iban

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/jbub/banking/bban"
	"github.com/jbub/banking/country"
)

const (
	// minIbanSize represents minimal length of iban.
	minIbanSize = 15

	// modCheck represents value used in mod check.
	modCheck = 98

	// modValue represents value used in mod check.
	modValue = 97

	// modMax is the maximum value allowed in mod check.
	modMax = 999999999

	// defaultCheckDigit is digit used in digit check.
	defaultCheckDigit = "00"
)

var (
	// regexCountryCode holds Regexp for matching country codes.
	regexCountryCode = regexp.MustCompile("^[A-Z]+$")
)

func validateMinLength(value string) error {
	if len(value) < minIbanSize {
		return ErrIbanTooShort
	}
	return nil
}

func validateCountryCode(value string) (string, error) {
	code := extractCountryCode(value)
	if code != strings.ToUpper(code) {
		return "", ErrCountryCodeNotUpper
	}

	if !regexCountryCode.MatchString(code) {
		return "", ErrCountryCodeNotAlpha
	}

	if !country.Exists(code) {
		return "", ErrCountryCodeNotPresent
	}

	return code, nil
}

func validateCheckDigit(value string, code string) error {
	digit := extractCheckDigit(value)
	expected, err := calculateCheckDigit(value, code)
	if err != nil {
		return err
	}

	if digit != expected {
		return ErrInvalidCheckDigit
	}

	return nil
}

func validateBbanLength(value string, structure bban.Structure) error {
	bbn := extractBban(value)
	if len(bbn) != structure.Length() {
		return ErrInvalidBbanLength
	}
	return nil
}

func validateBbanStructure(value string, structure bban.Structure) error {
	bbn := extractBban(value)
	offset := 0

	for _, part := range structure.Parts() {
		if value := bbn[offset : offset+part.Length]; !part.Validate(value) {
			return ErrInvalidBbanPart
		}

		offset += part.Length
	}

	return nil
}

func calculateCheckDigit(value string, code string) (string, error) {
	replaced := replaceCheckDigit(value, code)
	mod, err := calculateMod(replaced)
	if err != nil {
		return "", err
	}

	check := int(modCheck - mod)
	if check > 9 {
		return strconv.Itoa(check), nil
	}

	return fmt.Sprintf("0%d", check), nil
}

func calculateMod(value string) (int64, error) {
	var total int64

	for _, c := range reformatIban(value) {
		n, err := strconv.ParseInt(string(c), 36, 64)
		if err != nil {
			return 0, ErrInvalidIbanModulo
		}

		if n > 9 {
			total = total*100 + n
		} else {
			total = total*10 + n
		}

		if total > modMax {
			total = total % modValue
		}
	}

	return total % modValue, nil
}

func reformatIban(value string) string {
	return fmt.Sprintf("%s%s%s", extractBban(value), extractCountryCode(value), extractCheckDigit(value))
}

func replaceCheckDigit(value string, code string) string {
	return fmt.Sprintf("%s%s%s", code, defaultCheckDigit, extractBban(value))
}

func extractCountryCode(value string) string {
	return value[0:2]
}

func extractCheckDigit(value string) string {
	return value[2:4]
}

func extractBban(value string) string {
	return value[4:]
}

func extractBbanPart(value string, entryType bban.EntryType) string {
	bbn := extractBban(value)
	code := extractCountryCode(value)
	structure, _ := country.GetBbanStructure(code)
	offset := 0

	for _, part := range structure.Parts() {
		value := bbn[offset : offset+part.Length]
		if part.EntryType == entryType {
			return value
		}
		offset += part.Length
	}

	return ""
}

func extractBankCode(value string) string {
	return extractBbanPart(value, bban.BankCode)
}

func extractBranchCode(value string) string {
	return extractBbanPart(value, bban.BranchCode)
}

func extractAccountNumber(value string) string {
	return extractBbanPart(value, bban.AccountNumber)
}

func extractNationalCheckDigit(value string) string {
	return extractBbanPart(value, bban.NationalCheckDigit)
}

func extractAccountType(value string) string {
	return extractBbanPart(value, bban.AccountType)
}

func extractOwnerAccountType(value string) string {
	return extractBbanPart(value, bban.OwnerAccountType)
}

func extractIdentificationNumber(value string) string {
	return extractBbanPart(value, bban.IdentificationNumber)
}
