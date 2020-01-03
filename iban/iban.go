package iban

import (
	"errors"

	"github.com/jbub/banking/bban"
	"github.com/jbub/banking/country"
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

// Iban represents iban code. Zero value is not usable.
type Iban struct {
	value string
	struc bban.Structure
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
	return extractAccountNumber(i.value, i.struc)
}

// BankCode returns bank code of iban.
func (i *Iban) BankCode() string {
	return extractBankCode(i.value, i.struc)
}

// BranchCode returns branch code of iban.
func (i *Iban) BranchCode() string {
	return extractBranchCode(i.value, i.struc)
}

// NationalCheckDigit returns national check digit of iban.
func (i *Iban) NationalCheckDigit() string {
	return extractNationalCheckDigit(i.value, i.struc)
}

// AccountType returns account type of iban.
func (i *Iban) AccountType() string {
	return extractAccountType(i.value, i.struc)
}

// OwnerAccountType returns owner account type of iban.
func (i *Iban) OwnerAccountType() string {
	return extractOwnerAccountType(i.value, i.struc)
}

// IdentificationNumber returns identification number of iban.
func (i *Iban) IdentificationNumber() string {
	return extractIdentificationNumber(i.value, i.struc)
}

// Currency returns currency of iban.
func (i *Iban) Currency() string {
	return extractCurrency(i.value, i.struc)
}

// String returns text representation of iban.
func (i *Iban) String() string {
	return i.value
}

// Validate validates iban code.
func Validate(value string) error {
	_, err := validate(value)
	return err
}

// New validates and creates new iban code.
// Deprecated: Use Parse instead.
func New(value string) (*Iban, error) {
	struc, err := validate(value)
	if err != nil {
		return nil, err
	}
	return &Iban{
		value: value,
		struc: struc,
	}, nil
}

// Parse validates and creates new iban code.
func Parse(value string) (*Iban, error) {
	return New(value)
}

// MustParse tries to create new iban code, panics on failure.
func MustParse(value string) *Iban {
	ibn, err := New(value)
	if err != nil {
		panic(err)
	}
	return ibn
}

func validate(value string) (bban.Structure, error) {
	if err := validateMinLength(value); err != nil {
		return bban.Structure{}, err
	}

	code := extractCountryCode(value)
	if err := validateCountryCode(code); err != nil {
		return bban.Structure{}, err
	}

	struc, ok := country.GetBbanStructure(code)
	if !ok {
		return bban.Structure{}, ErrCountryCodeNotPresent
	}

	bbn := extractBban(value)
	if err := validateBban(bbn, struc); err != nil {
		return struc, err
	}

	if err := validateCheckDigit(value, code); err != nil {
		return struc, err
	}
	return struc, nil
}
