package bban

import (
	"unicode"
)

// EntryType represents a type of bban part.
type EntryType int

// CharType represents a character type of given bban part.
type charType int

const (
	// BankCode represents bank code part of iban.
	BankCode EntryType = iota

	// BranchCode represents branch code part of iban.
	BranchCode

	// AccountNumber represents account number part of iban.
	AccountNumber

	// NationalCheckDigit represents national check digit part of iban.
	NationalCheckDigit

	// AccountType represents account type part of iban.
	AccountType

	// OwnerAccountType represents owner account type part of iban.
	OwnerAccountType

	// IdentificationNumber represents identification number part of iban.
	IdentificationNumber

	// Num allows only numeric characters.
	Num charType = iota

	// AlphaUpper allows only uppercase alphabetic characters.
	AlphaUpper

	// AlphaNum allow only alphanumeric characters with any case.
	AlphaNum
)

// String returns text representation of EntryType.
func (e EntryType) String() string {
	switch e {
	case BankCode:
		return "BankCode"
	case BranchCode:
		return "BranchCode"
	case AccountNumber:
		return "AccountNumber"
	case NationalCheckDigit:
		return "NationalCheckDigit"
	case AccountType:
		return "AccountType"
	case OwnerAccountType:
		return "OwnerAccountType"
	case IdentificationNumber:
		return "IdentificationNumber"
	}
	return ""
}

// Validate validates given value against current CharType.
func (c charType) Validate(s string) bool {
	if s == "" {
		return false
	}
	switch c {
	case Num:
		return validateNum(s)
	case AlphaUpper:
		return validateAlphaUpper(s)
	case AlphaNum:
		return validateAlphaNum(s)
	}
	return false
}

func validateNum(s string) bool {
	return validateString(s, unicode.IsDigit)
}

func validateAlphaUpper(s string) bool {
	return validateString(s, func(r rune) bool {
		return unicode.IsUpper(r) && unicode.IsLetter(r)
	})
}

func validateAlphaNum(s string) bool {
	return validateString(s, func(r rune) bool {
		return unicode.IsLetter(r) || unicode.IsDigit(r)
	})
}

func validateString(s string, validator func(rune) bool) bool {
	for _, r := range s {
		if !validator(r) {
			return false
		}
	}
	return true
}

// Structure represents a bban structure which consists of bban Parts.
type Structure struct {
	parts []Part
}

// Parts returns slice of Parts.
func (s Structure) Parts() []Part {
	return s.parts
}

// Length returns a length of bban calculated by summing lengths of each Part.
func (s Structure) Length() int {
	length := 0
	for _, p := range s.Parts() {
		length += p.Length
	}
	return length
}

// NewStructure creates a new Structure from given Parts.
func NewStructure(parts ...Part) Structure {
	s := Structure{}
	s.parts = parts
	return s
}
