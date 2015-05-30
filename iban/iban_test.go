package iban

import (
	"testing"

	"github.com/jbub/banking/country"
	"github.com/stretchr/testify/suite"
)

var (
	validCases = []struct {
		Iban                 string
		CountryCode          string
		Digit                string
		Bban                 string
		Replaced             string
		BankCode             string
		BranchCode           string
		AccountNumber        string
		IdentificationNumber string
		AccountType          string
		OwnerAccountType     string
		NationalCheckDigit   string
	}{
		{
			Iban:               "AL47212110090000000235698741",
			CountryCode:        "AL",
			Digit:              "47",
			Bban:               "212110090000000235698741",
			Replaced:           "AL00212110090000000235698741",
			BankCode:           "212",
			BranchCode:         "1100",
			AccountNumber:      "0000000235698741",
			NationalCheckDigit: "9",
		},
		{
			Iban:               "BE68539007547034",
			CountryCode:        "BE",
			Digit:              "68",
			Bban:               "539007547034",
			Replaced:           "BE00539007547034",
			BankCode:           "539",
			AccountNumber:      "0075470",
			NationalCheckDigit: "34",
		},
		{
			Iban:          "GR1601101250000000012300695",
			CountryCode:   "GR",
			Digit:         "16",
			Bban:          "01101250000000012300695",
			Replaced:      "GR0001101250000000012300695",
			BankCode:      "011",
			BranchCode:    "0125",
			AccountNumber: "0000000012300695",
		},
		{
			Iban:          "GB29NWBK60161331926819",
			CountryCode:   "GB",
			Digit:         "29",
			Bban:          "NWBK60161331926819",
			Replaced:      "GB00NWBK60161331926819",
			BankCode:      "NWBK",
			BranchCode:    "601613",
			AccountNumber: "31926819",
		},
		{
			Iban:          "SA0380000000608010167519",
			CountryCode:   "SA",
			Digit:         "03",
			Bban:          "80000000608010167519",
			Replaced:      "SA0080000000608010167519",
			BankCode:      "80",
			AccountNumber: "000000608010167519",
		},
		{
			Iban:          "CH9300762011623852957",
			CountryCode:   "CH",
			Digit:         "93",
			Bban:          "00762011623852957",
			Replaced:      "CH0000762011623852957",
			BankCode:      "00762",
			AccountNumber: "011623852957",
		},
		{
			Iban:               "TR330006100519786457841326",
			CountryCode:        "TR",
			Digit:              "33",
			Bban:               "0006100519786457841326",
			Replaced:           "TR000006100519786457841326",
			BankCode:           "00061",
			AccountNumber:      "0519786457841326",
			NationalCheckDigit: "0",
		},
		{
			Iban:               "PL60102010260000042270201111",
			CountryCode:        "PL",
			Digit:              "60",
			Bban:               "102010260000042270201111",
			Replaced:           "PL00102010260000042270201111",
			BankCode:           "102",
			BranchCode:         "0102",
			AccountNumber:      "0000042270201111",
			NationalCheckDigit: "6",
		},
		{
			Iban:          "SK0611000000002920884960",
			CountryCode:   "SK",
			Digit:         "06",
			Bban:          "11000000002920884960",
			Replaced:      "SK0011000000002920884960",
			BankCode:      "1100",
			AccountNumber: "0000002920884960",
		},
		{
			Iban:               "NO9386011117947",
			CountryCode:        "NO",
			Digit:              "93",
			Bban:               "86011117947",
			Replaced:           "NO0086011117947",
			BankCode:           "8601",
			AccountNumber:      "111794",
			NationalCheckDigit: "7",
		},
	}
	invalidCases = []struct {
		Iban  string
		Error error
	}{
		{
			Iban:  "",
			Error: ErrIbanTooShort,
		},
		{
			Iban:  "AL4721211",
			Error: ErrIbanTooShort,
		},
		{
			Iban:  "sL472121142342323",
			Error: ErrCountryCodeNotUpper,
		},
		{
			Iban:  "S4472121142342323",
			Error: ErrCountryCodeNotAlpha,
		},
		{
			Iban:  "XX472121142342323",
			Error: ErrCountryCodeNotPresent,
		},
		{
			Iban:  "CH9300762011-623852957",
			Error: ErrInvalidBbanLength,
		},
		{
			Iban:  "SK0X2121142342323",
			Error: ErrInvalidBbanLength,
		},
		{
			Iban:  "PL67102010260000042270201111",
			Error: ErrInvalidCheckDigit,
		},
		{
			Iban:  "AT61190430023457320",
			Error: ErrInvalidBbanLength,
		},
		{
			Iban:  "SK061100A000002920884960",
			Error: ErrInvalidBbanPart,
		},
	}
	invalidBbanCases = []struct {
		Iban        string
		CountryCode string
		Error       error
	}{
		{
			Iban:        "SA03800000006080101675194",
			CountryCode: "SA",
			Error:       ErrInvalidBbanLength,
		},
		{
			Iban:        "GB29NWB161331926819",
			CountryCode: "GB",
			Error:       ErrInvalidBbanLength,
		},
		{
			Iban:        "BE68507547034",
			CountryCode: "BE",
			Error:       ErrInvalidBbanLength,
		},
		{
			Iban:        "SK0611000002920884960",
			CountryCode: "SK",
			Error:       ErrInvalidBbanLength,
		},
	}
)

type ValidateTestSuite struct {
	suite.Suite
}

func (s *ValidateTestSuite) TestValidateMinLength() {
	for _, cs := range validCases {
		err := validateMinLength(cs.Iban)
		s.NoError(err, "IBAN = %s", cs.Iban)
	}
}

func (s *ValidateTestSuite) TestExtractCountryCode() {
	for _, cs := range validCases {
		code := extractCountryCode(cs.Iban)
		s.Equal(cs.CountryCode, code, "IBAN = %s", cs.Iban)
	}
}

func (s *ValidateTestSuite) TestExtractDigit() {
	for _, cs := range validCases {
		digit := extractCheckDigit(cs.Iban)
		s.Equal(cs.Digit, digit, "IBAN = %s", cs.Iban)
	}
}

func (s *ValidateTestSuite) TestExtractBban() {
	for _, cs := range validCases {
		bban := extractBban(cs.Iban)
		s.Equal(cs.Bban, bban, "IBAN = %s", cs.Iban)
	}
}

func (s *ValidateTestSuite) TestCountryCodeExists() {
	for _, cs := range validCases {
		code := extractCountryCode(cs.Iban)
		exists := country.Exists(code)
		s.True(exists, "IBAN = %s", cs.Iban)
	}
}

func (s *ValidateTestSuite) TestReplaceDigit() {
	for _, cs := range validCases {
		code := extractCountryCode(cs.Iban)
		replaced := replaceCheckDigit(cs.Iban, code)
		s.Equal(cs.Replaced, replaced, "IBAN = %s", cs.Iban)
	}
}

func (s *ValidateTestSuite) TestValidateDigit() {
	for _, cs := range validCases {
		code := extractCountryCode(cs.Iban)
		err := validateCheckDigit(cs.Iban, code)
		s.NoError(err, "IBAN = %s", cs.Iban)
	}
}

func (s *ValidateTestSuite) TestValidateCountryCode() {
	for _, cs := range validCases {
		code, err := validateCountryCode(cs.Iban)
		s.NoError(err, "IBAN = %s", cs.Iban)
		s.Equal(cs.CountryCode, code, "IBAN = %s", cs.Iban)
	}
}

func (s *ValidateTestSuite) TestCalculateMod() {
	for _, cs := range validCases {
		mod, err := calculateMod(cs.Iban)
		s.NoError(err, "IBAN = %s", cs.Iban)
		s.Equal(mod, int64(1), "IBAN = %s", cs.Iban)
	}
}

func (s *ValidateTestSuite) TestValidateBbanLength() {
	for _, cs := range validCases {
		structure, ok := country.GetBbanStructure(cs.CountryCode)
		s.True(ok, "IBAN = %s", cs.Iban)
		err := validateBbanLength(cs.Iban, structure)
		s.NoError(err, "IBAN = %s", cs.Iban)
	}
}

func (s *ValidateTestSuite) TestValidateBbanLengthInvalid() {
	for _, cs := range invalidBbanCases {
		structure, ok := country.GetBbanStructure(cs.CountryCode)
		s.True(ok, "IBAN = %s", cs.Iban)
		err := validateBbanLength(cs.Iban, structure)
		s.Error(err, "IBAN = %s", cs.Iban)
		s.Equal(cs.Error, ErrInvalidBbanLength, "IBAN = %s", cs.Iban)
	}
}

func (s *ValidateTestSuite) TestValidateBbanStructure() {
	for _, cs := range validCases {
		structure, ok := country.GetBbanStructure(cs.CountryCode)
		s.True(ok, "IBAN = %s", cs.Iban)
		err := validateBbanStructure(cs.Iban, structure)
		s.NoError(err, "IBAN = %s", cs.Iban)
	}
}

func (s *ValidateTestSuite) TestNew() {
	for _, cs := range validCases {
		ib, err := New(cs.Iban)
		s.NotNil(ib)
		s.NoError(err)
		s.Equal(cs.BankCode, ib.BankCode(), "IBAN = %s", cs.Iban)
		s.Equal(cs.BranchCode, ib.BranchCode(), "IBAN = %s", cs.Iban)
		s.Equal(cs.AccountNumber, ib.AccountNumber(), "IBAN = %s", cs.Iban)
		s.Equal(cs.IdentificationNumber, ib.IdentificationNumber(), "IBAN = %s", cs.Iban)
		s.Equal(cs.AccountType, ib.AccountType(), "IBAN = %s", cs.Iban)
		s.Equal(cs.OwnerAccountType, ib.OwnerAccountType(), "IBAN = %s", cs.Iban)
		s.Equal(cs.NationalCheckDigit, ib.NationalCheckDigit(), "IBAN = %s", cs.Iban)
		s.Equal(cs.Digit, ib.CheckDigit(), "IBAN = %s", cs.Iban)
		s.Equal(cs.CountryCode, ib.CountryCode(), "IBAN = %s", cs.Iban)
		s.Equal(cs.Bban, ib.Bban(), "IBAN = %s", cs.Iban)
	}
}

func (s *ValidateTestSuite) TestMustParse() {
	for _, cs := range validCases {
		s.NotPanics(func() {
			ib := MustParse(cs.Iban)
			s.NotNil(ib, "IBAN = %s", cs.Iban)
		})
	}
}

func (s *ValidateTestSuite) TestNewInvalid() {
	for _, cs := range invalidCases {
		ib, err := New(cs.Iban)
		s.Nil(ib)
		s.Error(err, "IBAN = %s", ib)
		s.Equal(cs.Error, err, "IBAN = %s", cs.Iban)
	}
}

func (s *ValidateTestSuite) TestMustParseInvalid() {
	for _, cs := range invalidCases {
		s.Panics(func() {
			MustParse(cs.Iban)
		}, "IBAN = %s", cs.Iban)
	}
}

func TestValidateTestSuite(t *testing.T) {
	suite.Run(t, new(ValidateTestSuite))
}
