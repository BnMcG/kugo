package main

import (
	"crypto/x509"
	"testing"
	"time"
)

func TestPEMCertificateParsingWithCA(t *testing.T) {
	exampleCertificate := "LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSUN5RENDQWJDZ0F3SUJBZ0lCQURBTkJna3Foa2lHOXcwQkFRc0ZBREFWTVJNd0VRWURWUVFERXdwcmRXSmwKY201bGRHVnpNQjRYRFRFNE1USXpNREU0TURJd05Gb1hEVEk0TVRJeU56RTRNREl3TkZvd0ZURVRNQkVHQTFVRQpBeE1LYTNWaVpYSnVaWFJsY3pDQ0FTSXdEUVlKS29aSWh2Y05BUUVCQlFBRGdnRVBBRENDQVFvQ2dnRUJBTG1TCi9mM0FFTEk2bzQyOTArSnhyNGNJdEM0QmFsKytvTkwyZUpVZTRIUUlmOWpaTGhweGZmWHBBL1NTdzh5WTBpazEKdnFmaDlLZ2xtbGNHblROZm9lS0w1KzFVOHo1aWdwMU5LSS9qdG9meGxVMlFNaXY4aTVmZndNaExmczdCa0hNZQpuRjN5RnFtMkZsRG9aSE9weGlrRXlqWkNpMnZpcUtZM3FGWCt3VkFheGNpSURHalNQaDl5bTJRN3ZOcmRoVEFDCkdTRlVEdzUzS0JxVzdhWHF2dEpuTXJKTW5RdWlRUllHM0VRZ1F1dmU5TGlNekp2a0t3MEhYTUdNL1FuZzBvaFUKY3dlU3o1RE9SREpIaXl3c3hRSzVjQm1Tbjd5UUJaWkl5MzkzWW5vSDA0Vi95NkJraDJUeFVQNkVjcXFrTGkzago4Y3Z0Yk9qajRRTnZKenpkckZzQ0F3RUFBYU1qTUNFd0RnWURWUjBQQVFIL0JBUURBZ0trTUE4R0ExVWRFd0VCCi93UUZNQU1CQWY4d0RRWUpLb1pJaHZjTkFRRUxCUUFEZ2dFQkFKZE9BSVZ2Q2ZRTXRFN1RRMnpvTjhJeGNuOEsKZGR1ZEtmUURlRFdDeUZsZnM5bDI2OEQ3N1U2b0NJd041cURYRVhoNFNtM1JqemsvQThuZ1lZb3dOa01wZzN1WApQbWJWSUJDUHV0bXl5MWxIWDRkQjFsbDQxQ1BSVFFBTGRVZkFTZGFJZStMMzZVN2QrYkd2UmVmQmRuQjZJcEJLCmF6MmpLUnFveVdONXVuTzJwaHpXemoxbnFvelhvSVlpeEl4bjZHNzR2cmdNcmZFK2JGWUFjaE5qTXNBaWdBWVUKbE9LOWIyTDFsRTd5bk5CV2VjeStEK2NnSGpVRGY2aG5yVVdOQ1QvQUxrbm9MbkFFcFZTRUhoTWp5TURDb1p3QgpmVHc3Uyt4Rkd6VFRyL2JBdjY2MXExaENFTVVWWTRnU2o3ZExmZ080RHdoRkk2M0wwZlZXeHVxeElqZz0KLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="
	certificate, err := DecodeBase64EncodedPEMCertificate(exampleCertificate)
	if err != nil {
		t.Error(err)
	}

	if certificate.IsCA != true {
		t.Error("Could not parse CA status of certificate correctly")
	}

	if certificate.Subject.CommonName != "kubernetes" {
		t.Error("Could not parse subject common name correctly")
	}
}

func TestPEMCertificateParsingWithIssuedCert(t *testing.T) {
	exampleCertificate := "LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSUQyVENDQXNHZ0F3SUJBZ0lVRmt1VEpBSi9NVlpIQlJsYTNick8zb0dSTzhRd0RRWUpLb1pJaHZjTkFRRUwKQlFBd0ZURVRNQkVHQTFVRUF4TUthM1ZpWlhKdVpYUmxjekFlRncweE9UQTJNRGN3TnpFek1qUmFGdzB4T1RBMgpNRGN3TnpFNE5UUmFNRWt4RnpBVkJnTlZCQW9URG5ONWMzUmxiVHB0WVhOMFpYSnpNUk13RVFZRFZRUUxFd3ByCmRXSmxjbTVsZEdWek1Sa3dGd1lEVlFRREV4QnJkV0psY201bGRHVnpMV0ZrYldsdU1JSUJJakFOQmdrcWhraUcKOXcwQkFRRUZBQU9DQVE4QU1JSUJDZ0tDQVFFQTZvT2RHNWF6cld4aFozczJ2WkZMZmcrRXFDaWhESndzVmJlSgo0bkk4WUFOMktXZFRSV0p0am9Sb05qMEpwWWVOUG1HeGptZmZ3dnNGQk1BU0pzOHFnU0NlZlZncEhKSVBEdVhoCjJKVVpVdU40VE4yTFpYNEY2Q1JFVFNTYjJic05RWVhFTXFJSEpGZ05Wa2E0YkJDaWwrVUw0SEdidS9KY3czVkEKdTUvaXRIZ3l6TTIzeldmOTFLNDRwbWVPYzhTR01keGQwYXQwZHVjS201TGcvdWRQR050bGpjU2hlcWRqcS9ZQQpFRDR3NzFCNUlTRTJIVGRSQzR4elRxOStLa0JvMmUyRjRIWGx2VGhyNFJSRjFIbVpCRzlHYUtEdDIxVXJUSXBXCnp3cGYzWndVSTQ1WjdkUy9yVXF1M2VubUpXcU9SdVptdUFwMHQ1K1JqNC93VVIwYTdRSURBUUFCbzRIc01JSHAKTUE0R0ExVWREd0VCL3dRRUF3SURxREFkQmdOVkhTVUVGakFVQmdnckJnRUZCUWNEQVFZSUt3WUJCUVVIQXdJdwpIUVlEVlIwT0JCWUVGQWQwLzJMbHRoOGxqdWFqZmdBcUZKbG0zbnlVTUVJR0NDc0dBUVVGQndFQkJEWXdOREF5CkJnZ3JCZ0VGQlFjd0FvWW1hSFIwY0hNNkx5OTJZWFZzZEM1M2FYcGhjbVJvTG1GMEwzWXhMM0JyYVM5ck9ITXYKWTJFd0d3WURWUjBSQkJRd0VvSVFhM1ZpWlhKdVpYUmxjeTFoWkcxcGJqQTRCZ05WSFI4RU1UQXZNQzJnSzZBcApoaWRvZEhSd2N6b3ZMM1poZFd4MExuZHBlbUZ5WkdndVlYUXZkakV2Y0d0cEwyczRjeTlqY213d0RRWUpLb1pJCmh2Y05BUUVMQlFBRGdnRUJBQmF1RlhQcmhlb2NMWFI0UG1xb0dFcHhORnlmR1UrQzJvMnIwMHdxVENLeUc4MkcKbHFOOVJvRmdYQmVPSW42OVViQW1tMy9oaUxMaTJaQ3hlUHo1UjhMRHB4aEZuNU1kdWJ5OEp2L0IrTVd1V3hYRQo1TXlCcXVLbzIwa0dXam9lellIZ0dxU3lhYjY4VTc2NEd2TzR1cmNXK1VpbFJidTBDRUtKVTBOVjlnOWFjaUUvCkdiZlJjZkUvYy9pNDAyeWxrckxTcGdlMjQvbUVsVG5mdkppRnhWYWJUUm11bnRhTURPeHdEaDhPWHF0WkVid3QKUG9TWHB6eWNuMTVEUjFXTVg4Nk5oUFR0MTZZWGFIR0d5OUFOUUZhN3lzdUcxRkVpazczdGZMZGJjOStIbTM1aAovYWFKWGFvUHZ6SVBVMTVPcGt6RThCZlkzdUZFNEw5OEhEMHB3RXc9Ci0tLS0tRU5EIENFUlRJRklDQVRFLS0tLS0="
	certificate, err := DecodeBase64EncodedPEMCertificate(exampleCertificate)
	if err != nil {
		t.Error(err)
	}

	if certificate.IsCA != false {
		t.Error("Incorrectly parsed certificate as CA")
	}

	if certificate.Subject.CommonName != "kubernetes-admin" {
		t.Error("Could not parse subject common name correctly")
	}

	if len(certificate.Subject.Names) != 3 {
		t.Error("Could not parse correct number of subject names")
	}

	if len(certificate.DNSNames) != 1 {
		t.Error("Could not parse correct number of DNS names")
	}
}

func TestErrorsOnInvalidBase64(t *testing.T) {
	exampleCertificate := "hello-there"
	_, err := DecodeBase64EncodedPEMCertificate(exampleCertificate)
	if err == nil {
		t.Error("Did not error on invalid Base64 input")
	}
}

func TestExpiryVerificationNotExpired(t *testing.T) {
	duration, _ := time.ParseDuration("1m")
	cert := &x509.Certificate{
		NotAfter: time.Now().UTC().Add(duration),
	}

	if CertificateHasExpired(cert) {
		t.Error("Valid certificate returned as expired")
	}
}

func TestExpiryVerificationExpired(t *testing.T) {
	duration, _ := time.ParseDuration("-1h")
	cert := &x509.Certificate{
		NotAfter: time.Now().UTC().Add(duration),
	}

	if !CertificateHasExpired(cert) {
		t.Error("Expired certificate returned as valid")
	}
}
