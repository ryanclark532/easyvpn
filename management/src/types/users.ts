type User = {
	ID: number;
	Username: string;
	Name: string;
	IsAdmin: boolean;
	Enabled: boolean;
};

export interface GetUsersResponse {
	count: number;
	users: User[];
}
