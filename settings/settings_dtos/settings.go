package settings_dtos

import (
	"time"

	"github.com/uptrace/bun"
)

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
	VpnSubnet     string `json:"vpn_subnet" bun:"-"`
	VpnSubnetMask int    `json:"vpn_subnet_mask" bun:"-"`
	Port          int    `json:"port" bun:"-"`
	IPAddress     string `json:"ip_address" bun:",notnull"`
	WebServerPort int    `json:"web_server_port" bun:",notnull"`

	//Client
	UseAsGateway  bool   `json:"use_as_gateway" bun:"-"`
	PrivateAccess bool   `json:"private_access" bun:"-"`
	DNSServer1    string `json:"dnsserver1" bun:"-"`
	DNSServer2    string `json:"dnsserver2" bun:"-"`
}
