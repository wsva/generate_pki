package main

import (
	"crypto/x509"
	"fmt"
	"net"

	wl_crypto "github.com/wsva/lib_go/crypto"
)

type CertServer struct {
	CAName     string `json:"CAName"`
	CommonName string `json:"CommonName"`
	AltNames   struct {
		DNSNames []string `json:"DNSNames"`
		IPs      []string `json:"IPs"`
	} `json:"AltNames"`
	Filename string `json:"Filename"`
}

func (c *CertServer) Generate() {
	cacrt, cakey, err := wl_crypto.TryLoadCertAndKeyFromDisk(PKIPath, c.CAName)
	if err != nil {
		fmt.Println(err)
		return
	}

	toIP := func(IPs []string) []net.IP {
		result := make([]net.IP, len(IPs))
		for k, v := range IPs {
			result[k] = net.ParseIP(v)
		}
		return result
	}

	config := wl_crypto.CertConfig{
		CertConfigBase: wl_crypto.CertConfigBase{
			CommonName: c.CommonName,
			AltNames: wl_crypto.AltNames{
				DNSNames: c.AltNames.DNSNames,
				IPs:      toIP(c.AltNames.IPs),
			},
			Usages: []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
		},
		PublicKeyAlgorithm: x509.ECDSA,
	}
	crt, key, err := wl_crypto.NewCertAndKey(cacrt, cakey, &config)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = wl_crypto.WriteCertAndKey(PKIPath, c.Filename, crt, key)
	if err != nil {
		fmt.Println(err)
		return
	}
}
