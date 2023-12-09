package utils

import (
	"easyvpn/settings/settings_dtos"
	"fmt"
)

func AppendModifyableSettings(newFile []string, settings *settings_dtos.Settings) []string {
	dns1 := fmt.Sprintf("push \"dhcp-option DNS %s\"", *settings.Vpn.DNSServer1)
	dns2 := fmt.Sprintf("push \"dhcp-option DNS %s\"", *settings.Vpn.DNSServer2)
	ip := fmt.Sprintf("server %s 255.255.255.0", *settings.Network.IPAddress)
	port := fmt.Sprintf("port %d", *settings.Vpn.Port)
	gateway := "push \"redirect-gateway def1 bypass-dhcp\""
	newFile = append(newFile, "\n")
	newFile = append(newFile, "#Modifiable Settings")
	newFile = append(newFile, dns1)
	newFile = append(newFile, dns2)
	newFile = append(newFile, ip)
	newFile = append(newFile, port)

	useGateway := *settings.Vpn.UseAsGateway
	if useGateway {
		newFile = append(newFile, gateway)
	}

	return newFile
}
