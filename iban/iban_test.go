package iban

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/jbub/banking/country"
)

var (
	validCases = []struct {
		iban                 string
		countryCode          string
		checkDigit           string
		bban                 string
		replaced             string
		bankCode             string
		branchCode           string
		accountNumber        string
		identificationNumber string
		accountType          string
		ownerAccountType     string
		nationalCheckDigit   string
		currency             string
	}{
		{
			iban:               "AL47212110090000000235698741",
			countryCode:        "AL",
			checkDigit:         "47",
			bban:               "212110090000000235698741",
			replaced:           "AL00212110090000000235698741",
			bankCode:           "212",
			branchCode:         "1100",
			accountNumber:      "0000000235698741",
			nationalCheckDigit: "9",
		},
		{
			iban:               "BE68539007547034",
			countryCode:        "BE",
			checkDigit:         "68",
			bban:               "539007547034",
			replaced:           "BE00539007547034",
			bankCode:           "539",
			accountNumber:      "0075470",
			nationalCheckDigit: "34",
		},
		{
			iban:          "GR1601101250000000012300695",
			countryCode:   "GR",
			checkDigit:    "16",
			bban:          "01101250000000012300695",
			replaced:      "GR0001101250000000012300695",
			bankCode:      "011",
			branchCode:    "0125",
			accountNumber: "0000000012300695",
		},
		{
			iban:          "GB29NWBK60161331926819",
			countryCode:   "GB",
			checkDigit:    "29",
			bban:          "NWBK60161331926819",
			replaced:      "GB00NWBK60161331926819",
			bankCode:      "NWBK",
			branchCode:    "601613",
			accountNumber: "31926819",
		},
		{
			iban:          "SA0380000000608010167519",
			countryCode:   "SA",
			checkDigit:    "03",
			bban:          "80000000608010167519",
			replaced:      "SA0080000000608010167519",
			bankCode:      "80",
			accountNumber: "000000608010167519",
		},
		{
			iban:          "CH9300762011623852957",
			countryCode:   "CH",
			checkDigit:    "93",
			bban:          "00762011623852957",
			replaced:      "CH0000762011623852957",
			bankCode:      "00762",
			accountNumber: "011623852957",
		},
		{
			iban:               "TR330006100519786457841326",
			countryCode:        "TR",
			checkDigit:         "33",
			bban:               "0006100519786457841326",
			replaced:           "TR000006100519786457841326",
			bankCode:           "00061",
			accountNumber:      "0519786457841326",
			nationalCheckDigit: "0",
		},
		{
			iban:               "PL60102010260000042270201111",
			countryCode:        "PL",
			checkDigit:         "60",
			bban:               "102010260000042270201111",
			replaced:           "PL00102010260000042270201111",
			bankCode:           "102",
			branchCode:         "0102",
			accountNumber:      "0000042270201111",
			nationalCheckDigit: "6",
		},
		{
			iban:          "SK0611000000002920884960",
			countryCode:   "SK",
			checkDigit:    "06",
			bban:          "11000000002920884960",
			replaced:      "SK0011000000002920884960",
			bankCode:      "1100",
			accountNumber: "0000002920884960",
		},
		{
			iban:               "NO9386011117947",
			countryCode:        "NO",
			checkDigit:         "93",
			bban:               "86011117947",
			replaced:           "NO0086011117947",
			bankCode:           "8601",
			accountNumber:      "111794",
			nationalCheckDigit: "7",
		},
		{
			iban:          "VA59001123000012345678",
			countryCode:   "VA",
			checkDigit:    "59",
			bban:          "001123000012345678",
			replaced:      "VA00001123000012345678",
			bankCode:      "001",
			accountNumber: "123000012345678",
		},
		{
			iban:          "MU17BOMM0101101030300200000MUR",
			countryCode:   "MU",
			checkDigit:    "17",
			bban:          "BOMM0101101030300200000MUR",
			replaced:      "MU00BOMM0101101030300200000MUR",
			bankCode:      "BOMM01",
			accountNumber: "101030300200",
			branchCode:    "01",
			currency:      "MUR",
		},
	}
	invalidCases = []struct {
		iban string
		err  error
	}{
		{
			iban: "",
			err:  ErrIbanTooShort,
		},
		{
			iban: "AL4721211",
			err:  ErrIbanTooShort,
		},
		{
			iban: "sL472121142342323",
			err:  ErrCountryCodeNotUpper,
		},
		{
			iban: "S4472121142342323",
			err:  ErrCountryCodeNotAlpha,
		},
		{
			iban: "XX472121142342323",
			err:  ErrCountryCodeNotPresent,
		},
		{
			iban: "CH9300762011-623852957",
			err:  ErrInvalidBbanLength,
		},
		{
			iban: "SK0X2121142342323",
			err:  ErrInvalidBbanLength,
		},
		{
			iban: "PL67102010260000042270201111",
			err:  ErrInvalidCheckDigit,
		},
		{
			iban: "AT61190430023457320",
			err:  ErrInvalidBbanLength,
		},
		{
			iban: "SK061100A000002920884960",
			err:  ErrInvalidBbanPart,
		},
	}
	invalidBbanCases = []struct {
		iban        string
		countryCode string
		err         error
	}{
		{
			iban:        "SA03800000006080101675194",
			countryCode: "SA",
			err:         ErrInvalidBbanLength,
		},
		{
			iban:        "GB29NWB161331926819",
			countryCode: "GB",
			err:         ErrInvalidBbanLength,
		},
		{
			iban:        "BE68507547034",
			countryCode: "BE",
			err:         ErrInvalidBbanLength,
		},
		{
			iban:        "SK0611000002920884960",
			countryCode: "SK",
			err:         ErrInvalidBbanLength,
		},
	}
)

func TestValidateMinLength(t *testing.T) {
	for _, cs := range validCases {
		t.Run(cs.iban, func(t *testing.T) {
			err := validateMinLength(cs.iban)
			require.NoError(t, err)
		})
	}
}

func TestExtractCountryCode(t *testing.T) {
	for _, cs := range validCases {
		t.Run(cs.iban, func(t *testing.T) {
			code := extractCountryCode(cs.iban)
			require.Equal(t, cs.countryCode, code)
		})
	}
}

func TestExtractDigit(t *testing.T) {
	for _, cs := range validCases {
		t.Run(cs.iban, func(t *testing.T) {
			digit := extractCheckDigit(cs.iban)
			require.Equal(t, cs.checkDigit, digit)
		})
	}
}

func TestExtractBban(t *testing.T) {
	for _, cs := range validCases {
		t.Run(cs.iban, func(t *testing.T) {
			bbn := extractBban(cs.iban)
			require.Equal(t, cs.bban, bbn)
		})
	}
}

func TestCountryCodeExists(t *testing.T) {
	for _, cs := range validCases {
		t.Run(cs.iban, func(t *testing.T) {
			code := extractCountryCode(cs.iban)
			require.True(t, country.Exists(code))
		})
	}
}

func TestReplaceDigit(t *testing.T) {
	for _, cs := range validCases {
		t.Run(cs.iban, func(t *testing.T) {
			code := extractCountryCode(cs.iban)
			replaced := replaceCheckDigit(cs.iban, code)
			require.Equal(t, cs.replaced, replaced)
		})
	}
}

func TestValidateDigit(t *testing.T) {
	for _, cs := range validCases {
		t.Run(cs.iban, func(t *testing.T) {
			code := extractCountryCode(cs.iban)
			err := validateCheckDigit(cs.iban, code)
			require.NoError(t, err)
		})
	}
}

func TestValidateCountryCode(t *testing.T) {
	for _, cs := range validCases {
		t.Run(cs.iban, func(t *testing.T) {
			code, err := validateCountryCode(cs.iban)
			require.NoError(t, err)
			require.Equal(t, cs.countryCode, code)
		})
	}
}

func TestCalculateMod(t *testing.T) {
	for _, cs := range validCases {
		t.Run(cs.iban, func(t *testing.T) {
			mod, err := calculateMod(cs.iban)
			require.NoError(t, err)
			require.Equal(t, int64(1), mod)
		})
	}
}

func TestValidateBbanLength(t *testing.T) {
	for _, cs := range validCases {
		t.Run(cs.iban, func(t *testing.T) {
			structure, ok := country.GetBbanStructure(cs.countryCode)
			require.True(t, ok)
			err := validateBbanLength(cs.iban, structure)
			require.NoError(t, err)
		})
	}
}

func TestValidateBbanLengthInvalid(t *testing.T) {
	for _, cs := range invalidBbanCases {
		t.Run(cs.iban, func(t *testing.T) {
			structure, ok := country.GetBbanStructure(cs.countryCode)
			require.True(t, ok)
			err := validateBbanLength(cs.iban, structure)
			require.Equal(t, ErrInvalidBbanLength, err)
		})
	}
}

func TestValidateBbanStructure(t *testing.T) {
	for _, cs := range validCases {
		t.Run(cs.iban, func(t *testing.T) {
			structure, ok := country.GetBbanStructure(cs.countryCode)
			require.True(t, ok)
			err := validateBbanStructure(cs.iban, structure)
			require.NoError(t, err)
		})
	}
}

func TestNew(t *testing.T) {
	for _, cs := range validCases {
		t.Run(cs.iban, func(t *testing.T) {
			ib, err := New(cs.iban)
			require.NoError(t, err)
			require.Equal(t, cs.bankCode, ib.BankCode())
			require.Equal(t, cs.branchCode, ib.BranchCode())
			require.Equal(t, cs.accountNumber, ib.AccountNumber())
			require.Equal(t, cs.identificationNumber, ib.IdentificationNumber())
			require.Equal(t, cs.accountType, ib.AccountType())
			require.Equal(t, cs.ownerAccountType, ib.OwnerAccountType())
			require.Equal(t, cs.nationalCheckDigit, ib.NationalCheckDigit())
			require.Equal(t, cs.checkDigit, ib.CheckDigit())
			require.Equal(t, cs.countryCode, ib.CountryCode())
			require.Equal(t, cs.bban, ib.Bban())
			require.Equal(t, cs.currency, ib.Currency())
			require.Equal(t, cs.iban, ib.String())
		})
	}
}

func TestNewInvalid(t *testing.T) {
	for _, cs := range invalidCases {
		t.Run(cs.iban, func(t *testing.T) {
			ib, err := New(cs.iban)
			require.Nil(t, ib)
			require.Error(t, err)
		})
	}
}

func TestMustParse(t *testing.T) {
	for _, cs := range validCases {
		t.Run(cs.iban, func(t *testing.T) {
			require.NotPanics(t, func() {
				ib := MustParse(cs.iban)
				require.NotNil(t, ib)
			})
		})
	}
}

func TestMustParseInvalid(t *testing.T) {
	for _, cs := range invalidCases {
		t.Run(cs.iban, func(t *testing.T) {
			require.Panics(t, func() {
				MustParse(cs.iban)
			})
		})
	}
}
