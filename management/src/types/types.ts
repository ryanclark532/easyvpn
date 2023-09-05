export interface AuthResponse {
	token: string;
	is_admin: boolean;
}
export interface CheckTokenResponse {
	is_admin: boolean;
	token_valid: boolean;
}

export interface UserCreationResponse {
	user: User;
}

export interface UserResponse {
	users: User[];
	count: number;
}

export interface User {
	Username: string;
	ID: string;
	Name: string;
	IsAdmin: boolean;
	Enabled: boolean;
}
