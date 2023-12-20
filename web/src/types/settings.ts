export type Settings = {
	//Auth
	enforce_strong_pw: boolean;
	allow_change_pw: boolean;
	max_auth_attempts: number;
	lockout_timeout: number;

	//Server
	vpn_subnet: string;
	vpn_subnet_mask: number;
	port: number;
	web_server_port: number;
	ip_address: string;

	//Client
	dnsserver1: string;
	dnsserver2: string;
	private_access: boolean;
	use_as_gateway: boolean;
};
