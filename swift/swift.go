package swift

import "errors"

// Error codes returned by failures to validate an swift.
var (
	ErrInvalidLength         = errors.New("swift: invalid length")
	ErrInvalidCase           = errors.New("swift: invalid case")
	ErrInvalidBankCode       = errors.New("swift: invalid bank code")
	ErrInvalidCountryCode    = errors.New("swift: invalid country code")
	ErrCountryCodeNotPresent = errors.New("swift: country code does not exist")
	ErrInvalidLocationCode   = errors.New("swift: invalid location code")
	ErrInvalidBranchCode     = errors.New("swift: invalid branch code")
)

// Type represents type of swift code.
type Type int

const (
	// Type8 represents swift code with a length of 8 characters.
	Type8 Type = iota

	// Type11 represents swift code with a length of 11 characters.
	Type11
)

// Swift represents a swift/bic code.
type Swift struct {
	value string
}

// BankCode returns bank code of swift code.
func (s *Swift) BankCode() string {
	return extractBankCode(s.value)
}

// CountryCode returns country code of swift code.
func (s *Swift) CountryCode() string {
	return extractCountryCode(s.value)
}

// LocationCode returns location code of swift code.
func (s *Swift) LocationCode() string {
	return extractLocationCode(s.value)
}

// BranchCode returns branch code of swift code.
func (s *Swift) BranchCode() string {
	if hasBranchCode(s.value) {
		return extractBranchCode(s.value)
	}
	return ""
}

// Type returns type of swift code.
func (s *Swift) Type() Type {
	if hasBranchCode(s.value) {
		return Type11
	}
	return Type8
}

// Validate validates swift code.
func Validate(value string) error {
	if err := validateLength(value); err != nil {
		return err
	}

	if err := validateCase(value); err != nil {
		return err
	}

	if err := validateBankCode(value); err != nil {
		return err
	}

	if err := validateCountryCode(value); err != nil {
		return err
	}

	if err := validateLocationCode(value); err != nil {
		return err
	}

	return validateBranchCode(value)
}

// New validates and creates new swift code.
// Deprecated: Use Parse instead.
func New(value string) (*Swift, error) {
	if err := Validate(value); err != nil {
		return nil, err
	}
	return &Swift{value: value}, nil
}

// Parse validates and creates new swift code.
func Parse(value string) (*Swift, error) {
	return New(value)
}

// MustParse tries to create new swift code, panics on failure.
func MustParse(value string) *Swift {
	swft, err := Parse(value)
	if err != nil {
		panic(err)
	}
	return swft
}
