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
			bankCode:           "212",
			branchCode:         "1100",
			accountNumber:      "0000000235698741",
			nationalCheckDigit: "9",
		},
		{
			iban:          "BY13NBRB3600900000002Z00AB00",
			countryCode:   "BY",
			checkDigit:    "13",
			bban:          "NBRB3600900000002Z00AB00",
			bankCode:      "NBRB",
			branchCode:    "3600",
			accountNumber: "900000002Z00AB00",
		},
		{
			iban:               "BE68539007547034",
			countryCode:        "BE",
			checkDigit:         "68",
			bban:               "539007547034",
			bankCode:           "539",
			accountNumber:      "0075470",
			nationalCheckDigit: "34",
		},
		{
			iban:               "EG800002000156789012345180002",
			countryCode:        "EG",
			checkDigit:         "80",
			bban:               "0002000156789012345180002",
			bankCode:           "0002",
			branchCode:         "0001",
			accountNumber:      "56789012345180002",
		},
		{
			iban:               "FO6264600001631634",
			countryCode:        "FO",
			checkDigit:         "62",
			bban:               "64600001631634",
			bankCode:           "6460",
			accountNumber:      "000163163",
			nationalCheckDigit: "4",
		},
		{
			iban:          "GL8964710001000206",
			countryCode:   "GL",
			checkDigit:    "89",
			bban:          "64710001000206",
			bankCode:      "6471",
			accountNumber: "0001000206",
		},
		{
			iban:          "GR1601101250000000012300695",
			countryCode:   "GR",
			checkDigit:    "16",
			bban:          "01101250000000012300695",
			bankCode:      "011",
			branchCode:    "0125",
			accountNumber: "0000000012300695",
		},
		{
			iban:          "GB29NWBK60161331926819",
			countryCode:   "GB",
			checkDigit:    "29",
			bban:          "NWBK60161331926819",
			bankCode:      "NWBK",
			branchCode:    "601613",
			accountNumber: "31926819",
		},
		{
			iban:          "IQ98NBIQ850123456789012",
			countryCode:   "IQ",
			checkDigit:    "98",
			bban:          "NBIQ850123456789012",
			bankCode:      "NBIQ",
			branchCode:    "850",
			accountNumber: "123456789012",
		},
		{
			iban:          "SA0380000000608010167519",
			countryCode:   "SA",
			checkDigit:    "03",
			bban:          "80000000608010167519",
			bankCode:      "80",
			accountNumber: "000000608010167519",
		},
		{
			iban:          "CH9300762011623852957",
			countryCode:   "CH",
			checkDigit:    "93",
			bban:          "00762011623852957",
			bankCode:      "00762",
			accountNumber: "011623852957",
		},
		{
			iban:               "TL380080012345678910157",
			countryCode:        "TL",
			checkDigit:         "38",
			bban:               "0080012345678910157",
			bankCode:           "008",
			accountNumber:      "00123456789101",
			nationalCheckDigit: "57",
		},
		{
			iban:               "TR330006100519786457841326",
			countryCode:        "TR",
			checkDigit:         "33",
			bban:               "0006100519786457841326",
			bankCode:           "00061",
			accountNumber:      "0519786457841326",
			nationalCheckDigit: "0",
		},
		{
			iban:               "PL60102010260000042270201111",
			countryCode:        "PL",
			checkDigit:         "60",
			bban:               "102010260000042270201111",
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
			bankCode:      "1100",
			accountNumber: "0000002920884960",
		},
		{
			iban:               "NO9386011117947",
			countryCode:        "NO",
			checkDigit:         "93",
			bban:               "86011117947",
			bankCode:           "8601",
			accountNumber:      "111794",
			nationalCheckDigit: "7",
		},
		{
			iban:          "VA59001123000012345678",
			countryCode:   "VA",
			checkDigit:    "59",
			bban:          "001123000012345678",
			bankCode:      "001",
			accountNumber: "123000012345678",
		},
		{
			iban:          "MU17BOMM0101101030300200000MUR",
			countryCode:   "MU",
			checkDigit:    "17",
			bban:          "BOMM0101101030300200000MUR",
			bankCode:      "BOMM01",
			accountNumber: "101030300200",
			branchCode:    "01",
			currency:      "MUR",
		},
		{
			iban:          "XK051212012345678906",
			countryCode:   "XK",
			checkDigit:    "05",
			bban:          "1212012345678906",
			bankCode:      "1212",
			accountNumber: "012345678906",
		},
		{
			iban:          "LC55HEMM000100010012001200023015",
			countryCode:   "LC",
			checkDigit:    "55",
			bban:          "HEMM000100010012001200023015",
			bankCode:      "HEMM",
			accountNumber: "000100010012001200023015",
		},
		{
			iban:          "SC18SSCB11010000000000001497USD",
			countryCode:   "SC",
			checkDigit:    "18",
			bban:          "SSCB11010000000000001497USD",
			bankCode:      "SSCB",
			branchCode:    "1101",
			accountNumber: "0000000000001497",
			currency:      "USD",
		},
		{
			iban:          "UA213996220000026007233566001",
			countryCode:   "UA",
			checkDigit:    "21",
			bban:          "3996220000026007233566001",
			bankCode:      "399622",
			accountNumber: "0000026007233566001",
		},
		{
			iban:          "CR05015202001026284066",
			countryCode:   "CR",
			checkDigit:    "05",
			bban:          "015202001026284066",
			bankCode:      "152",
			accountNumber: "02001026284066",
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
			code := extractCountryCode(cs.iban)
			err := validateCountryCode(code)
			require.NoError(t, err)
		})
	}
}

func TestCalculateCheckDigit(t *testing.T) {
	for _, cs := range validCases {
		t.Run(cs.iban, func(t *testing.T) {
			digit, err := calculateCheckDigit(cs.iban, cs.countryCode)
			require.NoError(t, err)
			require.Equal(t, cs.checkDigit, digit)
		})
	}
}

func TestValidateBbanLength(t *testing.T) {
	for _, cs := range validCases {
		t.Run(cs.iban, func(t *testing.T) {
			struc, ok := country.GetBbanStructure(cs.countryCode)
			require.True(t, ok)
			bbn := extractBban(cs.iban)
			err := validateBban(bbn, struc)
			require.NoError(t, err)
		})
	}
}

func TestValidateBbanLengthInvalid(t *testing.T) {
	for _, cs := range invalidBbanCases {
		t.Run(cs.iban, func(t *testing.T) {
			struc, ok := country.GetBbanStructure(cs.countryCode)
			require.True(t, ok)
			bbn := extractBban(cs.iban)
			err := validateBban(bbn, struc)
			require.Equal(t, ErrInvalidBbanLength, err)
		})
	}
}

func TestValidateBbanStructure(t *testing.T) {
	for _, cs := range validCases {
		t.Run(cs.iban, func(t *testing.T) {
			struc, ok := country.GetBbanStructure(cs.countryCode)
			require.True(t, ok)
			bbn := extractBban(cs.iban)
			err := validateBban(bbn, struc)
			require.NoError(t, err)
		})
	}
}

func TestParse(t *testing.T) {
	for _, cs := range validCases {
		t.Run(cs.iban, func(t *testing.T) {
			ib, err := Parse(cs.iban)
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

func TestValidateValid(t *testing.T) {
	for _, cs := range validCases {
		t.Run(cs.iban, func(t *testing.T) {
			err := Validate(cs.iban)
			require.NoError(t, err)
		})
	}
}

func TestValidateInvalid(t *testing.T) {
	for _, cs := range invalidCases {
		t.Run(cs.iban, func(t *testing.T) {
			err := Validate(cs.iban)
			require.Error(t, err)
		})
	}
}

func TestParseInvalid(t *testing.T) {
	for _, cs := range invalidCases {
		t.Run(cs.iban, func(t *testing.T) {
			ib, err := Parse(cs.iban)
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

func BenchmarkValidate(b *testing.B) {
	b.ReportAllocs()
	b.SetBytes(2)

	for i := 0; i < b.N; i++ {
		_ = Validate("AL47212110090000000235698741")
	}
}
