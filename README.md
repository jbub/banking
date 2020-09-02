你好！
很冒昧用这样的方式来和你沟通，如有打扰请忽略我的提交哈。我是光年实验室（gnlab.com）的HR，在招Golang开发工程师，我们是一个技术型团队，技术氛围非常好。全职和兼职都可以，不过最好是全职，工作地点杭州。
我们公司是做流量增长的，Golang负责开发SAAS平台的应用，我们做的很多应用是全新的，工作非常有挑战也很有意思，是国内很多大厂的顾问。
如果有兴趣的话加我微信：13515810775  ，也可以访问 https://gnlab.com/，联系客服转发给HR。
# banking 
[![GoDoc](http://img.shields.io/badge/go-documentation-blue.svg?style=flat-square)](http://godoc.org/github.com/jbub/banking) 
[![Build Status](https://cloud.drone.io/api/badges/jbub/banking/status.svg)](https://cloud.drone.io/jbub/banking)
[![Go Report Card](https://goreportcard.com/badge/github.com/jbub/banking)](https://goreportcard.com/report/github.com/jbub/banking)

Banking library for Go.

## Install

```bash
go get github.com/jbub/banking
```

## Docs

http://godoc.org/github.com/jbub/banking

## Iban

```go
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
```

## Swift

```go
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
```
