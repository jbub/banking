package swift

import (
	"testing"

	"github.com/stretchr/testify/require"
)

var (
	validCases = []struct {
		swift        string
		bankCode     string
		countryCode  string
		locationCode string
		branchCode   string
		typ          Type
	}{
		{
			swift:        "TATRSKBX",
			bankCode:     "TATR",
			countryCode:  "SK",
			locationCode: "BX",
			branchCode:   "",
			typ:          Type8,
		},
		{
			swift:        "GIBASKBX",
			bankCode:     "GIBA",
			countryCode:  "SK",
			locationCode: "BX",
			branchCode:   "",
			typ:          Type8,
		},
		{
			swift:        "DEUTDEFF500",
			bankCode:     "DEUT",
			countryCode:  "DE",
			locationCode: "FF",
			branchCode:   "500",
			typ:          Type11,
		},
	}
	invalidCases = []struct {
		swift string
		err   error
	}{
		{
			swift: "",
			err:   ErrInvalidLength,
		},
		{
			swift: "KU78N78",
			err:   ErrInvalidLength,
		},
		{
			swift: "KU78N78K43KL",
			err:   ErrInvalidLength,
		},
		{
			swift: "MK23MjK2",
			err:   ErrInvalidCase,
		},
		{
			swift: "MK23MjK2D23",
			err:   ErrInvalidCase,
		},
		{
			swift: "MK23KDLF",
			err:   ErrInvalidBankCode,
		},
		{
			swift: "24KM3KDLFDS",
			err:   ErrInvalidBankCode,
		},
		{
			swift: "JMKM3KDL",
			err:   ErrInvalidCountryCode,
		},
		{
			swift: "JMKM3KDLFDS",
			err:   ErrInvalidCountryCode,
		},
		{
			swift: "JMKMXXDL",
			err:   ErrCountryCodeNotPresent,
		},
		{
			swift: "JMKMXXDLFDS",
			err:   ErrCountryCodeNotPresent,
		},
		{
			swift: "JMKMSK--",
			err:   ErrInvalidLocationCode,
		},
		{
			swift: "JMKMSK--DSL",
			err:   ErrInvalidLocationCode,
		},
		{
			swift: "JMKMSKLDDS-",
			err:   ErrInvalidBranchCode,
		},
	}
)

func TestValidateLength(t *testing.T) {
	for _, cs := range validCases {
		t.Run(cs.swift, func(t *testing.T) {
			err := validateLength(cs.swift)
			require.NoError(t, err)
		})
	}
}

func TestValidateCase(t *testing.T) {
	for _, cs := range validCases {
		t.Run(cs.swift, func(t *testing.T) {
			err := validateCase(cs.swift)
			require.NoError(t, err)
		})
	}
}

func TestValidateBankCode(t *testing.T) {
	for _, cs := range validCases {
		t.Run(cs.swift, func(t *testing.T) {
			err := validateBankCode(cs.swift)
			require.NoError(t, err)
		})
	}
}

func TestValidateCountryCode(t *testing.T) {
	for _, cs := range validCases {
		t.Run(cs.swift, func(t *testing.T) {
			err := validateCountryCode(cs.swift)
			require.NoError(t, err)
		})
	}
}

func TestValidateLocationCode(t *testing.T) {
	for _, cs := range validCases {
		t.Run(cs.swift, func(t *testing.T) {
			err := validateLocationCode(cs.swift)
			require.NoError(t, err)
		})
	}
}

func TestValidateBranchCode(t *testing.T) {
	for _, cs := range validCases {
		t.Run(cs.swift, func(t *testing.T) {
			err := validateBranchCode(cs.swift)
			require.NoError(t, err)
		})
	}
}

func TestParse(t *testing.T) {
	for _, cs := range validCases {
		t.Run(cs.swift, func(t *testing.T) {
			sw, err := Parse(cs.swift)
			require.NoError(t, err)
			require.Equal(t, cs.bankCode, sw.BankCode())
			require.Equal(t, cs.countryCode, sw.CountryCode())
			require.Equal(t, cs.locationCode, sw.LocationCode())
			require.Equal(t, cs.branchCode, sw.BranchCode())
			require.Equal(t, cs.typ, sw.Type())
		})
	}
}

func TestParseInvalid(t *testing.T) {
	for _, cs := range invalidCases {
		t.Run(cs.swift, func(t *testing.T) {
			sw, err := Parse(cs.swift)
			require.Nil(t, sw)
			require.Equal(t, cs.err, err)
		})
	}
}

func TestMustParse(t *testing.T) {
	for _, cs := range validCases {
		t.Run(cs.swift, func(t *testing.T) {
			require.NotPanics(t, func() {
				sw := MustParse(cs.swift)
				require.NotNil(t, sw)
			})
		})
	}
}

func TestMustParseInvalid(t *testing.T) {
	for _, cs := range invalidCases {
		t.Run(cs.swift, func(t *testing.T) {
			require.Panics(t, func() {
				MustParse(cs.swift)
			})
		})
	}
}

func BenchmarkValidate(b *testing.B) {
	b.ReportAllocs()
	b.SetBytes(2)

	for i := 0; i < b.N; i++ {
		_ = Validate("DEUTDEFF500")
	}
}
