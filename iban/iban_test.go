package iban

import (
	"testing"

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
			if err := validateMinLength(cs.iban); err != nil {
				t.Errorf("unexpected error %v", err)
			}
		})
	}
}

func TestExtractCountryCode(t *testing.T) {
	for _, cs := range validCases {
		t.Run(cs.iban, func(t *testing.T) {
			code := extractCountryCode(cs.iban)
			if cs.countryCode != code {
				t.Errorf("expected %v got %v", cs.countryCode, code)
			}
		})
	}
}

func TestExtractDigit(t *testing.T) {
	for _, cs := range validCases {
		t.Run(cs.iban, func(t *testing.T) {
			digit := extractCheckDigit(cs.iban)
			if cs.checkDigit != digit {
				t.Errorf("expected %v got %v", cs.checkDigit, digit)
			}
		})
	}
}

func TestExtractBban(t *testing.T) {
	for _, cs := range validCases {
		t.Run(cs.iban, func(t *testing.T) {
			bban := extractBban(cs.iban)
			if cs.bban != bban {
				t.Errorf("expected %v got %v", cs.bban, bban)
			}
		})
	}
}

func TestCountryCodeExists(t *testing.T) {
	for _, cs := range validCases {
		t.Run(cs.iban, func(t *testing.T) {
			code := extractCountryCode(cs.iban)
			exists := country.Exists(code)
			if !exists {
				t.Errorf("expected code %v to exist", code)
			}
		})
	}
}

func TestReplaceDigit(t *testing.T) {
	for _, cs := range validCases {
		t.Run(cs.iban, func(t *testing.T) {
			code := extractCountryCode(cs.iban)
			replaced := replaceCheckDigit(cs.iban, code)
			if cs.replaced != replaced {
				t.Errorf("expected %v got %v", cs.replaced, replaced)
			}
		})
	}
}

func TestValidateDigit(t *testing.T) {
	for _, cs := range validCases {
		t.Run(cs.iban, func(t *testing.T) {
			code := extractCountryCode(cs.iban)
			if err := validateCheckDigit(cs.iban, code); err != nil {
				t.Errorf("unexpected error %v", err)
			}
		})
	}
}

func TestValidateCountryCode(t *testing.T) {
	for _, cs := range validCases {
		t.Run(cs.iban, func(t *testing.T) {
			code, err := validateCountryCode(cs.iban)
			if err != nil {
				t.Errorf("unexpected error %v", err)
			}
			if cs.countryCode != code {
				t.Errorf("expected %v got %v", cs.countryCode, code)
			}
		})
	}
}

func TestCalculateMod(t *testing.T) {
	for _, cs := range validCases {
		t.Run(cs.iban, func(t *testing.T) {
			mod, err := calculateMod(cs.iban)
			if err != nil {
				t.Errorf("unexpected error %v", err)
			}
			want := int64(1)
			if mod != want {
				t.Errorf("expected %v got %v", want, mod)
			}
		})
	}
}

func TestValidateBbanLength(t *testing.T) {
	for _, cs := range validCases {
		t.Run(cs.iban, func(t *testing.T) {
			structure, ok := country.GetBbanStructure(cs.countryCode)
			if !ok {
				t.Errorf("expected structure for code %v to exist", cs.countryCode)
			}
			if err := validateBbanLength(cs.iban, structure); err != nil {
				t.Errorf("unexpected error %v", err)
			}
		})
	}
}

func TestValidateBbanLengthInvalid(t *testing.T) {
	for _, cs := range invalidBbanCases {
		t.Run(cs.iban, func(t *testing.T) {
			structure, ok := country.GetBbanStructure(cs.countryCode)
			if !ok {
				t.Errorf("expected structure for code %v to exist", cs.countryCode)
			}
			if err := validateBbanLength(cs.iban, structure); err != ErrInvalidBbanLength {
				t.Errorf("expected error, got %v", err)
			}
		})
	}
}

func TestValidateBbanStructure(t *testing.T) {
	for _, cs := range validCases {
		t.Run(cs.iban, func(t *testing.T) {
			structure, ok := country.GetBbanStructure(cs.countryCode)
			if !ok {
				t.Errorf("expected structure for code %v to exist", cs.countryCode)
			}
			if err := validateBbanStructure(cs.iban, structure); err != nil {
				t.Errorf("unexpected error %v", err)
			}
		})
	}
}

func TestNew(t *testing.T) {
	for _, cs := range validCases {
		t.Run(cs.iban, func(t *testing.T) {
			ib, err := New(cs.iban)
			if err != nil {
				t.Errorf("unexpected error %v", err)
			}
			if cs.bankCode != ib.BankCode() {
				t.Errorf("expected %v got %v", cs.bankCode, ib.BankCode())
			}
			if cs.branchCode != ib.BranchCode() {
				t.Errorf("expected %v got %v", cs.branchCode, ib.BranchCode())
			}
			if cs.accountNumber != ib.AccountNumber() {
				t.Errorf("expected %v got %v", cs.accountNumber, ib.AccountNumber())
			}
			if cs.identificationNumber != ib.IdentificationNumber() {
				t.Errorf("expected %v got %v", cs.identificationNumber, ib.IdentificationNumber())
			}
			if cs.accountType != ib.AccountType() {
				t.Errorf("expected %v got %v", cs.accountType, ib.AccountType())
			}
			if cs.ownerAccountType != ib.OwnerAccountType() {
				t.Errorf("expected %v got %v", cs.ownerAccountType, ib.OwnerAccountType())
			}
			if cs.nationalCheckDigit != ib.NationalCheckDigit() {
				t.Errorf("expected %v got %v", cs.nationalCheckDigit, ib.NationalCheckDigit())
			}
			if cs.checkDigit != ib.CheckDigit() {
				t.Errorf("expected %v got %v", cs.checkDigit, ib.CheckDigit())
			}
			if cs.countryCode != ib.CountryCode() {
				t.Errorf("expected %v got %v", cs.countryCode, ib.CountryCode())
			}
			if cs.bban != ib.Bban() {
				t.Errorf("expected %v got %v", cs.bban, ib.Bban())
			}
		})
	}
}

func TestNewInvalid(t *testing.T) {
	for _, cs := range invalidCases {
		t.Run(cs.iban, func(t *testing.T) {
			ib, err := New(cs.iban)
			if ib != nil {
				t.Errorf("expected nil got %v", ib)
			}
			if cs.err != err {
				t.Errorf("expected %v got %v", cs.err, err)
			}
		})
	}
}

func TestMustParse(t *testing.T) {
	for _, cs := range validCases {
		t.Run(cs.iban, func(t *testing.T) {
			defer func() {
				if r := recover(); r != nil {
					t.Errorf("unexpected panic")
				}
			}()

			if ib := MustParse(cs.iban); ib == nil {
				t.Error("unexpected nil")
			}
		})
	}
}

func TestMustParseInvalid(t *testing.T) {
	for _, cs := range invalidCases {
		t.Run(cs.iban, func(t *testing.T) {
			defer func() {
				if r := recover(); r == nil {
					t.Errorf("expected panic")
				}
			}()

			if ib := MustParse(cs.iban); ib != nil {
				t.Error("expected nil")
			}
		})
	}
}
