export interface AuthResponse {
	token: string;
	is_admin: boolean;
	error?: string;
}

export interface CheckTokenResponse {
	is_admin: boolean;
	token_valid: boolean;
}
