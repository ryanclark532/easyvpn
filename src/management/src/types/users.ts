export type User = {
	ID: number;
	Username: string;
	Name: string;
	IsAdmin: boolean;
	Enabled: boolean;
};

export type GetUsersResponse = {
	users: User[];
};

export type CreateUserRequest = {
	username: string;
	name: string;
	password: string;
	is_admin: boolean;
	enabled: boolean;
};

export type UserCreationResponse = {
	user: User;
};

export type UpdateUserResponse = {
	users: User[];
};
