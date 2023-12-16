package utils

import (
	"easyvpn/settings/settings_dtos"
	"fmt"
)

func AppendModifyableSettings(newFile []string, settings settings_dtos.Settings) []string {
	dns1 := fmt.Sprintf("push \"dhcp-option DNS %s\"", settings.Client.DNSServer1)
	dns2 := fmt.Sprintf("push \"dhcp-option DNS %s\"", settings.Client.DNSServer2)
	ip := fmt.Sprintf("server %s 255.255.255.0", settings.Server.VpnSubnet)
	port := fmt.Sprintf("port %d", settings.Server.Port)
	gateway := "push \"redirect-gateway def1 bypass-dhcp\""
	newFile = append(newFile, "\n")
	newFile = append(newFile, "#Modifiable Settings")
	newFile = append(newFile, dns1)
	newFile = append(newFile, dns2)
	newFile = append(newFile, ip)
	newFile = append(newFile, port)

	if settings.Client.UseAsGateway {
		newFile = append(newFile, gateway)
	}

	return newFile
}
