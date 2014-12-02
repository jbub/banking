package bban

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

var (
	charTypeTests = []struct {
		Value    string
		CharType charType
		Expected bool
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
		Length    int
		EntryType EntryType
		CharType  charType
		Value     string
		Expected  bool
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
		New      func(length int, char charType) Part
		Expected EntryType
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

type StructureTestSuite struct {
	suite.Suite
}

func (s *StructureTestSuite) TestCharTypeValidate() {
	for _, ctt := range charTypeTests {
		result := ctt.CharType.Validate(ctt.Value)
		s.Equal(ctt.Expected, result, "Value = %s", ctt.Value)
	}
}

func (s *StructureTestSuite) TestPartValidate() {
	for _, tt := range partTests {
		part := NewPart(tt.Length, tt.CharType, tt.EntryType)
		result := part.Validate(tt.Value)
		s.Equal(tt.Expected, result, "EntryType = %s, Value = %s", tt.EntryType, tt.Value)
	}
}

func (s *StructureTestSuite) TestNewPart() {
	for _, tt := range newPartTests {
		part := tt.New(4, Num)
		s.Equal(tt.Expected, part.EntryType)
		s.Equal(part.String(), part.EntryType.String())
	}
}

func (s *StructureTestSuite) TestPartString() {
	for _, tt := range newPartTests {
		part := tt.New(3, AlphaNum)
		s.Equal(part.String(), part.EntryType.String())
	}
}

func (s *StructureTestSuite) TestStructureLength() {
	st1 := NewStructure()
	s.Equal(0, st1.Length())
	st2 := NewStructure(Part{Length: 2}, Part{Length: 4})
	s.Equal(6, st2.Length())
}

func TestStructureTestSuite(t *testing.T) {
	suite.Run(t, new(StructureTestSuite))
}
