package bban

import (
	"testing"
)

var (
	charTypeTests = []struct {
		in       string
		charType charType
		want     bool
	}{
		{"0123", Num, true},
		{"AB23", Num, false},
		{"AB", Num, false},
		{"", Num, false},
		{"DSA", AlphaUpper, true},
		{"dsa", AlphaUpper, false},
		{"32", AlphaUpper, false},
		{"", AlphaUpper, false},
		{"AB2", AlphaNum, true},
		{"AB", AlphaNum, true},
		{"", AlphaNum, false},
	}
	partTests = []struct {
		length    int
		entryType EntryType
		charType  charType
		val       string
		want      bool
	}{
		{4, BankCode, Num, "213", true},
		{5, BranchCode, AlphaNum, "213", true},
		{8, AccountNumber, AlphaUpper, "213", false},
		{4, NationalCheckDigit, Num, "213", true},
		{3, AccountType, AlphaNum, "213", true},
		{1, OwnerAccountType, AlphaUpper, "ABCD", true},
		{2, IdentificationNumber, Num, "213", true},
	}
	newPartTests = []struct {
		new  func(length int, char charType) Part
		want EntryType
	}{
		{NewBankCode, BankCode},
		{NewBranchCode, BranchCode},
		{NewAccountNumber, AccountNumber},
		{NewNationalCheckDigit, NationalCheckDigit},
		{NewAccountType, AccountType},
		{NewOwnerAccountType, OwnerAccountType},
		{NewIdentificationNumber, IdentificationNumber},
	}
)

func TestCharTypeValidate(t *testing.T) {
	for _, tc := range charTypeTests {
		t.Run(tc.in, func(t *testing.T) {
			result := tc.charType.Validate(tc.in)
			if tc.want != result {
				t.Errorf("expected %v got %v", tc.want, result)
			}
		})
	}
}

func TestPartValidate(t *testing.T) {
	for _, tc := range partTests {
		t.Run(tc.val, func(t *testing.T) {
			part := NewPart(tc.length, tc.charType, tc.entryType)
			result := part.Validate(tc.val)
			if tc.want != result {
				t.Errorf("expected %v got %v", tc.want, result)
			}
		})
	}
}

func TestNewPart(t *testing.T) {
	for _, tc := range newPartTests {
		t.Run(tc.want.String(), func(t *testing.T) {
			part := tc.new(4, Num)
			if tc.want != part.EntryType {
				t.Errorf("expected %v got %v", tc.want, part.EntryType)
			}
			if part.String() != part.EntryType.String() {
				t.Errorf("expected %v got %v", part.String(), part.EntryType.String())
			}
		})
	}
}

func TestPartString(t *testing.T) {
	for _, tc := range newPartTests {
		t.Run(tc.want.String(), func(t *testing.T) {
			part := tc.new(3, AlphaNum)
			if part.String() != part.EntryType.String() {
				t.Errorf("expected %v got %v", part.String(), part.EntryType.String())
			}
		})
	}
}

func TestStructureLength(t *testing.T) {
	st1 := NewStructure()
	if 0 != st1.Length() {
		t.Errorf("expected 0 got %v", st1.Length())
	}

	st2 := NewStructure(Part{Length: 2}, Part{Length: 4})
	if 6 != st2.Length() {
		t.Errorf("expected 0 got %v", st2.Length())
	}
}
