package settings

import (
	"bufio"
	"context"
	"easyvpn/internal/common"
	"easyvpn/internal/database"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/uptrace/bun"
)

const SETTINGS_BREAK = "#Modifiable Settings"

var dnsPattern = `^push "dhcp-option DNS ([0-9.]+)"(.*)$`
var dnsRegex, _ = regexp.Compile(dnsPattern)
var portPattern = `^port (\d+)$`
var portRegex, _ = regexp.Compile(portPattern)
var subnetPattern = `^server (\d+\.\d+\.\d+\.\d+) (\d+\.\d+\.\d+\.\d+)$`
var subnetRegex, _ = regexp.Compile(subnetPattern)

type Settings struct {
	bun.BaseModel `bun:"table:settings,alias:s"`
	Version       int  `json:"version" bun:",pk,autoincrement"`
	Latest        bool `json:"latest" bun:",notnull"`

	//Auth
	AllowChangePW   bool          `json:"allow_change_pw" bun:",notnull"`
	EnforceStrongPW bool          `json:"enforce_strong_pw" bun:",notnull"`
	MaxAuthAttempts int           `json:"max_auth_attempts" bun:",notnull"`
	LockoutTimeout  time.Duration `json:"lockout_timeout" bun:",notnull"`
	//Server
	VpnSubnet      string `json:"vpn_subnet" bun:"-"`
	VpnSubnetMask  int    `json:"vpn_subnet_mask" bun:"-"`
	TCPPort        int    `json:"port" bun:"-"`
	UDPPort        int    `json:"port" bun:"-"`
	IPAddress      string `json:"ip_address" bun:",notnull"`
	WebServerPort  int    `json:"web_server_port" bun:",notnull"`
	MaxConnections int    `json:"max_connections" bun:",notnull"`

	//Client
	UseAsGateway  bool   `json:"use_as_gateway" bun:"-"`
	PrivateAccess bool   `json:"private_access" bun:"-"`
	DNSServer1    string `json:"dnsserver1" bun:"-"`
	DNSServer2    string `json:"dnsserver2" bun:"-"`
}

func (s Settings) SaveUDPSettings() error {
	_, err := database.DB.NewUpdate().Table("settings").Set("latest = ?", false).Where("latest = ?", true).Exec(context.Background())
	if err != nil {
		return err
	}
	s.Version += 1
	s.Latest = true
	_, err = database.DB.NewInsert().Model(s).Exec(context.Background())
	if err != nil {
		return err
	}
	err = s.RewriteConfig(common.VPN_UDP_CONFIG_FILE)
	return err
}

func (s Settings) SaveTCPSettings() error {
	_, err := database.DB.NewUpdate().Table("settings").Set("latest = ?", false).Where("latest = ?", true).Exec(context.Background())
	if err != nil {
		return err
	}
	s.Version += 1
	s.Latest = true
	_, err = database.DB.NewInsert().Model(s).Exec(context.Background())
	if err != nil {
		return err
	}
	err = s.RewriteConfig(common.VPN_TCP_CONFIG_FILE)
	return err
}

func (s Settings) RewriteConfig(config string) error {

	file, err := os.OpenFile(config, os.O_RDWR, 0644)
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

	vpnMask, err := cidrToSubnetMask(s.VpnSubnetMask)
	if err != nil {
		return err
	}

	modifiableSettings := [4]string{fmt.Sprintf("push \"dhcp-option DNS %s\" \npush \"dhcp-option DNS %s\"", s.DNSServer1, s.DNSServer2), fmt.Sprintf("server %s %s", s.VpnSubnet, vpnMask), fmt.Sprintf("port %s", strconv.Itoa(s.TCPPort))}
	if s.PrivateAccess {
		modifiableSettings[3] = "push \"redirect-gateway def1 bypass-dhcp\""
	}
	contentLen := len(modifiedContent)
	modifiedContent = append(modifiedContent, make([]string, 4)...)

	copy(modifiedContent[contentLen:], modifiableSettings[:])
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

func GetTCPSettings() (*Settings, error) {
	newSettings := new(Settings)
	err := database.DB.NewSelect().Model(newSettings).Where("latest = 1").Scan(context.Background())
	if err != nil {
		return nil, err
	}

	file, err := os.Open(common.VPN_TCP_CONFIG_FILE)
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
	fmt.Println(newSettings.VpnSubnet)
	return newSettings, nil
}

func GetUDPSettings() (*Settings, error) {
	newSettings := new(Settings)
	err := database.DB.NewSelect().Model(newSettings).Where("latest = 1").Scan(context.Background())
	if err != nil {
		return nil, err
	}

	file, err := os.Open(common.VPN_UDP_CONFIG_FILE)
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
	fmt.Println(newSettings.VpnSubnet)
	return newSettings, nil
}

func readConfigLine(line string, settings *Settings) {
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
		fmt.Println(subnetMatch)
		settings.VpnSubnet = subnetMatch[1]

		s, _ := subnetMaskToCIDR(subnetMatch[2])
		settings.VpnSubnetMask = s
	}
	if len(portMatch) > 0 {
		portint, _ := strconv.ParseInt(portMatch[1], 0, 0)
		settings.TCPPort = int(portint)
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

func cidrToSubnetMask(cidr int) (string, error) {
	if cidr < 0 || cidr > 32 {
		return "", fmt.Errorf("invalid CIDR value")
	}

	binaryRepresentation := strings.Repeat("1", cidr) + strings.Repeat("0", 32-cidr)

	var octets []string
	for i := 0; i < 4; i++ {
		start := i * 8
		end := start + 8
		octets = append(octets, binaryRepresentation[start:end])
	}

	var decimalOctets []string
	for _, octet := range octets {
		decimalValue, _ := strconv.ParseInt(octet, 2, 64)
		decimalOctets = append(decimalOctets, fmt.Sprintf("%d", decimalValue))
	}

	subnetMask := strings.Join(decimalOctets, ".")
	return subnetMask, nil
}
