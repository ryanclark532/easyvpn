package utils

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"math/big"
	"os"
	"time"
)

const KeyDir string = "./src/keys/"

func GenerateRootKeyPair() error {
	key, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return err
	}
	cert, err := createCertificate(key)
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

func GenerateCertificateKeyPair(keyFile string, certFile string) error {
	key, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return err
	}
	cert, err := createCertificate(key)
	if err != nil {
		return err
	}

	rootCert, rootKey, err := loadKeyPair(KeyDir+"root.crt", KeyDir+"root.key")
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

	err = savePrivateKey(KeyDir+keyFile, key)
	if err != nil {
		return err
	}

	return nil
}

func createCertificate(key *rsa.PrivateKey) (*x509.Certificate, error) {
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

	certDER, err := x509.CreateCertificate(rand.Reader, &template, &template, &key.PublicKey, key)
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
		err = os.WriteFile(filename, keyBytes, 0600)
	}
	return err
}

func loadKeyPair(certFile, keyFile string) (*x509.Certificate, *rsa.PrivateKey, error) {
	certPEM, err := os.ReadFile(certFile)
	if err != nil {
		return nil, nil, err
	}

	keyPEM, err := os.ReadFile(keyFile)
	if err != nil {
		return nil, nil, err
	}

	certBlock, _ := pem.Decode(certPEM)
	keyBlock, _ := pem.Decode(keyPEM)

	rootCert, err := x509.ParseCertificate(certBlock.Bytes)
	if err != nil {
		return nil, nil, err
	}

	rootKey, err := x509.ParsePKCS1PrivateKey(keyBlock.Bytes)
	if err != nil {
		return nil, nil, err
	}

	return rootCert, rootKey, nil
}
