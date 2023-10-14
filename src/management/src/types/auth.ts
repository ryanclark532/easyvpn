export interface AuthUser {
	name: string;
	id: string;
	picture: string;
	attrs: {
		admin: boolean;
	};
}

export interface AuthError {
	error: string;
}
