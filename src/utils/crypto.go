package utils

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"errors"
	"math/big"
	"os"
	"os/exec"
	"time"
)

var KeyDir = "./src/keys/"

func GenerateRootKeyPair() error {
	key, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return err
	}

	cert, err := createCertificate(&key.PublicKey, key)
	if err != nil {
		return err
	}

	certDer, err := selfSignCertificate(cert, key)
	if err != nil {
		return err
	}

	err = saveCertificate(KeyDir+"root.crt", certDer)
	if err != nil {
		return err
	}

	err = savePrivateKey(KeyDir+"root.key", key)
	if err != nil {
		return err
	}

	return nil
}

func GenerateCertificateKeyPair(certFile string) error {
	rootCert, rootKey, err := readCertificatePair(KeyDir+"root.crt", KeyDir+"root.key")
	if err != nil {
		return err
	}

	cert, err := createChildCertificate(&rootKey.PublicKey, rootKey, rootCert)
	if err != nil {
		return err
	}

	certDer, err := signCertificate(cert, rootKey, rootCert)
	if err != nil {
		return err
	}

	err = saveCertificate(KeyDir+certFile, certDer)
	if err != nil {
		return err
	}

	return nil
}

func createCertificate(pub *rsa.PublicKey, priv *rsa.PrivateKey) (*x509.Certificate, error) {
	serialNumber, err := rand.Int(rand.Reader, new(big.Int).Lsh(big.NewInt(1), 128))
	if err != nil {
		return nil, err
	}

	template := x509.Certificate{
		SerialNumber: serialNumber,
		Subject: pkix.Name{
			CommonName: "Root Certificate",
		},
		NotBefore:             time.Now(),
		NotAfter:              time.Now().AddDate(10, 0, 0),
		KeyUsage:              x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature | x509.KeyUsageCertSign,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
		BasicConstraintsValid: true,
	}
	certDER, err := x509.CreateCertificate(rand.Reader, &template, &template, pub, priv)
	if err != nil {
		return nil, err
	}

	cert, err := x509.ParseCertificate(certDER)
	if err != nil {
		return nil, err
	}

	return cert, nil
}

func createChildCertificate(pub *rsa.PublicKey, priv *rsa.PrivateKey, parentCert *x509.Certificate) (*x509.Certificate, error) {
	serialNumber, err := rand.Int(rand.Reader, new(big.Int).Lsh(big.NewInt(1), 128))
	if err != nil {
		return nil, err
	}

	template := x509.Certificate{
		SerialNumber: serialNumber,
		Subject: pkix.Name{
			CommonName: "Root Certificate",
		},
		NotBefore:             time.Now(),
		NotAfter:              time.Now().AddDate(10, 0, 0),
		KeyUsage:              x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature | x509.KeyUsageCertSign,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
		BasicConstraintsValid: true,
	}
	certDER, err := x509.CreateCertificate(rand.Reader, &template, parentCert, pub, priv)
	if err != nil {
		return nil, err
	}

	cert, err := x509.ParseCertificate(certDER)
	if err != nil {
		return nil, err
	}

	return cert, nil
}

func signCertificate(cert *x509.Certificate, rootKey *rsa.PrivateKey, rootCert *x509.Certificate) ([]byte, error) {
	certDER, err := x509.CreateCertificate(rand.Reader, cert, rootCert, &rootKey.PublicKey, rootKey)
	if err != nil {
		return nil, err
	}

	return pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: certDER}), nil

}

func selfSignCertificate(cert *x509.Certificate, selfKey *rsa.PrivateKey) ([]byte, error) {
	certDER, err := x509.CreateCertificate(rand.Reader, cert, cert, &selfKey.PublicKey, selfKey)
	if err != nil {
		return nil, err
	}

	return pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: certDER}), nil
}

func saveCertificate(filename string, cert []byte) error {
	_, err := os.Stat(filename)
	if os.IsNotExist(err) {
		err = os.WriteFile(filename, cert, 0644)
	}
	return err
}

func savePrivateKey(filename string, key *rsa.PrivateKey) error {
	_, err := os.Stat(filename)
	if os.IsNotExist(err) {
		keyBytes := x509.MarshalPKCS1PrivateKey(key)

		pemBlock := &pem.Block{
			Type:  "RSA PRIVATE KEY",
			Bytes: keyBytes,
		}

		pemData := pem.EncodeToMemory(pemBlock)

		err = os.WriteFile(filename, pemData, 0600)
	}
	return err
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

func readCertificatePair(certPath, keyPath string) (*x509.Certificate, *rsa.PrivateKey, error) {
	certData, err := os.ReadFile(certPath)
	if err != nil {
		return nil, nil, err
	}

	certBlock, _ := pem.Decode(certData)
	if certBlock == nil {
		return nil, nil, errors.New("failed to decode certificate PEM block")
	}

	cert, err := x509.ParseCertificate(certBlock.Bytes)
	if err != nil {
		return nil, nil, err
	}

	keyData, err := os.ReadFile(keyPath)
	if err != nil {
		return nil, nil, err
	}

	keyBlock, _ := pem.Decode(keyData)
	if keyBlock == nil {
		return nil, nil, errors.New("failed to decode private key PEM block")
	}

	key, err := x509.ParsePKCS1PrivateKey(keyBlock.Bytes)
	if err != nil {
		return nil, nil, err
	}

	return cert, key, nil
}
