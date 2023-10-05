package utils

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"math/big"
	"os"
	"os/exec"
	"time"
)

var KeyDir = "./src/keys/"

func GenerateRootCACertificate() error {
	privateKey, err := ecdsa.GenerateKey(elliptic.P384(), rand.Reader)
	if err != nil {
		return err
	}

	template := x509.Certificate{
		SerialNumber: big.NewInt(1),
		Subject: pkix.Name{
			Organization: []string{"Root CA"},
		},
		NotBefore:             time.Now(),
		NotAfter:              time.Now().AddDate(10, 0, 0), // Valid for 10 years
		KeyUsage:              x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature | x509.KeyUsageCertSign,
		BasicConstraintsValid: true,
		IsCA:                  true,
	}

	derBytes, err := x509.CreateCertificate(rand.Reader, &template, &template, &privateKey.PublicKey, privateKey)
	if err != nil {
		return err
	}

	keyBytes, err := x509.MarshalECPrivateKey(privateKey)
	if err != nil {
		return err
	}

	certFile, err := os.Create(KeyDir + "root.crt")
	if err != nil {
		return err
	}
	defer certFile.Close()

	keyFile, err := os.Create(KeyDir + "root.key")
	if err != nil {
		return err
	}
	defer keyFile.Close()

	pem.Encode(certFile, &pem.Block{Type: "CERTIFICATE", Bytes: derBytes})
	pem.Encode(keyFile, &pem.Block{Type: "EC PRIVATE KEY", Bytes: keyBytes})

	return nil
}

func GenerateSignedCertificate(name string) error {
	rootCACertData, err := os.ReadFile(KeyDir + "root.crt")
	if err != nil {
		return err
	}

	rootCAKeyData, err := os.ReadFile(KeyDir + "root.key")
	if err != nil {
		return err
	}

	rootCACert, _ := pem.Decode(rootCACertData)
	rootCAKey, _ := pem.Decode(rootCAKeyData)

	rootCA, err := x509.ParseCertificate(rootCACert.Bytes)
	if err != nil {
		return err
	}

	rootPrivateKey, err := x509.ParseECPrivateKey(rootCAKey.Bytes)
	if err != nil {
		return err
	}

	privateKey, err := ecdsa.GenerateKey(elliptic.P384(), rand.Reader)
	if err != nil {
		return err
	}

	template := x509.Certificate{
		SerialNumber: big.NewInt(2),
		Subject: pkix.Name{
			Organization: []string{"Example Organization"},
			CommonName:   name,
		},
		NotBefore:   time.Now(),
		NotAfter:    time.Now().AddDate(1, 0, 0),
		KeyUsage:    x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature,
		ExtKeyUsage: []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth, x509.ExtKeyUsageServerAuth},
	}

	derBytes, err := x509.CreateCertificate(rand.Reader, &template, rootCA, &privateKey.PublicKey, rootPrivateKey)
	if err != nil {
		return err
	}

	keyBytes, err := x509.MarshalECPrivateKey(privateKey)
	if err != nil {
		return err
	}

	certFile, err := os.Create(KeyDir + name + ".crt")
	if err != nil {
		return err
	}
	defer certFile.Close()

	keyFile, err := os.Create(KeyDir + name + ".key")
	if err != nil {
		return err
	}
	defer keyFile.Close()

	pem.Encode(certFile, &pem.Block{Type: "CERTIFICATE", Bytes: derBytes})
	pem.Encode(keyFile, &pem.Block{Type: "EC PRIVATE KEY", Bytes: keyBytes})

	return nil
}

func GenerateDHKey() error {
	_, err := os.Stat(KeyDir + "dh.pem")
	if os.IsNotExist(err) {
		cmd := exec.Command("openssl", "dhparam", "-out", KeyDir+"dh.pem", "2048")
		err = cmd.Run()
		return err
	}
	return nil
}
