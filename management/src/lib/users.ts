import { typedFetch } from '$lib/fetch';
import { getToken } from '$lib/auth';
import type { DataWithStatus } from '../types/data-with-status';
import type {
	CreateUserRequest,
	GetUsersResponse,
	User,
	UserCreationResponse
} from '../types/users';
import { writable } from 'svelte/store';

export const users = writable<User[]>([]);
export const createUserResponse = writable<DataWithStatus<string | undefined>>({
	status: "initial",
	data: undefined,
});

export async function getUsers(
	token: string
): Promise<DataWithStatus<GetUsersResponse | undefined>> {
	const headers = new Headers();
	headers.append('Authorization', `Bearer ${token}`);
	const response = await typedFetch<GetUsersResponse>('http://localhost:8080/admin/get-users', {
		method: 'GET',
		headers
	});

	if (response.status >= 400) {
		return {
			status: 'error',
			data: undefined
		};
	}
	return {
		status: 'ready',
		data: response.data
	};
}

export async function createUser(e: Event) {
	e.preventDefault();
	const formData = new FormData(e.target as HTMLFormElement);
	const body = {
		username: formData.get('username') as string,
		name: formData.get('name') as string,
		password: formData.get('password') as string,
		is_admin: (formData.get('isAdmin') as string) === 'on',
		enabled: (formData.get('enabled') as string) === 'on'
	};

	if (!validateNewUser(body)) {
		createUserResponse.set({
			status: 'error',
			data: 'Please fill out all fields'
		});
	}
	const token = getToken();
	const headers = new Headers();
	headers.append('Authorization', `Bearer ${token}`);

	const response = await typedFetch<UserCreationResponse>(
		'http://localhost:8080/admin/create-user',
		{
			body: JSON.stringify(body),
			method: 'POST',
			headers
		}
	);
	if (response.status === 409) {
		createUserResponse.set({
			status: 'error',
			data: 'User already exists'
		});
	}
	if (response.status >= 400) {
		createUserResponse.set({
			status: 'error',
			data: 'Something went wrong. Please try again later.'
		});
	}
	createUserResponse.set({
		status: 'ready',
		data: `User ${body.username} created sucessfully`
	});
	users.update((users) => [...users, response.data.user]);
}

function validateNewUser(user: CreateUserRequest) {
	//todo some validation here
	return true;
}
