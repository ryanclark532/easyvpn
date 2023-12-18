package settings

import (
	"bufio"
	"context"
	"easyvpn/database"
	"easyvpn/settings/settings_dtos"
	"easyvpn/utils"
	"easyvpn/vpn"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var dnsPattern = `^push "dhcp-option DNS ([0-9.]+)"(.*)$`
var dnsRegex, _ = regexp.Compile(dnsPattern)
var portPattern = `^port (\d+)$`
var portRegex, _ = regexp.Compile(portPattern)
var subnetPattern = `^server (\d+\.\d+\.\d+\.\d+) (\d+\.\d+\.\d+\.\d+)$`
var subnetRegex, _ = regexp.Compile(subnetPattern)

func GetSettings() (*settings_dtos.Settings, error) {

	file, err := os.Open(`C:\Program Files\OpenVPN\config-auto\server-dev.ovpn`)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var dns1 string
	var dns2 string
	var subnet string
	var subnetMask string
	var port int64
	var gateway bool

	for scanner.Scan() {
		line := scanner.Text()
		if line == "#Modifiable Settings" {
			break
		}
	}

	for scanner.Scan() {
		line := scanner.Text()

		dnsMatch := dnsRegex.FindStringSubmatch(line)
		portMatch := portRegex.FindStringSubmatch(line)
		subnetMatch := subnetRegex.FindStringSubmatch(line)

		if len(dnsMatch) > 0 {
			if dns1 == "" {
				dns1 = dnsMatch[1]
			} else {
				dns2 = dnsMatch[1]
			}
		}
		if len(subnetMatch) > 0 {
			subnet = subnetMatch[1]
			subnetMask = subnetMatch[2] //TODO add transform to / notation
		}
		if len(portMatch) > 0 {
			port, _ = strconv.ParseInt(portMatch[1], 0, 0)
		}
		if line == "push \"redirect-gateway def1 bypass-dhcp\"" {
			gateway = true
		}
	}

	newSettings := new(settings_dtos.Settings)
	err = database.DB.NewSelect().Model(newSettings).Where("latest = 1").Scan(context.Background())
	if err != nil {
		return nil, err
	}
	s, err := subnetMaskToCIDR(subnetMask)
	if err != nil {
		return nil, err
	}
	newSettings.DNSServer1 = dns1
	newSettings.DNSServer2 = dns2
	newSettings.VpnSubnet = subnet
	newSettings.VpnSubnetMask = s
	newSettings.Port = int(port)
	newSettings.UseAsGateway = gateway

	return newSettings, nil
}

func SetSettings(settings *settings_dtos.Settings) error {
	err := RewriteVPNConfig()
	if err != nil {
		return err
	}

	err = database.DB.NewUpdate().Set("latest = ?", false).Where("latest != ?", true).Scan(context.Background())
	if err != nil {
		return err
	}
	settings.Latest = true
	err = database.DB.NewInsert().Model(settings).Scan(context.Background())
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

	settings := new(settings_dtos.Settings)
	err = database.DB.NewSelect().Model(settings).Scan(context.Background())
	if err != nil {
		return err
	}
	newFile = AppendModifyableSettings(newFile, *settings)

	err = utils.WriteFile(`C:\Program Files\OpenVPN\config-auto\server-dev.ovpn`, []byte(strings.Join(newFile, "\n")))
	return err
}

func AppendModifyableSettings(newFile []string, settings settings_dtos.Settings) []string {
	dns1 := fmt.Sprintf("#DNS1 \n push \"dhcp-option DNS %s\"", settings.DNSServer1)
	dns2 := fmt.Sprintf("#DNS2 \n push \"dhcp-option DNS %s\"", settings.DNSServer2)
	ip := fmt.Sprintf("server %s 255.255.255.0", settings.VpnSubnet)
	port := fmt.Sprintf("port %d", settings.Port)
	gateway := "push \"redirect-gateway def1 bypass-dhcp\""
	newFile = append(newFile, "\n")
	newFile = append(newFile, "#Modifiable Settings")
	newFile = append(newFile, dns1)
	newFile = append(newFile, dns2)
	newFile = append(newFile, ip)
	newFile = append(newFile, port)

	if settings.UseAsGateway {
		newFile = append(newFile, gateway)
	}

	return newFile
}
func subnetMaskToCIDR(subnetMask string) (int, error) {
	parts := strings.Split(subnetMask, ".")
	if len(parts) != 4 {
		return 0, fmt.Errorf("invalid subnet mask format")
	}
	var binaryRepresentation string
	for _, part := range parts {
		octet, err := strconv.Atoi(part)
		if err != nil || octet < 0 || octet > 255 {
			return 0, fmt.Errorf("invalid octet in subnet mask")
		}

		binaryRepresentation += fmt.Sprintf("%08b", octet)
	}
	onesCount := strings.Count(binaryRepresentation, "1")
	return onesCount, nil
}
