export type Settings = {
	network: NetworkSettings;
	vpn: VpnSettings;
	auth: AuthSettings;
};

export type NetworkSettings = {
	ip_address: string;
	protocol: string;
	web_server_port: string;
};

export type VpnSettings = {
	vpn_subnet: string;
	vpn_subnet_mask: number;
	private_access: boolean;
	use_as_gateway: boolean;
	dnsserver1: string;
	dnsserver2: string;
};

export type AuthSettings = {
	allow_change_pw: boolean;
	enforce_strong_pw: boolean;
	max_auth_attempts: number;
	lockout_timeout: number;
};
