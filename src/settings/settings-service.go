package settings

import (
	"bufio"
	"context"
	"easyvpn/src/common"
	"easyvpn/src/database"
	"easyvpn/src/settings/settings_dtos"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)


const SETTINGS_BREAK = "#Modifiable Settings" 

var dnsPattern = `^push "dhcp-option DNS ([0-9.]+)"(.*)$`
var dnsRegex, _ = regexp.Compile(dnsPattern)
var portPattern = `^port (\d+)$`
var portRegex, _ = regexp.Compile(portPattern)
var subnetPattern = `^server (\d+\.\d+\.\d+\.\d+) (\d+\.\d+\.\d+\.\d+)$`
var subnetRegex, _ = regexp.Compile(subnetPattern)

func GetSettings() (*settings_dtos.Settings, error) {
	newSettings := new(settings_dtos.Settings)
	err := database.DB.NewSelect().Model(newSettings).Where("latest = 1").Scan(context.Background())
	if err != nil {
		return nil, err
	}

	file, err := os.Open(common.VPNCONFIG_FILE)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		if line == SETTINGS_BREAK {
			break
		}
	}

	for scanner.Scan() {
		line := scanner.Text()
		readConfigLine(line, newSettings)
	}


	return newSettings, nil
}

func readConfigLine(line string, settings *settings_dtos.Settings) {
		dnsMatch := dnsRegex.FindStringSubmatch(line)
		portMatch := portRegex.FindStringSubmatch(line)
		subnetMatch := subnetRegex.FindStringSubmatch(line)

		if len(dnsMatch) > 0 {
			if settings.DNSServer1 == "" {
				settings.DNSServer1 = dnsMatch[1]
			} else {
				settings.DNSServer2 = dnsMatch[1]
			}
		}
		if len(subnetMatch) > 0 {
			settings.VpnSubnet = subnetMatch[1]

			s, _:= subnetMaskToCIDR(subnetMatch[2])
			settings.VpnSubnetMask = s //TODO add transform to / notation
		}
		if len(portMatch) > 0 {
		portint , _ := strconv.ParseInt(portMatch[1], 0,0)
			settings.Port = int(portint) 
		}
		if line == "push \"redirect-gateway def1 bypass-dhcp\"" {
			settings.UseAsGateway = true
		}
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

func SetSettings(settings *settings_dtos.Settings) error{
	fmt.Println(settings)
	_, err := database.DB.NewUpdate().Table("settings").Set("latest = ?", false).Where("latest = ?", true).Exec(context.Background())
	if err != nil {
		return err
	}
	settings.Version +=1
	settings.Latest = true
	_, err = database.DB.NewInsert().Model(settings).Exec(context.Background())
	if err != nil {
		return err
	}
	err = RewriteConfigFile(settings)
	return err
}

func RewriteConfigFile(settings *settings_dtos.Settings) error {
	file, err:=os.OpenFile(common.VPNCONFIG_FILE, os.O_RDWR,0644)
	if err != nil {
		return err
	}
	defer file.Close()
	
	scanner := bufio.NewScanner(file)

	var modifiedContent []string
	for scanner.Scan() {
		line := scanner.Text()
		modifiedContent = append(modifiedContent, line)
		if line == SETTINGS_BREAK {
			break
		}
	}

	for scanner.Scan() {
		line := scanner.Text()
		line = writeConfigFile(line, settings)	
		modifiedContent = append(modifiedContent, line)
		if settings.UseAsGateway {
			modifiedContent = append(modifiedContent, "push \"redirect-gateway def1 bypass-dhcp\"" )
		}
	}
	if err := scanner.Err(); err != nil {
		return err
	}

	if err := file.Truncate(0); err != nil {
		return err
	}

	if _, err := file.Seek(0, 0); err != nil {
		return err
	}

	for _, line := range modifiedContent {
		fmt.Fprintln(file, line)
	}
	return nil
}

func writeConfigFile(line string, settings *settings_dtos.Settings) string {
		dnsMatch := dnsRegex.FindStringSubmatch(line)
		portMatch := portRegex.FindStringSubmatch(line)
		subnetMatch := subnetRegex.FindStringSubmatch(line)

		if len(dnsMatch) > 0 {
		return fmt.Sprintf("push \"dhcp-option DNS %s\" \n push \"dhcp-option DNS %s\"", settings.DNSServer1, settings.DNSServer2)	
		}
		if len(subnetMatch) > 0 {
		return fmt.Sprintf("server %s %s", settings.VpnSubnet, strconv.Itoa(settings.VpnSubnetMask))
		}
		if len(portMatch) > 0 {
		return fmt.Sprintf("port %s", strconv.Itoa(settings.Port))
		}
	return line
}
