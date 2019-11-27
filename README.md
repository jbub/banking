# banking 
[![GoDoc](http://img.shields.io/badge/go-documentation-blue.svg?style=flat-square)](http://godoc.org/github.com/anecsoiu/banking) 
[![Build Status](https://cloud.drone.io/api/badges/jbub/banking/status.svg)](https://cloud.drone.io/jbub/banking)
[![Go Report Card](https://goreportcard.com/badge/github.com/anecsoiu/banking)](https://goreportcard.com/report/github.com/anecsoiu/banking)

Banking library for Go.

## Install

```bash
go get github.com/anecsoiu/banking
```

## Docs

http://godoc.org/github.com/anecsoiu/banking

## Iban

```go
package main

import (
    "fmt"
    "log"

    "github.com/anecsoiu/banking/iban"
)

var (
    testIban = iban.MustParse("BE68539007547034")
)

func main() {
    ibn, err := iban.New("BE68539007547034")
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
```

## Swift

```go
package main

import (
    "fmt"
    "log"

    "github.com/anecsoiu/banking/swift"
)

var (
    testSwift = swift.MustParse("DEUTDEFF500")
)

func main() {
    swft, err := swift.New("DEUTDEFF500")
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
```