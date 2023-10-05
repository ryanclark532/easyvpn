package vpn

import (
	"easyvpn/src/utils"
	"fmt"
)

func GetVpnServerStatus() (string, error) {
	running, err := utils.IsProcessRunning("openvpn")
	if err != nil {
		return "unknown", err
	}
	if !running {
		return "notRunning", nil
	}

	containsInit, err := utils.ContainsSequence("src/log/openvpn.log", "Initialization Sequence Completed")
	if containsInit {
		return "running", err
	}
	return "starting", nil
}

func VpnOperation(operation string) error {
	var err error
	fmt.Println(operation)
	switch operation {
	case "start":
		go utils.StartVPNServer()
		return nil
	case "stop":
		err = utils.StopVPNServer()
		return err
	case "restart":
		err = utils.StopVPNServer()
		go utils.StartVPNServer()
		return err
	}
	return nil
}

func GetActiveConnections() error {
	result, err := utils.TelnetCMD("status")
	if err != nil {
		return err
	}
	fmt.Println(result)
	return nil
}
