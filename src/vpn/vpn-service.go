package vpn

import (
	"easyvpn/src/utils"
)

func GetVpnServerStatus() (string, error) {
	running, err := utils.IsProcessRunning("openvpn")
	if err != nil {
		return "", err
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
