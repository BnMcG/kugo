package main

import (
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"time"
)

// DecodeBase64EncodedPEMCertificate parses PEM data into Certificate object
func DecodeBase64EncodedPEMCertificate(b64Data string) (*x509.Certificate, error) {
	pemBytes, err := base64.StdEncoding.DecodeString(b64Data)
	if err != nil {
		return &x509.Certificate{}, err
	}

	pem, _ := pem.Decode(pemBytes)
	if pem == nil {
		return &x509.Certificate{}, errors.New("could not find any PEM data")
	}

	certificate, err := x509.ParseCertificate(pem.Bytes)
	if err != nil {
		return &x509.Certificate{}, err
	}

	return certificate, nil
}

// CertificateHasExpired verifies whether or not the given certificate has expired
func CertificateHasExpired(certificate *x509.Certificate) bool {
	return time.Now().UTC().After(certificate.NotAfter)
}
