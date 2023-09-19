package utils

import "os"

func SetupVPNServer() error {
	err := GenerateRootKeyPair()
	if err != nil {
		return err
	}
	err = GenerateCertificateKeyPair("server.key", "server.crt")
	if err != nil {
		return err
	}

	err = checkVPNCertificates()
	if err != nil {
		return err
	}
	return nil
}

func checkVPNCertificates() error {
	_, err := os.Stat(KeyDir + "root.crt")
	if !os.IsNotExist(err) {
		return err
	}

	_, err = os.Stat(KeyDir + "root.key")
	if !os.IsNotExist(err) {
		return err
	}

	_, err = os.Stat(KeyDir + "server.crt")
	if !os.IsNotExist(err) {
		return err
	}

	_, err = os.Stat(KeyDir + "server.key")
	if !os.IsNotExist(err) {
		return err
	}
	return nil
}
