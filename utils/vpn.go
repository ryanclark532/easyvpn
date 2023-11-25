package utils

import (
	"fmt"
	"os"
	"os/exec"
)

func SetupVPNServer() error {
	path := `C:\Program Files\OpenVPN\config-auto\`
	if _, err := os.Stat(path + `keys\`); os.IsNotExist(err) {
		err := os.Mkdir(path+`keys\`, 0777)
		if err != nil {
			return err
		}
		keychan := make(chan error)
		defer close(keychan)
		go func() {
			err := GenerateDHKey(path + `keys\`)
			if err != nil {
				keychan <- err
			}
			err = GenerateRootCACertificate(path + `keys\`)
			if err != nil {
				keychan <- err
			}
			err = GenerateSignedCertificate(path+`keys\`, "server")
			if err != nil {
				keychan <- err
			}
		}()

		err = <-keychan
		return err
	}

	err := CopyFile(`.\vpn-config\server-dev.ovpn`, path+`\server-dev.ovpn`)
	if err != nil {
		return err
	}
	/*
		err = StopVPNServer()
		if err != nil {
			return err
		}
		err = StartVPNServer()
	*/
	return err
}

func StartVPNServer() error {
	cmd := exec.Command("sc", "start", "OpenVPNService")
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("failed to start service: %v", err)
	}
	return nil
}

func StopVPNServer() error {
	cmd := exec.Command("sc", "stop", "OpenVPNService")
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("failed to stop service: %v", err)
	}
	return nil
}

func GetActiveUsers() {
	// call telnet cmd

	//format users
}
