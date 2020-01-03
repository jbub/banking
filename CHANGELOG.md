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