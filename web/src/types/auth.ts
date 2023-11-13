export interface AuthUser {
	name: string;
	id: string;
	picture: string;
	attrs: {
		admin: boolean;
		password_expiry: string;
		enabled: boolean;
	};
}

export interface AuthError {
	error: string;
}
