package country

import (
	"github.com/jbub/banking/bban"
)

// Country holds country related banking info.
type Country struct {
	Name       string
	Alpha2Code string
	Alpha3Code string
	Structure  bban.Structure
}

// String returns text representation of country.
func (c Country) String() string {
	return c.Name
}

var (
	countries = map[string]Country{
		"AL": {
			Name:       "Albania",
			Alpha2Code: "AL",
			Alpha3Code: "ALB",
			Structure: bban.NewStructure(
				bban.NewBankCode(3, bban.Num),
				bban.NewBranchCode(4, bban.Num),
				bban.NewNationalCheckDigit(1, bban.Num),
				bban.NewAccountNumber(16, bban.AlphaNum),
			),
		},
		"AD": {
			Name:       "Andorra",
			Alpha2Code: "AD",
			Alpha3Code: "AND",
			Structure: bban.NewStructure(
				bban.NewBankCode(4, bban.Num),
				bban.NewBranchCode(4, bban.Num),
				bban.NewAccountNumber(12, bban.AlphaNum),
			),
		},
		"AT": {
			Name:       "Austria",
			Alpha2Code: "AT",
			Alpha3Code: "AUT",
			Structure: bban.NewStructure(
				bban.NewBankCode(5, bban.Num),
				bban.NewAccountNumber(11, bban.Num),
			),
		},
		"AZ": {
			Name:       "Azerbaijan",
			Alpha2Code: "AZ",
			Alpha3Code: "AZE",
			Structure: bban.NewStructure(
				bban.NewBankCode(4, bban.AlphaUpper),
				bban.NewAccountNumber(20, bban.AlphaNum),
			),
		},
		"BH": {
			Name:       "Bahrain",
			Alpha2Code: "BH",
			Alpha3Code: "BHR",
			Structure: bban.NewStructure(
				bban.NewBankCode(4, bban.AlphaUpper),
				bban.NewAccountNumber(14, bban.Num),
			),
		},
		"BE": {
			Name:       "Belgium",
			Alpha2Code: "BE",
			Alpha3Code: "BEL",
			Structure: bban.NewStructure(
				bban.NewBankCode(3, bban.Num),
				bban.NewAccountNumber(7, bban.Num),
				bban.NewNationalCheckDigit(2, bban.Num),
			),
		},
		"BA": {
			Name:       "Bosnia and Herzegovina",
			Alpha2Code: "BA",
			Alpha3Code: "BIH",
			Structure: bban.NewStructure(
				bban.NewBankCode(3, bban.Num),
				bban.NewBranchCode(3, bban.Num),
				bban.NewAccountNumber(8, bban.Num),
				bban.NewNationalCheckDigit(2, bban.Num),
			),
		},
		"BR": {
			Name:       "Brazil",
			Alpha2Code: "BR",
			Alpha3Code: "BRA",
			Structure: bban.NewStructure(
				bban.NewBankCode(8, bban.Num),
				bban.NewBranchCode(5, bban.Num),
				bban.NewAccountNumber(10, bban.Num),
				bban.NewAccountType(1, bban.AlphaUpper),
				bban.NewOwnerAccountType(1, bban.AlphaNum),
			),
		},
		"VG": {
			Name:       "British Virgin Islands",
			Alpha2Code: "VG",
			Alpha3Code: "VGB",
			Structure: bban.NewStructure(
				bban.NewBankCode(4, bban.AlphaNum),
				bban.NewAccountNumber(16, bban.Num),
			),
		},
		"BG": {
			Name:       "Bulgaria",
			Alpha2Code: "BG",
			Alpha3Code: "BGR",
			Structure: bban.NewStructure(
				bban.NewBankCode(4, bban.AlphaNum),
				bban.NewBranchCode(4, bban.Num),
				bban.NewAccountType(2, bban.Num),
				bban.NewAccountNumber(8, bban.AlphaNum),
			),
		},
		"CR": {
			Name:       "Costa Rica",
			Alpha2Code: "CR",
			Alpha3Code: "CRI",
			Structure: bban.NewStructure(
				bban.NewBankCode(3, bban.Num),
				bban.NewAccountNumber(14, bban.Num),
			),
		},
		"HR": {
			Name:       "Croatia",
			Alpha2Code: "HR",
			Alpha3Code: "HRV",
			Structure: bban.NewStructure(
				bban.NewBankCode(7, bban.Num),
				bban.NewAccountNumber(10, bban.Num),
			),
		},
		"CY": {
			Name:       "Cyprus",
			Alpha2Code: "CY",
			Alpha3Code: "CYP",
			Structure: bban.NewStructure(
				bban.NewBankCode(3, bban.Num),
				bban.NewBranchCode(5, bban.Num),
				bban.NewAccountNumber(16, bban.AlphaNum),
			),
		},
		"CN": {
			Name:       "China",
			Alpha2Code: "CN",
			Alpha3Code: "CHN",
			Structure: bban.NewStructure(
				bban.NewBankCode(4, bban.Num),
				bban.NewAccountNumber(22, bban.Num),
			),
		},
		"CZ": {
			Name:       "Czech Republic",
			Alpha2Code: "CZ",
			Alpha3Code: "CZE",
			Structure: bban.NewStructure(
				bban.NewBankCode(4, bban.Num),
				bban.NewAccountNumber(16, bban.Num),
			),
		},
		"DK": {
			Name:       "Denmark",
			Alpha2Code: "DK",
			Alpha3Code: "DNK",
			Structure: bban.NewStructure(
				bban.NewBankCode(4, bban.Num),
				bban.NewAccountNumber(10, bban.Num),
			),
		},
		"DO": {
			Name:       "Dominican Republic",
			Alpha2Code: "DO",
			Alpha3Code: "DOM",
			Structure: bban.NewStructure(
				bban.NewBankCode(4, bban.AlphaNum),
				bban.NewAccountNumber(20, bban.Num),
			),
		},
		"EE": {
			Name:       "Estonia",
			Alpha2Code: "EE",
			Alpha3Code: "EST",
			Structure: bban.NewStructure(
				bban.NewBankCode(2, bban.Num),
				bban.NewBranchCode(2, bban.Num),
				bban.NewAccountNumber(11, bban.Num),
				bban.NewNationalCheckDigit(1, bban.Num),
			),
		},
		"FI": {
			Name:       "Finland",
			Alpha2Code: "FI",
			Alpha3Code: "FIN",
			Structure: bban.NewStructure(
				bban.NewBankCode(6, bban.Num),
				bban.NewAccountNumber(7, bban.Num),
				bban.NewNationalCheckDigit(1, bban.Num),
			),
		},
		"FR": {
			Name:       "France",
			Alpha2Code: "FR",
			Alpha3Code: "FRA",
			Structure: bban.NewStructure(
				bban.NewBankCode(5, bban.Num),
				bban.NewBranchCode(5, bban.Num),
				bban.NewAccountNumber(11, bban.AlphaNum),
				bban.NewNationalCheckDigit(2, bban.Num),
			),
		},
		"GE": {
			Name:       "Georgia",
			Alpha2Code: "GE",
			Alpha3Code: "GEO",
			Structure: bban.NewStructure(
				bban.NewBankCode(2, bban.AlphaUpper),
				bban.NewAccountNumber(16, bban.Num),
			),
		},
		"DE": {
			Name:       "Germany",
			Alpha2Code: "DE",
			Alpha3Code: "DEU",
			Structure: bban.NewStructure(
				bban.NewBankCode(8, bban.Num),
				bban.NewAccountNumber(10, bban.Num),
			),
		},
		"GI": {
			Name:       "Gibraltar",
			Alpha2Code: "GI",
			Alpha3Code: "GIB",
			Structure: bban.NewStructure(
				bban.NewBankCode(4, bban.AlphaNum),
				bban.NewAccountNumber(15, bban.AlphaNum),
			),
		},
		"GR": {
			Name:       "Greece",
			Alpha2Code: "GR",
			Alpha3Code: "GRC",
			Structure: bban.NewStructure(
				bban.NewBankCode(3, bban.Num),
				bban.NewBranchCode(4, bban.Num),
				bban.NewAccountNumber(16, bban.AlphaNum),
			),
		},
		"GT": {
			Name:       "Guatemala",
			Alpha2Code: "GT",
			Alpha3Code: "GTM",
			Structure: bban.NewStructure(
				bban.NewBankCode(4, bban.AlphaNum),
				bban.NewAccountNumber(20, bban.AlphaNum),
			),
		},
		"HU": {
			Name:       "Hungary",
			Alpha2Code: "HU",
			Alpha3Code: "HUN",
			Structure: bban.NewStructure(
				bban.NewBankCode(3, bban.Num),
				bban.NewBranchCode(4, bban.Num),
				bban.NewAccountNumber(16, bban.Num),
				bban.NewNationalCheckDigit(1, bban.Num),
			),
		},
		"IS": {
			Name:       "Iceland",
			Alpha2Code: "IS",
			Alpha3Code: "ISL",
			Structure: bban.NewStructure(
				bban.NewBankCode(4, bban.Num),
				bban.NewBranchCode(2, bban.Num),
				bban.NewAccountNumber(6, bban.Num),
				bban.NewIdentificationNumber(10, bban.Num),
			),
		},
		"IE": {
			Name:       "Ireland",
			Alpha2Code: "IE",
			Alpha3Code: "IRL",
			Structure: bban.NewStructure(
				bban.NewBankCode(4, bban.AlphaUpper),
				bban.NewBranchCode(6, bban.Num),
				bban.NewAccountNumber(8, bban.Num),
			),
		},
		"IL": {
			Name:       "Israel",
			Alpha2Code: "IL",
			Alpha3Code: "ISR",
			Structure: bban.NewStructure(
				bban.NewBankCode(3, bban.Num),
				bban.NewBranchCode(3, bban.Num),
				bban.NewAccountNumber(13, bban.Num),
			),
		},
		"IT": {
			Name:       "Italy",
			Alpha2Code: "IT",
			Alpha3Code: "ITA",
			Structure: bban.NewStructure(
				bban.NewNationalCheckDigit(1, bban.AlphaUpper),
				bban.NewBankCode(5, bban.Num),
				bban.NewBranchCode(5, bban.Num),
				bban.NewAccountNumber(12, bban.AlphaNum),
			),
		},
		"JO": {
			Name:       "Jordan",
			Alpha2Code: "JO",
			Alpha3Code: "JOR",
			Structure: bban.NewStructure(
				bban.NewBankCode(4, bban.AlphaUpper),
				bban.NewBranchCode(4, bban.Num),
				bban.NewAccountNumber(18, bban.AlphaNum),
			),
		},
		"KZ": {
			Name:       "Kazakhstan",
			Alpha2Code: "KZ",
			Alpha3Code: "KAZ",
			Structure: bban.NewStructure(
				bban.NewBankCode(3, bban.Num),
				bban.NewAccountNumber(13, bban.AlphaNum),
			),
		},
		"KW": {
			Name:       "Kuwait",
			Alpha2Code: "KW",
			Alpha3Code: "KWT",
			Structure: bban.NewStructure(
				bban.NewBankCode(4, bban.AlphaUpper),
				bban.NewAccountNumber(22, bban.AlphaNum),
			),
		},
		"LV": {
			Name:       "Latvia",
			Alpha2Code: "LV",
			Alpha3Code: "LVA",
			Structure: bban.NewStructure(
				bban.NewBankCode(4, bban.AlphaUpper),
				bban.NewAccountNumber(13, bban.AlphaNum),
			),
		},
		"LB": {
			Name:       "Lebanon",
			Alpha2Code: "LB",
			Alpha3Code: "LBN",
			Structure: bban.NewStructure(
				bban.NewBankCode(4, bban.Num),
				bban.NewAccountNumber(20, bban.AlphaNum),
			),
		},
		"LI": {
			Name:       "Liechtenstein",
			Alpha2Code: "LI",
			Alpha3Code: "LIE",
			Structure: bban.NewStructure(
				bban.NewBankCode(5, bban.Num),
				bban.NewAccountNumber(12, bban.AlphaNum),
			),
		},
		"LT": {
			Name:       "Lithuania",
			Alpha2Code: "LT",
			Alpha3Code: "LTU",
			Structure: bban.NewStructure(
				bban.NewBankCode(5, bban.Num),
				bban.NewAccountNumber(11, bban.Num),
			),
		},
		"LU": {
			Name:       "Luxembourg",
			Alpha2Code: "LU",
			Alpha3Code: "LUX",
			Structure: bban.NewStructure(
				bban.NewBankCode(3, bban.Num),
				bban.NewAccountNumber(13, bban.AlphaNum),
			),
		},
		"MK": {
			Name:       "Macedonia",
			Alpha2Code: "MK",
			Alpha3Code: "MKD",
			Structure: bban.NewStructure(
				bban.NewBankCode(3, bban.Num),
				bban.NewAccountNumber(10, bban.AlphaNum),
				bban.NewNationalCheckDigit(2, bban.Num),
			),
		},
		"MT": {
			Name:       "Malta",
			Alpha2Code: "MT",
			Alpha3Code: "MLT",
			Structure: bban.NewStructure(
				bban.NewBankCode(4, bban.AlphaUpper),
				bban.NewBranchCode(5, bban.Num),
				bban.NewAccountNumber(18, bban.AlphaNum),
			),
		},
		"MR": {
			Name:       "Mauritania",
			Alpha2Code: "MR",
			Alpha3Code: "MRT",
			Structure: bban.NewStructure(
				bban.NewBankCode(5, bban.Num),
				bban.NewBranchCode(5, bban.Num),
				bban.NewAccountNumber(11, bban.Num),
				bban.NewNationalCheckDigit(2, bban.Num),
			),
		},
		"MU": {
			Name:       "Mauritius",
			Alpha2Code: "MU",
			Alpha3Code: "MUS",
			Structure: bban.NewStructure(
				bban.NewBankCode(6, bban.AlphaUpper),
				bban.NewBranchCode(2, bban.Num),
				bban.NewAccountNumber(18, bban.AlphaNum),
			),
		},
		"MD": {
			Name:       "Moldova",
			Alpha2Code: "MD",
			Alpha3Code: "MDA",
			Structure: bban.NewStructure(
				bban.NewBankCode(2, bban.AlphaNum),
				bban.NewAccountNumber(18, bban.AlphaNum),
			),
		},
		"MC": {
			Name:       "Monaco",
			Alpha2Code: "MC",
			Alpha3Code: "MCO",
			Structure: bban.NewStructure(
				bban.NewBankCode(5, bban.Num),
				bban.NewBranchCode(5, bban.Num),
				bban.NewAccountNumber(11, bban.AlphaNum),
				bban.NewNationalCheckDigit(2, bban.Num),
			),
		},
		"ME": {
			Name:       "Montenegro",
			Alpha2Code: "ME",
			Alpha3Code: "MNE",
			Structure: bban.NewStructure(
				bban.NewBankCode(3, bban.Num),
				bban.NewAccountNumber(13, bban.Num),
				bban.NewNationalCheckDigit(2, bban.Num),
			),
		},
		"NL": {
			Name:       "Netherlands",
			Alpha2Code: "NL",
			Alpha3Code: "NLD",
			Structure: bban.NewStructure(
				bban.NewBankCode(4, bban.AlphaUpper),
				bban.NewAccountNumber(10, bban.Num),
			),
		},
		"NO": {
			Name:       "Norway",
			Alpha2Code: "NO",
			Alpha3Code: "NOR",
			Structure: bban.NewStructure(
				bban.NewBankCode(4, bban.Num),
				bban.NewAccountNumber(6, bban.Num),
				bban.NewNationalCheckDigit(1, bban.Num),
			),
		},
		"PK": {
			Name:       "Pakistan",
			Alpha2Code: "PK",
			Alpha3Code: "PAK",
			Structure: bban.NewStructure(
				bban.NewBankCode(4, bban.AlphaNum),
				bban.NewAccountNumber(16, bban.Num),
			),
		},
		"PS": {
			Name:       "Palestine",
			Alpha2Code: "PS",
			Alpha3Code: "PSE",
			Structure: bban.NewStructure(
				bban.NewBankCode(4, bban.AlphaUpper),
				bban.NewAccountNumber(21, bban.AlphaNum),
			),
		},
		"PL": {
			Name:       "Poland",
			Alpha2Code: "PL",
			Alpha3Code: "POL",
			Structure: bban.NewStructure(
				bban.NewBankCode(3, bban.Num),
				bban.NewBranchCode(4, bban.Num),
				bban.NewNationalCheckDigit(1, bban.Num),
				bban.NewAccountNumber(16, bban.Num),
			),
		},
		"PT": {
			Name:       "Portugal",
			Alpha2Code: "PT",
			Alpha3Code: "PRT",
			Structure: bban.NewStructure(
				bban.NewBankCode(4, bban.Num),
				bban.NewBranchCode(4, bban.Num),
				bban.NewAccountNumber(11, bban.Num),
				bban.NewNationalCheckDigit(2, bban.Num),
			),
		},
		"QA": {
			Name:       "Qatar",
			Alpha2Code: "QA",
			Alpha3Code: "QAT",
			Structure: bban.NewStructure(
				bban.NewBankCode(4, bban.AlphaUpper),
				bban.NewAccountNumber(21, bban.AlphaNum),
			),
		},
		"RO": {
			Name:       "Romania",
			Alpha2Code: "RO",
			Alpha3Code: "ROU",
			Structure: bban.NewStructure(
				bban.NewBankCode(4, bban.AlphaUpper),
				bban.NewAccountNumber(16, bban.AlphaNum),
			),
		},
		"SM": {
			Name:       "San Marino",
			Alpha2Code: "SM",
			Alpha3Code: "SMR",
			Structure: bban.NewStructure(
				bban.NewNationalCheckDigit(1, bban.AlphaUpper),
				bban.NewBankCode(5, bban.Num),
				bban.NewBranchCode(5, bban.Num),
				bban.NewAccountNumber(12, bban.AlphaNum),
			),
		},
		"SA": {
			Name:       "Saudi Arabia",
			Alpha2Code: "SA",
			Alpha3Code: "SAU",
			Structure: bban.NewStructure(
				bban.NewBankCode(2, bban.Num),
				bban.NewAccountNumber(18, bban.AlphaNum),
			),
		},
		"RS": {
			Name:       "Serbia",
			Alpha2Code: "RS",
			Alpha3Code: "SRB",
			Structure: bban.NewStructure(
				bban.NewBankCode(3, bban.Num),
				bban.NewAccountNumber(13, bban.Num),
				bban.NewNationalCheckDigit(2, bban.Num),
			),
		},
		"SK": {
			Name:       "Slovakia",
			Alpha2Code: "SK",
			Alpha3Code: "SVK",
			Structure: bban.NewStructure(
				bban.NewBankCode(4, bban.Num),
				bban.NewAccountNumber(16, bban.Num),
			),
		},
		"SI": {
			Name:       "Slovenia",
			Alpha2Code: "SI",
			Alpha3Code: "SVN",
			Structure: bban.NewStructure(
				bban.NewBankCode(2, bban.Num),
				bban.NewBranchCode(3, bban.Num),
				bban.NewAccountNumber(8, bban.Num),
				bban.NewNationalCheckDigit(2, bban.Num),
			),
		},
		"ES": {
			Name:       "Spain",
			Alpha2Code: "ES",
			Alpha3Code: "ESP",
			Structure: bban.NewStructure(
				bban.NewBankCode(4, bban.Num),
				bban.NewBranchCode(4, bban.Num),
				bban.NewNationalCheckDigit(2, bban.Num),
				bban.NewAccountNumber(10, bban.Num),
			),
		},
		"SE": {
			Name:       "Sweden",
			Alpha2Code: "SE",
			Alpha3Code: "SWE",
			Structure: bban.NewStructure(
				bban.NewBankCode(3, bban.Num),
				bban.NewAccountNumber(17, bban.Num),
			),
		},
		"CH": {
			Name:       "Switzerland",
			Alpha2Code: "CH",
			Alpha3Code: "CHE",
			Structure: bban.NewStructure(
				bban.NewBankCode(5, bban.Num),
				bban.NewAccountNumber(12, bban.AlphaNum),
			),
		},
		"TN": {
			Name:       "Tunisia",
			Alpha2Code: "TN",
			Alpha3Code: "TUN",
			Structure: bban.NewStructure(
				bban.NewBankCode(2, bban.Num),
				bban.NewBranchCode(3, bban.Num),
				bban.NewAccountNumber(15, bban.AlphaNum),
			),
		},
		"TR": {
			Name:       "Turkey",
			Alpha2Code: "TR",
			Alpha3Code: "TUR",
			Structure: bban.NewStructure(
				bban.NewBankCode(5, bban.Num),
				bban.NewNationalCheckDigit(1, bban.AlphaNum),
				bban.NewAccountNumber(16, bban.AlphaNum),
			),
		},
		"AE": {
			Name:       "United Arab Emirates",
			Alpha2Code: "AE",
			Alpha3Code: "ARE",
			Structure: bban.NewStructure(
				bban.NewBankCode(3, bban.Num),
				bban.NewAccountNumber(16, bban.AlphaNum),
			),
		},
		"GB": {
			Name:       "United Kingdom",
			Alpha2Code: "GB",
			Alpha3Code: "GBR",
			Structure: bban.NewStructure(
				bban.NewBankCode(4, bban.AlphaUpper),
				bban.NewBranchCode(6, bban.Num),
				bban.NewAccountNumber(8, bban.Num),
			),
		},
	}
)

// Exists returns true if country code exists.
func Exists(code string) bool {
	_, ok := countries[code]
	return ok
}

// Get returns country by given country code.
func Get(code string) (Country, bool) {
	country, ok := countries[code]
	return country, ok
}

// GetBbanStructure returns bban.Structure by given country code.
func GetBbanStructure(code string) (bban.Structure, bool) {
	if country, ok := Get(code); ok {
		return country.Structure, true
	}
	return bban.Structure{}, false
}
