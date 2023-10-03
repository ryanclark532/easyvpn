import type { DataWithStatus } from '../types/data-with-status';
import type {
	CreateUserRequest,
	GetUsersResponse,
	UpdateUserResponse,
	User,
	UserCreationResponse
} from '../types/users';
import { writable } from 'svelte/store';

export const users = writable<User[]>([]);
export const createUserResponse = writable<DataWithStatus<string>>({
	status: 'initial',
	data: undefined
});
export const actionResponse = writable<DataWithStatus<string>>();
export const searchFilter = writable<string>();
export const selectedUsers = writable<User[]>([]);
export const masterCheckbox = writable<boolean>(false);
