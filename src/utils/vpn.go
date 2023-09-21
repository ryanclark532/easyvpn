package utils

import (
	"os"
	"os/exec"
)

func SetupVPNServer() error {
	if checkVPNCertificates() {
		return nil
	}

	err := GenerateDHKey()
	if err != nil {
		return err
	}

	err = GenerateRootKeyPair()
	if err != nil {
		return err
	}

	err = GenerateCertificateKeyPair("server.crt")
	if err != nil {
		return err
	}

	return nil
}

func checkVPNCertificates() bool {
	_, err := os.Stat(KeyDir + "root.crt")
	if os.IsNotExist(err) {
		return false
	}

	_, err = os.Stat(KeyDir + "root.key")
	if os.IsNotExist(err) {
		return false
	}

	_, err = os.Stat(KeyDir + "server.crt")
	if os.IsNotExist(err) {
		return false
	}

	_, err = os.Stat(KeyDir + "server.key")
	if os.IsNotExist(err) {
		return false
	}

	_, err = os.Stat(KeyDir + "dh.pem")
	if os.IsNotExist(err) {
		return false
	}
	return true
}

func StartVPNServer() error {
	cmd := exec.Command("openvpn", "server.conf")
	err := cmd.Run()
	return err
}
