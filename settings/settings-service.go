package settings

import (
	"bufio"
	"easyvpn/settings/settings_dtos"
	"easyvpn/utils"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"
)

func GetSettings() (*settings_dtos.Settings, error) {
	js, err := os.Open(`./settings.json`)
	if err != nil {
		return nil, err
	}
	defer js.Close()
	jsonBytes, err := io.ReadAll(js)
	if err != nil {
		return nil, err
	}
	var settings settings_dtos.Settings
	err = json.Unmarshal(jsonBytes, &settings)
	return &settings, err
}

func SetSettings(settings *settings_dtos.Settings) error {
	err := RewriteVPNConfig(settings)
	if err != nil {
		return err
	}
	json, err := json.Marshal(settings)
	if err != nil {
		return err
	}
	err = utils.WriteFile("./settings.json", json)
	return err
}

func RewriteVPNConfig(settings *settings_dtos.Settings) error {
	config, err := os.Open(`C:\Program Files\OpenVPN\config-auto\server-dev.ovpn`)
	if err != nil {
		return err
	}
	defer config.Close()
	var newConfig []string
	scanner := bufio.NewScanner(config)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, `push "dhcp-option DNS`) && settings.Vpn.DNSServer1 != nil && settings.Vpn.DNSServer2 != nil {
			newConfig = processDns(line, settings, newConfig)
		} else if strings.HasPrefix(line, "proto") && settings.Network.Protocol != nil {
			newConfig = processProtocol(line, settings, newConfig)
		} else if strings.HasPrefix(line, `push "route `) && settings.Vpn.VpnSubnet != nil {
			newConfig = processClientRoute(line, settings, newConfig)
		} else {
			newConfig = append(newConfig, line)
		}
	}

	content := []byte(strings.Join(newConfig, "\n"))
	err = utils.WriteFile(`C:\Program Files\OpenVPN\config-auto\server-dev.ovpn`, content)
	return err
}

func processDns(line string, settings *settings_dtos.Settings, newConfig []string) []string {
	newConfig = append(newConfig, fmt.Sprintf(`push "dhcp-option DNS %s"`, *settings.Vpn.DNSServer1))
	newConfig = append(newConfig, fmt.Sprintf(`push "dhcp-option DNS %s"`, *settings.Vpn.DNSServer2))
	return newConfig
}

func processProtocol(line string, settings *settings_dtos.Settings, newConfig []string) []string {
	var tcp = "tcp"
	var udp = "udp"
	if settings.Network.Protocol == &tcp {
		newConfig = append(newConfig, "proto tcp-server")
	} else if settings.Network.Protocol == &udp {
		newConfig = append(newConfig, "proto udp-server")
	}
	return newConfig
}

func processClientRoute(line string, settings *settings_dtos.Settings, newConfig []string) []string {
	newConfig = append(newConfig, fmt.Sprintf(`push "route %s %d"`, *settings.Vpn.VpnSubnet, *settings.Vpn.VpnSubnetMask))
	return newConfig
}
