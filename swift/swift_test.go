package swift

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

var (
	validCases = []struct {
		Swift        string
		BankCode     string
		CountryCode  string
		LocationCode string
		BranchCode   string
		Type         Type
	}{
		{
			Swift:        "TATRSKBX",
			BankCode:     "TATR",
			CountryCode:  "SK",
			LocationCode: "BX",
			BranchCode:   "",
			Type:         Type8,
		},
		{
			Swift:        "GIBASKBX",
			BankCode:     "GIBA",
			CountryCode:  "SK",
			LocationCode: "BX",
			BranchCode:   "",
			Type:         Type8,
		},
		{
			Swift:        "DEUTDEFF500",
			BankCode:     "DEUT",
			CountryCode:  "DE",
			LocationCode: "FF",
			BranchCode:   "500",
			Type:         Type11,
		},
	}
	invalidCases = []struct {
		Swift string
		Error error
	}{
		{
			Swift: "",
			Error: ErrInvalidLength,
		},
		{
			Swift: "KU78N78",
			Error: ErrInvalidLength,
		},
		{
			Swift: "KU78N78K43KL",
			Error: ErrInvalidLength,
		},
		{
			Swift: "MK23MjK2",
			Error: ErrInvalidCase,
		},
		{
			Swift: "MK23MjK2D23",
			Error: ErrInvalidCase,
		},
		{
			Swift: "MK23KDLF",
			Error: ErrInvalidBankCode,
		},
		{
			Swift: "24KM3KDLFDS",
			Error: ErrInvalidBankCode,
		},
		{
			Swift: "JMKM3KDL",
			Error: ErrInvalidCountryCode,
		},
		{
			Swift: "JMKM3KDLFDS",
			Error: ErrInvalidCountryCode,
		},
		{
			Swift: "JMKMXXDL",
			Error: ErrCountryCodeNotPresent,
		},
		{
			Swift: "JMKMXXDLFDS",
			Error: ErrCountryCodeNotPresent,
		},
		{
			Swift: "JMKMSK--",
			Error: ErrInvalidLocationCode,
		},
		{
			Swift: "JMKMSK--DSL",
			Error: ErrInvalidLocationCode,
		},
		{
			Swift: "JMKMSKLDDS-",
			Error: ErrInvalidBranchCode,
		},
	}
)

type ValidateTestSuite struct {
	suite.Suite
}

func (s *ValidateTestSuite) TestValidateLength() {
	for _, cs := range validCases {
		err := validateLength(cs.Swift)
		s.NoError(err, "SWIFT = %s", cs.Swift)
	}
}

func (s *ValidateTestSuite) TestValidateCase() {
	for _, cs := range validCases {
		err := validateCase(cs.Swift)
		s.NoError(err, "SWIFT = %s", cs.Swift)
	}
}

func (s *ValidateTestSuite) TestValidateBankCode() {
	for _, cs := range validCases {
		err := validateBankCode(cs.Swift)
		s.NoError(err, "SWIFT = %s", cs.Swift)
	}
}

func (s *ValidateTestSuite) TestValidateCountryCode() {
	for _, cs := range validCases {
		err := validateCountryCode(cs.Swift)
		s.NoError(err, "SWIFT = %s", cs.Swift)
	}
}

func (s *ValidateTestSuite) TestValidateLocationCode() {
	for _, cs := range validCases {
		err := validateLocationCode(cs.Swift)
		s.NoError(err, "SWIFT = %s", cs.Swift)
	}
}

func (s *ValidateTestSuite) TestValidateBranchCode() {
	for _, cs := range validCases {
		err := validateBranchCode(cs.Swift)
		s.NoError(err, "SWIFT = %s", cs.Swift)
	}
}

func (s *ValidateTestSuite) TestNew() {
	for _, cs := range validCases {
		sw, err := New(cs.Swift)
		s.NotNil(sw, "SWIFT = %s", cs.Swift)
		s.Nil(err, "SWIFT = %s", cs.Swift)
		s.Equal(cs.BankCode, sw.BankCode(), "SWIFT = %s", cs.Swift)
		s.Equal(cs.CountryCode, sw.CountryCode(), "SWIFT = %s", cs.Swift)
		s.Equal(cs.LocationCode, sw.LocationCode(), "SWIFT = %s", cs.Swift)
		s.Equal(cs.BranchCode, sw.BranchCode(), "SWIFT = %s", cs.Swift)
		s.Equal(cs.Type, sw.Type(), "SWIFT = %s", cs.Swift)
	}
}

func (s *ValidateTestSuite) TestMustParse() {
	for _, cs := range validCases {
		s.NotPanics(func() {
			sw := MustParse(cs.Swift)
			s.NotNil(sw, "SWIFT = %s", cs.Swift)
		})
	}
}

func (s *ValidateTestSuite) TestNewInvalid() {
	for _, cs := range invalidCases {
		sw, err := New(cs.Swift)
		s.Nil(sw, "SWIFT = %s", cs.Swift)
		s.Error(err, "SWIFT = %s", cs.Swift)
		s.Equal(cs.Error, err, "SWIFT = %s", cs.Swift)
	}
}

func (s *ValidateTestSuite) TestMustParseInvalid() {
	for _, cs := range invalidCases {
		s.Panics(func() {
			MustParse(cs.Swift)
		}, "SWIFT = %s", cs.Swift)
	}
}

func TestValidateTestSuite(t *testing.T) {
	suite.Run(t, new(ValidateTestSuite))
}
