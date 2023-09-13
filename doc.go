/*
Package banking provides banking utilities.

IBAN:

	package main

	import (
	    "fmt"
	    "log"

	    "github.com/jbub/banking/iban"
	)

	var (
	    testIban = iban.MustParse("BE68539007547034")
	)

	func main() {
	    ibn, err := iban.Parse("BE68539007547034")
	    if err != nil {
	        log.Fatal(err)
	    }

	    fmt.Println(ibn.BankCode())
	    fmt.Println(ibn.AccountNumber())

	    err = iban.Validate("BE68539007547034")
	    if err != nil {
	        log.Fatal(err)
	    }
	}

Swift:

	package main

	import (
	    "fmt"
	    "log"

	    "github.com/jbub/banking/swift"
	)

	var (
	    testSwift = swift.MustParse("DEUTDEFF500")
	)

	func main() {
	    swft, err := swift.Parse("DEUTDEFF500")
	    if err != nil {
	        log.Fatal(err)
	    }

	    fmt.Println(swft.BankCode())
	    fmt.Println(swft.CountryCode())

	    err = swift.Validate("DEUTDEFF500")
	    if err != nil {
	        log.Fatal(err)
	    }
	}
*/
package banking
