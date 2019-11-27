package iban

import (
	"errors"

	"github.com/anecsoiu/banking/country"
)

// Error codes returned by failures to validate an iban.
var (
	ErrIbanTooShort          = errors.New("iban: iban too short")
	ErrCountryCodeNotUpper   = errors.New("iban: country code contains lowercase letters")
	ErrCountryCodeNotAlpha   = errors.New("iban: country code contains non alphabetic letters")
	ErrCountryCodeNotPresent = errors.New("iban: country code does not exist")
	ErrInvalidCheckDigit     = errors.New("iban: invalid check digit")
	ErrInvalidIbanModulo     = errors.New("iban: invalid modulo")
	ErrInvalidBbanLength     = errors.New("iban: invalid bban length")
	ErrInvalidBbanPart       = errors.New("iban: invalid bban part")
)

// Iban represents iban code.
type Iban struct {
	value string
}

// CheckDigit returns check digit of iban.
func (i *Iban) CheckDigit() string {
	return extractCheckDigit(i.value)
}

// CountryCode returns country code of iban.
func (i *Iban) CountryCode() string {
	return extractCountryCode(i.value)
}

// Bban returns bban part of iban.
func (i *Iban) Bban() string {
	return extractBban(i.value)
}

// AccountNumber returns account number of iban.
func (i *Iban) AccountNumber() string {
	return extractAccountNumber(i.value)
}

// BankCode returns bank code of iban.
func (i *Iban) BankCode() string {
	return extractBankCode(i.value)
}

// BranchCode returns branch code of iban.
func (i *Iban) BranchCode() string {
	return extractBranchCode(i.value)
}

// NationalCheckDigit returns national check digit of iban.
func (i *Iban) NationalCheckDigit() string {
	return extractNationalCheckDigit(i.value)
}

// AccountType returns account type of iban.
func (i *Iban) AccountType() string {
	return extractAccountType(i.value)
}

// OwnerAccountType returns owner account type of iban.
func (i *Iban) OwnerAccountType() string {
	return extractOwnerAccountType(i.value)
}

// IdentificationNumber returns identification number of iban.
func (i *Iban) IdentificationNumber() string {
	return extractIdentificationNumber(i.value)
}

// String returns text representation of iban.
func (i *Iban) String() string {
	return i.value
}

// Validate validates iban code.
func Validate(value string) error {
	if err := validateMinLength(value); err != nil {
		return err
	}

	code, err := validateCountryCode(value)
	if err != nil {
		return err
	}

	structure, ok := country.GetBbanStructure(code)
	if !ok {
		return ErrCountryCodeNotPresent
	}

	if err = validateBbanLength(value, structure); err != nil {
		return err
	}

	if err = validateBbanStructure(value, structure); err != nil {
		return err
	}

	if err = validateCheckDigit(value, code); err != nil {
		return err
	}

	return nil
}

// New validates and creates new iban code.
func New(value string) (*Iban, error) {
	if err := Validate(value); err != nil {
		return nil, err
	}
	return &Iban{value}, nil
}

// MustParse tries to create new iban code, panics on failure.
func MustParse(value string) *Iban {
	ibn, err := New(value)
	if err != nil {
		panic(err)
	}
	return ibn
}
