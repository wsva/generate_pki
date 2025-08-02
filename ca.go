package main

import (
	"crypto/x509"
	"fmt"

	wl_crypto "github.com/wsva/lib_go/crypto"
)

type CA struct {
	CommonName   string `json:"CommonName"`
	Organization string `json:"Organization"`
	Filename     string `json:"Filename"`
}

func (c *CA) Generate() {
	config := wl_crypto.CertConfig{
		CertConfigBase: wl_crypto.CertConfigBase{
			CommonName:   c.CommonName,
			Organization: []string{c.Organization},
		},
		PublicKeyAlgorithm: x509.ECDSA,
	}
	cacrt, cakey, err := wl_crypto.NewCertificateAuthority(&config)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = wl_crypto.WriteCertAndKey(PKIPath, c.Filename, cacrt, cakey)
	if err != nil {
		fmt.Println(err)
		return
	}
}
