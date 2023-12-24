package utils

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func GenerateClientConfig(userName string) error {
	key, err := ReadFile(fmt.Sprintf(`C:\Program Files\OpenVPN\config-auto\keys\%s.key`, userName))
	if err != nil {
		return err
	}

	cert, err := ReadFile(fmt.Sprintf(`C:\Program Files\OpenVPN\config-auto\keys\%s.crt`, userName))
	if err != nil {
		return err
	}

	ca, err := ReadFile(`C:\Program Files\OpenVPN\config-auto\keys\root.crt`)
	if err != nil {
		return err
	}

	base, err := ReadFile("./vpn-config/base-client-dev.ovpn")
	if err != nil {
		return err
	}
	config := append(base, []byte("<ca>\n")...)
	config = append(config, ca...)
	config = append(config, []byte("</ca>\n")...)
	config = append(config, []byte("<cert>\n")...)
	config = append(config, cert...)
	config = append(config, []byte("</cert>\n")...)
	config = append(config, []byte("<key>\n")...)
	config = append(config, key...)
	config = append(config, []byte("</key>")...)
	err = WriteFile(fmt.Sprintf("./tmp/%s.ovpn", userName), config)
	return err
}

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

	err := RestartVPNServer()
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

func RestartVPNServer() error {
	status, err := GetVpnServerStatus()
	if err != nil {
		return err
	}
	if status == "running" {
		err = StopVPNServer()
	}
	go StartVPNServer()
	return err
}

func GetVpnServerStatus() (string, error) {
	cmd := exec.Command("sc", "query", "OpenVPNService")
	output, err := cmd.Output()
	if err != nil {
		return "unknown", err
	}

	if strings.Contains(string(output), "RUNNING") {
		initFinished, err := ContainsSequence(`C:\Program Files\OpenVPN\log\server-dev.log`, "Initialization Sequence Completed")
		if err != nil {
			return "unknown", err
		}
		if initFinished {
			return "running", nil
		} else {
			return "starting", nil
		}
	}
	return "notRunning", nil
}
