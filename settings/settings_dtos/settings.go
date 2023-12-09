package settings_dtos

import "time"

type NetworkSettings struct {
	IPAddress     *string `json:"ip_address"`
	WebServerPort *string `json:"web_server_port"`
}

type VpnSettings struct {
	Port          *int    `json:"port"`
	VpnSubnet     *string `json:"vpn_subnet"`
	VpnSubnetMask *int    `json:"vpn_subnet_mask"`
	PrivateAccess *bool   `json:"private_access"`
	UseAsGateway  *bool   `json:"use_as_gateway"`
	DNSServer1    *string `json:"dnsserver1"`
	DNSServer2    *string `json:"dnsserver2"`
}

type AuthSettings struct {
	AllowChangePW   bool          `json:"allow_change_pw"`
	EnforceStrongPW bool          `json:"enforce_strong_pw"`
	MaxAuthAttempts int           `json:"max_auth_attempts"`
	LockoutTimeout  time.Duration `json:"lockout_timeout"`
}

type Settings struct {
	Network *NetworkSettings `json:"network"`
	Vpn     *VpnSettings     `json:"vpn"`
	Auth    *AuthSettings    `json:"auth"`
}
