package settings

import (
	"bufio"
	"easyvpn/settings/settings_dtos"
	"easyvpn/utils"
	"easyvpn/vpn"
	"encoding/json"
	"os"
	"strings"
)

func SetSettings(settings *settings_dtos.Settings) error {
	utils.Settings = settings
	err := RewriteVPNConfig()
	if err != nil {
		return err
	}
	json, err := json.Marshal(settings)
	if err != nil {
		return err
	}
	err = utils.WriteFile("./settings.json", json)
	if err != nil {
		return err
	}

	err = vpn.VpnOperation("restart")
	return err
}

func RewriteVPNConfig() error {
	content, err := os.Open(`C:\Program Files\OpenVPN\config-auto\server-dev.ovpn`)
	if err != nil {
		return err
	}
	defer content.Close()
	scanner := bufio.NewScanner(content)
	var newFile []string
	for scanner.Scan() {
		line := scanner.Text()
		if line == "#Modifiable Settings" {
			break
		}
		newFile = append(newFile, line)
	}
	newFile = utils.AppendModifyableSettings(newFile, *utils.Settings)

	err = utils.WriteFile(`C:\Program Files\OpenVPN\config-auto\server-dev.ovpn`, []byte(strings.Join(newFile, "\n")))
	return err
}
