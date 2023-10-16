export interface AuthUser {
	name: string;
	id: string;
	picture: string;
	attrs: {
		admin: boolean;
		password_expiry: string;
	};
}

export interface AuthError {
	error: string;
}
