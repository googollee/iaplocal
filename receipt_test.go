package iaplocal

import (
	"encoding/base64"
	"log"
	"testing"
)

func TestCertCheck(t *testing.T) {
	// at := assert.New(t)
	// _, err := Parse(nil, base64ToBytes(receipt1))
	// at.Equal(ErrInvalidCertificate, err)
}

func TestParseReceipt1(t *testing.T) {
	// rootCA, err := x509.ParseCertificate(base64ToBytes(rootCA))
	// if err != nil {
	// 	t.Fatal("parse ca error:", err)
	// }
	// receipt, err := Parse(rootCA, base64ToBytes(receipt1))
	// if err != nil {
	// 	t.Fatal("parse receipt error:", err)
	// }
	// at := assert.New(t)
	// at.Equal("1", receipt.ApplicationVersion)
	// at.Equal([]byte{0xf6, 0xf0, 0xf5, 0x8b, 0x39, 0xaf, 0x26, 0xe2, 0x51, 0x2b, 0x52, 0xad, 0xa1, 0xed, 0xcd, 0x4a}, receipt.OpaqueValue)
	// at.Equal([]byte{0xc5, 0xb4, 0x83, 0x90, 0xcf, 0xea, 0x65, 0x89, 0xfd, 0x63, 0xad, 0x72, 0x2c, 0x8, 0x1c, 0xcb, 0x1f, 0x6d, 0xbd, 0x28}, receipt.SHA1Hash)
	// at.Equal(2, len(receipt.InApp))
	// at.Equal("1.0", receipt.OriginalApplicationVersion)
	// at.True(receipt.ExpirationDate.IsZero())
	//
	// inApp1 := receipt.InApp[0]
	// at.Equal(1, inApp1.Quantity)
	// at.Equal("1000000225325901", inApp1.TransactionID)
	// at.Equal("2016-07-23T06:21:11Z", inApp1.PurchaseDate.Format(time.RFC3339))
	// at.Equal("1000000225325901", inApp1.OriginalTransactionID)
	// at.Equal("2016-07-23T06:21:11Z", inApp1.OriginalPurchaseDate.Format(time.RFC3339))
	// at.True(inApp1.ExpiresDate.IsZero())
	// at.Equal(0, inApp1.WebOrderLineItemID)
	// at.True(inApp1.CancellationDate.IsZero())
	//
	// inApp2 := receipt.InApp[1]
	// at.Equal(1, inApp2.Quantity)
	// at.Equal("1000000225334938", inApp2.TransactionID)
	// at.Equal("2016-07-23T11:00:59Z", inApp2.PurchaseDate.Format(time.RFC3339))
	// at.Equal("1000000225334938", inApp2.OriginalTransactionID)
	// at.Equal("2016-07-23T11:00:59Z", inApp2.OriginalPurchaseDate.Format(time.RFC3339))
	// at.True(inApp2.ExpiresDate.IsZero())
	// at.Equal(0, inApp2.WebOrderLineItemID)
	// at.True(inApp2.CancellationDate.IsZero())
}

func TestReceiptValidate(t *testing.T) {
	// at := assert.New(t)
	// rootCA, err := x509.ParseCertificate(base64ToBytes(rootCA))
	// if err != nil {
	// 	t.Fatal("parse ca error:", err)
	// }

	// receipt, err := Parse(rootCA, base64ToBytes(receipt1))
	// if err != nil {
	// 	t.Fatal("parse receipt error:", err)
	// }
	//
	// guid, err := uuid.FromString(hostGUID)
	// if err != nil {
	// 	t.Fatal("parse guid error:", err)
	// }
	// at.True(receipt.Verify(guid.Bytes()))
}

func base64ToBytes(b64 string) []byte {
	d, err := base64.StdEncoding.DecodeString(b64)
	if err != nil {
		log.Fatalln("decode error:", err)
	}
	return d
}

var (
	rootCA   = `MIIEuzCCA6OgAwIBAgIBAjANBgkqhkiG9w0BAQUFADBiMQswCQYDVQQGEwJVUzETMBEGA1UEChMKQXBwbGUgSW5jLjEmMCQGA1UECxMdQXBwbGUgQ2VydGlmaWNhdGlvbiBBdXRob3JpdHkxFjAUBgNVBAMTDUFwcGxlIFJvb3QgQ0EwHhcNMDYwNDI1MjE0MDM2WhcNMzUwMjA5MjE0MDM2WjBiMQswCQYDVQQGEwJVUzETMBEGA1UEChMKQXBwbGUgSW5jLjEmMCQGA1UECxMdQXBwbGUgQ2VydGlmaWNhdGlvbiBBdXRob3JpdHkxFjAUBgNVBAMTDUFwcGxlIFJvb3QgQ0EwggEiMA0GCSqGSIb3DQEBAQUAA4IBDwAwggEKAoIBAQDkkakJH5HbHkdQ6wXtXnmELes2oldMVeyLGYne+Uts9QerIjAC6Bg++FAJ039BqJj50cpmnCRrEdCju+QbKsMflZ56DKRHi1vUFjczy8QPTc4UadHJGXL1XQ7Vf1+b8iUDulWPTV0N8WQ1IxVLFVkds5T39pyez1C6wVhQZ48ItCD3y6wsIG9wtj8BMIy3Q88PnT3zK0koGsj+zrW5DtleHNbLPbU6rfQPDgCSC7EhFi501TwN22IWq6NxkkdTVcGvL0Gz+PvjcM3mo0xFfh9Ma1CWQYnEdGILEINBhzOKgbEwWOxaBDKMaLOPHd5lc/9nXmW8Sdh2nzMUZaF3lMktAgMBAAGjggF6MIIBdjAOBgNVHQ8BAf8EBAMCAQYwDwYDVR0TAQH/BAUwAwEB/zAdBgNVHQ4EFgQUK9BpR5R2Cf70a40uQKb3R01/CF4wHwYDVR0jBBgwFoAUK9BpR5R2Cf70a40uQKb3R01/CF4wggERBgNVHSAEggEIMIIBBDCCAQAGCSqGSIb3Y2QFATCB8jAqBggrBgEFBQcCARYeaHR0cHM6Ly93d3cuYXBwbGUuY29tL2FwcGxlY2EvMIHDBggrBgEFBQcCAjCBthqBs1JlbGlhbmNlIG9uIHRoaXMgY2VydGlmaWNhdGUgYnkgYW55IHBhcnR5IGFzc3VtZXMgYWNjZXB0YW5jZSBvZiB0aGUgdGhlbiBhcHBsaWNhYmxlIHN0YW5kYXJkIHRlcm1zIGFuZCBjb25kaXRpb25zIG9mIHVzZSwgY2VydGlmaWNhdGUgcG9saWN5IGFuZCBjZXJ0aWZpY2F0aW9uIHByYWN0aWNlIHN0YXRlbWVudHMuMA0GCSqGSIb3DQEBBQUAA4IBAQBcNplMLXi37Yyb3PN3m/J20ncwT8EfhYOFG5k9RzfyqZtAjizUsZAS2L70c5vu0mQPy3lPNNiiPvl4/2vIB+x9OYOLUyDTOMSxv5pPCmv/K/xZpwUJfBdAVhEedNO3iyM7R6PVbyTi69G3cN8PReEnyvFteO3ntRcXqNx+IjXKJdXZD9Zr1KIkIxH3oayPc4FgxhtbCS+SsvhESPBgOJ4V9T0mZyCKM2r3DYLP3uujL/lTaltkwGMzd/c6ByxW69oPIQ7aunMZT7XZNn/Bh1XZp5m5MkL72NVxnn6hUrcbvZNCJBIqxw8dtk2cXmPIS4AXUKqK1drk/NAJBzewdXUh`
	hostGUID = `urn:uuid:xxxx-xxx-xxx-xxx`
	receipt1 = ``
)
