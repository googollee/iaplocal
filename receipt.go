// Package iaplocal supports Apple Local In-App Purchase (IAP)
// receipt processing.
//
// It loads the receipt from binary, parses the receipt's
// attributes, and verifies the receipt signature and hash.
package iaplocal

import (
	"crypto/sha1"
	"crypto/x509"
	"encoding/asn1"
	"errors"
	"time"

	"github.com/fullsailor/pkcs7"
)

// Receipt is the receipt for an in-app purchase.
type Receipt struct {
	Quantity              int
	ProductID             string
	TransactionID         string
	PurchaseDate          time.Time
	OriginalTransactionID string
	OriginalPurchaseDate  time.Time
	ExpiresDate           time.Time
	WebOrderLineItemID    int
	CancellationDate      time.Time
}

// Receipts is the app receipt.
type Receipts struct {
	BundleID                   string
	ApplicationVersion         string
	OpaqueValue                []byte
	SHA1Hash                   []byte
	InApp                      []Receipt
	OriginalApplicationVersion string
	ExpirationDate             time.Time

	rawBundleID []byte
}

var (
	// ErrInvalidCertificate returns when parse a receipt
	// with invalid certificate from given root certificate.
	ErrInvalidCertificate = errors.New("iaplocal: invalid certificate in receipt")
	// ErrInvalidSignature returns when parse a receipt
	// which improperly signed.
	ErrInvalidSignature = errors.New("iaplocal: invalid signature of receipt")
)

// Parse parses a receipt binary which certificates with
// root certificate.
// Need decode to DER binary if recevied a base64 file.
func Parse(root *x509.Certificate, data []byte) (Receipts, error) {
	pkcs, err := pkcs7.Parse(data)
	if err != nil {
		return Receipts{}, err
	}

	if !verifyCertificates(root, pkcs.Certificates) {
		return Receipts{}, ErrInvalidCertificate
	}

	if !verifyPKCS(pkcs) {
		return Receipts{}, ErrInvalidSignature
	}

	return parsePKCS(pkcs)
}

// Verify verifys the receipts with given guid.
// TestReceiptValidate shows how to get GUID from string.
// Check https://developer.apple.com/library/ios/releasenotes/General/ValidateAppStoreReceipt/Chapters/ValidateLocally.html#//apple_ref/doc/uid/TP40010573-CH1-SW5
func (r *Receipts) Verify(guid []byte) bool {
	hash := sha1.New()
	hash.Write(guid)
	hash.Write(r.OpaqueValue)
	hash.Write([]byte(r.rawBundleID))
	sign := hash.Sum(nil)
	if len(sign) != len(r.SHA1Hash) {
		return false
	}
	for i := range sign {
		if sign[i] != r.SHA1Hash[i] {
			return false
		}
	}
	return true
}

func verifyCertificates(root *x509.Certificate, certs []*x509.Certificate) bool {
	roots := x509.NewCertPool()
	if root != nil {
		roots.AddCert(root)
	}
	for _, cert := range certs {
		roots.AddCert(cert)
	}
	opts := x509.VerifyOptions{
		Roots: roots,
	}
	for _, cert := range certs {
		chain, err := cert.Verify(opts)
		for _, c := range chain {
			if c[0] == c[1] {
				// self certificate
				return false
			}
		}
		if err != nil {
			return false
		}
	}
	return true
}

func verifyPKCS(pkcs *pkcs7.PKCS7) bool {
	return pkcs.Verify() == nil
}

type attribute struct {
	Type    int
	Version int
	Value   []byte
}

func parsePKCS(pkcs *pkcs7.PKCS7) (ret Receipts, err error) {
	var r asn1.RawValue
	_, err = asn1.Unmarshal(pkcs.Content, &r)
	if err != nil {
		return
	}
	rest := r.Bytes
	for len(rest) > 0 {
		var ra attribute
		rest, err = asn1.Unmarshal(rest, &ra)
		if err != nil {
			return
		}
		switch ra.Type {
		case 2:
			if _, err = asn1.Unmarshal(ra.Value, &ret.BundleID); err != nil {
				return
			}
			ret.rawBundleID = ra.Value
		case 3:
			if _, err = asn1.Unmarshal(ra.Value, &ret.ApplicationVersion); err != nil {
				return
			}
		case 4:
			ret.OpaqueValue = ra.Value
		case 5:
			ret.SHA1Hash = ra.Value
		case 17:
			var inApp Receipt
			inApp, err = parseInApp(ra.Value)
			if err != nil {
				return
			}
			ret.InApp = append(ret.InApp, inApp)
		case 19:
			if _, err = asn1.Unmarshal(ra.Value, &ret.OriginalApplicationVersion); err != nil {
				return
			}
		case 21:
			ret.ExpirationDate, err = asn1ParseTime(ra.Value)
			if err != nil {
				return
			}
		}
	}
	return
}

func parseInApp(data []byte) (ret Receipt, err error) {
	var r asn1.RawValue
	_, err = asn1.Unmarshal(data, &r)
	if err != nil {
		return
	}
	data = r.Bytes
	for len(data) > 0 {
		var ra attribute
		data, err = asn1.Unmarshal(data, &ra)
		if err != nil {
			return
		}
		switch ra.Type {
		case 1701:
			if _, err = asn1.Unmarshal(ra.Value, &ret.Quantity); err != nil {
				return
			}
		case 1702:
			if _, err = asn1.Unmarshal(ra.Value, &ret.ProductID); err != nil {
				return
			}
		case 1703:
			if _, err = asn1.Unmarshal(ra.Value, &ret.TransactionID); err != nil {
				return
			}
		case 1704:
			ret.PurchaseDate, err = asn1ParseTime(ra.Value)
			if err != nil {
				return
			}
		case 1705:
			if _, err = asn1.Unmarshal(ra.Value, &ret.OriginalTransactionID); err != nil {
				return
			}
		case 1706:
			ret.OriginalPurchaseDate, err = asn1ParseTime(ra.Value)
			if err != nil {
				return
			}
		case 1708:
			ret.ExpiresDate, err = asn1ParseTime(ra.Value)
			if err != nil {
				return
			}
		case 1711:
			if _, err = asn1.Unmarshal(ra.Value, &ret.WebOrderLineItemID); err != nil {
				return
			}
		case 1712:
			ret.CancellationDate, err = asn1ParseTime(ra.Value)
			if err != nil {
				return
			}
		}
	}
	return
}

func asn1ParseTime(data []byte) (time.Time, error) {
	var str string
	if _, err := asn1.Unmarshal(data, &str); err != nil {
		return time.Time{}, err
	}
	if str == "" {
		return time.Time{}, nil
	}
	return time.Parse(time.RFC3339, str)
}
