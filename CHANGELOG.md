## 0.7.0

* Use Go 1.17.

## 0.6.0

* Use Go 1.16.

## 0.5.0

* Add support for Egypt iban.
* Use Go 1.14.

## 0.4.0

* Fix Costa Rica iban.
* Add support for Ukraine iban.
* Add support for Seychelles iban.
* Add support for Saint Lucia iban.
* Add support for Kosovo iban.
* Add support for Iraq iban.
* Add support for Greenland iban.
* Add support for Faroe Islands iban.
* Add support for East Timor iban.
* Add support for bban.Zero.
* Drop unicode dependency.

## 0.3.0

* Do not use regexp to validate swift codes.
* Optimize calculateMod to not use strconv.ParseInt.
* Refactor bban validation.
* Move bban structure to iban to eliminate duplicate parsing.
* Drop fmt.Sprintf usage in calculateCheckDigit.
* Drop fmt.Sprintf usage in favour of string concat to save allocations.
* Refactor iban country code validation.
* Deprecate New and use Parse instead.
* Add support for Belarus iban.

## 0.2.0

* Add padding support to fix the Mauritius account number.
* Add support for iban currency.
* Add Makefile.
* Setup linter.

## 0.1.0

* Initial release.