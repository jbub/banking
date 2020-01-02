package bban

// Part represents a substring part of bban.
type Part struct {
	Length    int
	EntryType EntryType
	charType  charType
}

// Validate validates given value against part CharType.
func (p Part) Validate(value string) bool {
	return p.charType.Validate(value)
}

// String returns a text representation of Part.
func (p Part) String() string {
	return p.EntryType.String()
}

// NewPart creates a new Part.
func NewPart(length int, char charType, entry EntryType) Part {
	return Part{
		Length:    length,
		EntryType: entry,
		charType:  char,
	}
}

// NewBankCode creates a new Part with BankCode EntryType.
func NewBankCode(length int, char charType) Part {
	return NewPart(length, char, BankCode)
}

// NewBranchCode creates a new Part with BranchCode EntryType.
func NewBranchCode(length int, char charType) Part {
	return NewPart(length, char, BranchCode)
}

// NewAccountNumber creates a new Part with AccountNumber EntryType.
func NewAccountNumber(length int, char charType) Part {
	return NewPart(length, char, AccountNumber)
}

// NewNationalCheckDigit creates a new Part with NationalCheckDigit EntryType.
func NewNationalCheckDigit(length int, char charType) Part {
	return NewPart(length, char, NationalCheckDigit)
}

// NewAccountType creates a new Part with AccountType EntryType.
func NewAccountType(length int, char charType) Part {
	return NewPart(length, char, AccountType)
}

// NewOwnerAccountType creates a new Part with OwnerAccountType EntryType.
func NewOwnerAccountType(length int, char charType) Part {
	return NewPart(length, char, OwnerAccountType)
}

// NewIdentificationNumber creates a new Part with IdentificationNumber EntryType.
func NewIdentificationNumber(length int, char charType) Part {
	return NewPart(length, char, IdentificationNumber)
}

// NewCurrency creates a new Part with Currency EntryType.
func NewCurrency(length int, char charType) Part {
	return NewPart(length, char, Currency)
}

// NewPadding creates a new Part with Padding EntryType.
func NewPadding(length int, char charType) Part {
	return NewPart(length, char, Padding)
}
