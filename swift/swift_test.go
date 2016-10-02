package swift

import (
	"testing"
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
			if err := validateLength(cs.swift); err != nil {
				t.Errorf("unexpected error %v", err)
			}
		})
	}
}

func TestValidateCase(t *testing.T) {
	for _, cs := range validCases {
		t.Run(cs.swift, func(t *testing.T) {
			if err := validateCase(cs.swift); err != nil {
				t.Errorf("unexpected error %v", err)
			}
		})
	}
}

func TestValidateBankCode(t *testing.T) {
	for _, cs := range validCases {
		t.Run(cs.swift, func(t *testing.T) {
			if err := validateBankCode(cs.swift); err != nil {
				t.Errorf("unexpected error %v", err)
			}
		})
	}
}

func TestValidateCountryCode(t *testing.T) {
	for _, cs := range validCases {
		t.Run(cs.swift, func(t *testing.T) {
			if err := validateCountryCode(cs.swift); err != nil {
				t.Errorf("unexpected error %v", err)
			}
		})
	}
}

func TestValidateLocationCode(t *testing.T) {
	for _, cs := range validCases {
		t.Run(cs.swift, func(t *testing.T) {
			if err := validateLocationCode(cs.swift); err != nil {
				t.Errorf("unexpected error %v", err)
			}
		})
	}
}

func TestValidateBranchCode(t *testing.T) {
	for _, cs := range validCases {
		t.Run(cs.swift, func(t *testing.T) {
			if err := validateBranchCode(cs.swift); err != nil {
				t.Errorf("unexpected error %v", err)
			}
		})
	}
}

func TestNew(t *testing.T) {
	for _, cs := range validCases {
		t.Run(cs.swift, func(t *testing.T) {
			sw, err := New(cs.swift)
			if err != nil {
				t.Errorf("unexpected error %v", err)
			}
			if cs.bankCode != sw.BankCode() {
				t.Errorf("expected %v got %v", cs.bankCode, sw.BankCode())
			}
			if cs.countryCode != sw.CountryCode() {
				t.Errorf("expected %v got %v", cs.countryCode, sw.CountryCode())
			}
			if cs.locationCode != sw.LocationCode() {
				t.Errorf("expected %v got %v", cs.locationCode, sw.LocationCode())
			}
			if cs.branchCode != sw.BranchCode() {
				t.Errorf("expected %v got %v", cs.branchCode, sw.BranchCode())
			}
			if cs.typ != sw.Type() {
				t.Errorf("expected %v got %v", cs.typ, sw.Type())
			}
		})
	}
}

func TestNewInvalid(t *testing.T) {
	for _, cs := range invalidCases {
		t.Run(cs.swift, func(t *testing.T) {
			sw, err := New(cs.swift)
			if sw != nil {
				t.Errorf("expected nil got %v", sw)
			}
			if cs.err != err {
				t.Errorf("expected %v got %v", cs.err, err)
			}
		})
	}
}

func TestMustParse(t *testing.T) {
	for _, cs := range validCases {
		t.Run(cs.swift, func(t *testing.T) {
			defer func() {
				if r := recover(); r != nil {
					t.Errorf("unexpected panic")
				}
			}()

			if sw := MustParse(cs.swift); sw == nil {
				t.Error("unexpected nil")
			}
		})
	}
}

func TestMustParseInvalid(t *testing.T) {
	for _, cs := range invalidCases {
		t.Run(cs.swift, func(t *testing.T) {
			defer func() {
				if r := recover(); r == nil {
					t.Errorf("expected panic")
				}
			}()

			if sw := MustParse(cs.swift); sw != nil {
				t.Error("expected nil")
			}
		})
	}
}
