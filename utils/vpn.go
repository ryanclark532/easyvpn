package utils

import (
	"fmt"
	"os"
	"os/exec"
)

func GenerateClientConfig(userName string) error{
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
	err = WriteFile(fmt.Sprintf("./vpn-config/temp/%s.ovpn", userName),config) 
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
