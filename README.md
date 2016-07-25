[![GoDoc](https://godoc.org/github.com/googollee/iaplocal?status.svg)](https://godoc.org/github.com/googollee/iaplocal)
[![Build Status](https://travis-ci.org/googollee/iaplocal.svg?branch=master)](https://travis-ci.org/googollee/iaplocal)

# Description


`iaplocal` is a Go library that supports Apple Local In-App Purchase
(IAP) receipt processing.

- Verify the receipt signature against [App Root CA certificate](https://www.apple.com/certificateauthority/).
- Parse the receipt from binary, extract in-app receipts.
- Validate the receipts hash with [host GUID](https://developer.apple.com/library/ios/releasenotes/General/ValidateAppStoreReceipt/Chapters/ValidateLocally.html#//apple_ref/doc/uid/TP40010573-CH1-SW5).

# Installation

```
go install github.com/googollee/iaplocal
```

# Usage

The simplest possible usage is:

```go
rootBytes, _ := ioutil.ReadFile("./AppleComputerRootCertificate.cer")
rootCA, _ := x509.ParseCertificate(rootBytes)

receiptB64, _ := ioutil.ReadFile("./receipt_b64")
receiptBytes, _ := base64.StdEncoding.DecodeString(string(receiptB64))
receipt, _ := iaplocal.Parse(rootCA, receiptBytes)

guid, _ := uuid.FromString(hostGUID)
receipt.Verify(guid)
```

# License

See `LICENSE`.
