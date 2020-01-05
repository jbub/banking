package bban

import (
	"testing"

	"github.com/stretchr/testify/require"
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
		{"", Zero, false},
		{"0", Zero, true},
		{"000", Zero, true},
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
		{3, Currency, AlphaUpper, "MUR", true},
		{3, Padding, AlphaNum, "00", true},
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
		{NewCurrency, Currency},
		{NewPadding, Padding},
	}
)

func TestCharTypeValidate(t *testing.T) {
	for _, tc := range charTypeTests {
		t.Run(tc.in, func(t *testing.T) {
			result := tc.charType.Validate(tc.in)
			require.Equal(t, tc.want, result)
		})
	}
}

func TestPartValidate(t *testing.T) {
	for _, tc := range partTests {
		t.Run(tc.val, func(t *testing.T) {
			part := NewPart(tc.length, tc.charType, tc.entryType)
			result := part.Validate(tc.val)
			require.Equal(t, tc.want, result)
		})
	}
}

func TestNewPart(t *testing.T) {
	for _, tc := range newPartTests {
		t.Run(tc.want.String(), func(t *testing.T) {
			part := tc.new(4, Num)
			require.Equal(t, tc.want, part.EntryType)
			require.Equal(t, part.String(), part.EntryType.String())
		})
	}
}

func TestPartString(t *testing.T) {
	for _, tc := range newPartTests {
		t.Run(tc.want.String(), func(t *testing.T) {
			part := tc.new(3, AlphaNum)
			require.Equal(t, part.String(), part.EntryType.String())
		})
	}
}

func TestStructureLength(t *testing.T) {
	st1 := NewStructure()
	require.Equal(t, 0, st1.Length())

	st2 := NewStructure(Part{Length: 2}, Part{Length: 4})
	require.Equal(t, 6, st2.Length())
}
