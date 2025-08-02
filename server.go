package main

import (
	"crypto/x509"
	"fmt"
	"net"

	wl_crypto "github.com/wsva/lib_go/crypto"
)

type CertServer struct {
	CAName     string   `json:"CAName"`
	CommonName string   `json:"CommonName"`
	IPList     []string `json:"IPList"`
	Filename   string   `json:"Filename"`
}

func (c *CertServer) Generate() {
	cacrt, cakey, err := wl_crypto.TryLoadCertAndKeyFromDisk(PKIPath, c.CAName)
	if err != nil {
		fmt.Println(err)
		return
	}

	var ipList []net.IP
	for _, v := range c.IPList {
		ipList = append(ipList, net.ParseIP(v))
	}

	config := wl_crypto.CertConfig{
		CertConfigBase: wl_crypto.CertConfigBase{
			CommonName: c.CommonName,
			AltNames: wl_crypto.AltNames{
				IPs: ipList,
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
