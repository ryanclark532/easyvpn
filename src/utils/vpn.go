package utils

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func SetupVPNServer() error {
	if checkVPNCertificates() {
		return nil
	}

	err := GenerateDHKey()
	if err != nil {
		return err
	}

	err = GenerateRootCACertificate()
	if err != nil {
		return err
	}

	err = GenerateSignedCertificate("server")
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
	cmd := exec.Command("openvpn", "src/server-dev.conf")
	err := cmd.Run()
	return err
}

func StopOpenVPNServer() error {
	cmd := exec.Command("pkill", "openvpn")

	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("error stopping OpenVPN process: %v\n%s", err, output)
	}

	if strings.Contains(string(output), "no process found") {
		return fmt.Errorf("OpenVPN process not found")
	}

	fmt.Println("OpenVPN process stopped successfully.")
	return nil
}
