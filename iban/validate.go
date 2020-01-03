package iban

import (
	"strconv"
	"unicode"

	"github.com/jbub/banking/bban"
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

func validateMinLength(value string) error {
	if len(value) < minIbanSize {
		return ErrIbanTooShort
	}
	return nil
}

func validateCountryCode(code string) error {
	for _, r := range code {
		if unicode.IsLower(r) {
			return ErrCountryCodeNotUpper
		}
		if !unicode.IsLetter(r) {
			return ErrCountryCodeNotAlpha
		}
	}
	return nil
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

func validateBban(bbn string, struc bban.Structure) error {
	if len(bbn) != struc.Length() {
		return ErrInvalidBbanLength
	}

	offset := 0
	for _, part := range struc.Parts() {
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
	return "0" + strconv.Itoa(check), nil
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
			total %= modValue
		}
	}

	return total % modValue, nil
}

func reformatIban(value string) string {
	return extractBban(value) + extractCountryCode(value) + extractCheckDigit(value)
}

func replaceCheckDigit(value string, code string) string {
	return code + defaultCheckDigit + extractBban(value)
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

func extractBbanPart(value string, struc bban.Structure, entryType bban.EntryType) string {
	bbn := extractBban(value)
	offset := 0
	for _, part := range struc.Parts() {
		value := bbn[offset : offset+part.Length]
		if part.EntryType == entryType {
			return value
		}
		offset += part.Length
	}
	return ""
}

func extractBankCode(value string, struc bban.Structure) string {
	return extractBbanPart(value, struc, bban.BankCode)
}

func extractBranchCode(value string, struc bban.Structure) string {
	return extractBbanPart(value, struc, bban.BranchCode)
}

func extractAccountNumber(value string, struc bban.Structure) string {
	return extractBbanPart(value, struc, bban.AccountNumber)
}

func extractNationalCheckDigit(value string, struc bban.Structure) string {
	return extractBbanPart(value, struc, bban.NationalCheckDigit)
}

func extractAccountType(value string, struc bban.Structure) string {
	return extractBbanPart(value, struc, bban.AccountType)
}

func extractOwnerAccountType(value string, struc bban.Structure) string {
	return extractBbanPart(value, struc, bban.OwnerAccountType)
}

func extractIdentificationNumber(value string, struc bban.Structure) string {
	return extractBbanPart(value, struc, bban.IdentificationNumber)
}

func extractCurrency(value string, struc bban.Structure) string {
	return extractBbanPart(value, struc, bban.Currency)
}
