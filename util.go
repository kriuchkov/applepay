package applepay

import (
	"crypto/x509"
	"encoding/pem"

	"github.com/pkg/errors"
)

func loadRootCertificate(data []byte) (*x509.Certificate, error) {
	rootPEM, _ := pem.Decode(data)
	if rootPEM == nil {
		return nil, errors.New("couldn't decode the root certificate")
	}

	root, err := x509.ParseCertificate(rootPEM.Bytes)
	if err != nil {
		return nil, errors.Wrap(err, "couldn't parse the root certificate")
	}

	if !root.IsCA {
		return nil, errors.New("the certificate is not be a CA")
	}
	return root, nil
}
