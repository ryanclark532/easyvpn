package settings_dtos

import "time"

type Settings struct {
	Client *ClientSettings `json:"network"`
	Server *ServerSettings `json:"vpn"`
	Auth   *AuthSettings   `json:auth"`
}

type AuthSettings struct {
	AllowChangePW   bool          `json:"allow_change_pw"`
	EnforceStrongPW bool          `json:"enforce_strong_pw"`
	MaxAuthAttempts int           `json:"max_auth_attempts"`
	LockoutTimeout  time.Duration `json:"lockout_timeout"`
}

type ServerSettings struct {
	VpnSubnet     string `json:"vpn_subnet"`
	VpnSubnetMask int    `json:"vpn_subnet_mask"`
	Port          int    `json:"port"`
	IPAddress     string `json:"ip_address"`
	WebServerPort string `json:"web_server_port"`
	PrivateAccess bool   `json:"private_access"`
	UseAsGateway  bool   `json:"use_as_gateway"`
}

type ClientSettings struct {
	DNSServer1 *string `json:"dnsserver1"`
	DNSServer2 *string `json:"dnsserver2"`
}
