import { typedFetch } from '$lib/fetch';
import { getToken } from '$lib/auth';
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
export const createUserResponse = writable<DataWithStatus<string | undefined>>({
	status: 'initial',
	data: undefined
});
export const actionResponse = writable<DataWithStatus<string | undefined>>();
export const searchFilter = writable<string>();
export const selectedUsers = writable<User[]>([]);
export const masterCheckbox = writable<boolean>(false);

export async function getUsers(
	token: string | undefined
): Promise<DataWithStatus<GetUsersResponse | undefined>> {
	const headers = new Headers();
	headers.append('Authorization', `Bearer ${token}`);
	const response = await typedFetch<GetUsersResponse>('http://localhost:8080/api/user', {
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

	const response = await typedFetch<UserCreationResponse>('http://localhost:8080/api/user', {
		body: JSON.stringify(body),
		method: 'POST',
		headers
	});
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

export async function deleteUsers(usersToDelete: User[]) {
	const body = {
		ID: usersToDelete.map((u) => u.ID)
	};
	if (body.ID.length === 0) {
		actionResponse.set({
			status: 'error',
			data: 'There was an error deleting users. Please try again later.'
		});
		return;
	}
	const token = getToken();
	const headers = new Headers();
	headers.append('Authorization', `Bearer ${token}`);

	const response = await fetch('http://localhost:8080/api/user', {
		body: JSON.stringify(body),
		method: 'DELETE',
		headers
	});

	if (response.status >= 400) {
		actionResponse.set({
			status: 'error',
			data: 'There was an error deleting users. Please try again later.'
		});
	}

	users.update((prev) => {
		return prev.filter((u) => {
			const x = usersToDelete.map((u) => u.ID);
			return !x.includes(u.ID);
		});
	});

	actionResponse.set({
		status: 'ready',
		data: 'Users deleted successfully'
	});
}

export async function setTemporaryPassword(usersToSet: User[]) {
	const body = {
		ID: usersToSet.map((u) => u.ID)
	};
	if (body.ID.length === 0) {
		actionResponse.set({
			status: 'error',
			data: 'There was an error setting a temporary password for the selected users. Please try again later.'
		});
		return;
	}
	const token = getToken();
	const headers = new Headers();
	headers.append('Authorization', `Bearer ${token}`);

	const response = await fetch('http://localhost:8080/api/user/set-temporary-password', {
		body: JSON.stringify(body),
		method: 'PUT',
		headers
	});

	if (response.status >= 400) {
		actionResponse.set({
			status: 'error',
			data: 'There was an error setting a temporary password for the selected users. Please try again later.'
		});
	}

	actionResponse.set({
		status: 'ready',
		data: 'Users updated successfully'
	});
}

export async function updateUser(usersToSet: User[]) {
	const body = {
		users: usersToSet
	};
	if (body.users.length === 0) {
		actionResponse.set({
			status: 'error',
			data: 'There was an error updating the selected users. Please try again later.'
		});
		return;
	}
	const token = getToken();
	const headers = new Headers();
	headers.append('Authorization', `Bearer ${token}`);

	const response = await typedFetch<UpdateUserResponse>('http://localhost:8080/api/user', {
		body: JSON.stringify(body),
		method: 'PUT',
		headers
	});

	if (response.status >= 400) {
		actionResponse.set({
			status: 'error',
			data: 'There was an error updating the selected users. Please try again later.'
		});
		return;
	}

	actionResponse.set({
		status: 'ready',
		data: 'Users updated successfully'
	});

	window.location.reload();
}

function validateNewUser(user: CreateUserRequest) {
	//todo some validation here
	return true;
}
