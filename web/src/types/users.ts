export type User = {
	id: number;
	username: string;
	name: string;
	is_admin: boolean;
	enabled: boolean;
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
