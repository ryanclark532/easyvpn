export type ClientSettings = {
	dnsserver1: string;
	dnsserver2: string;
	private_access: boolean;
	use_as_gateway: boolean;
};

export type ServerSettings = {
	vpn_subnet: string;
	vpn_subnet_mask: number;
	port: number;
	web_server_port: number;
	ip_address: string;
};

export type AuthSettings = {
	allow_change_pw: boolean;
	enforce_strong_pw: boolean;
	max_auth_attempts: number;
	lockout_timeout: number;
};

export type Settings = {
	client: ClientSettings;
	server: ServerSettings;
	auth: AuthSettings;
};
