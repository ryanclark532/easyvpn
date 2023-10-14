import type { User } from '../types/users';
import { writable } from 'svelte/store';
import { getToken } from './auth';
import { invalidate } from '$app/navigation';

export const searchFilter = writable<string>();
export const selectedUsers = writable<User[]>([]);
export const masterCheckbox = writable<boolean>(false);

export async function setTempPw(u: User[]) {
	const body = {
		ID: u.map((u) => u.ID)
	};
	if (body.ID.length === 0) {
		return new Error('No Users Selected');
	}
	const headers = new Headers();
	headers.append('JWT', getToken() ?? '');

	const response = await fetch('http://localhost:8080/user/set-temporary-password', {
		body: JSON.stringify(body),
		method: 'PUT',
		headers,
		credentials: 'include'
	});

	if (response.status >= 400) {
		return new Error('Error setting users temporary password, please try again later');
	}
}

export async function deleteUser(u: User[]) {
	const body = {
		ID: u.map((u) => u.ID)
	};
	if (body.ID.length === 0) {
		return new Error('No Users Selected');
	}
	const headers = new Headers();
	headers.append('JWT', getToken() ?? '');

	const response = await fetch('http://localhost:8080/user', {
		body: JSON.stringify(body),
		method: 'DELETE',
		headers
	});

	if (response.status >= 400) {
		return new Error('Error retrieving users, please try again later');
	}

	invalidate('http://localhost:8080/user');
}

export async function updateUser(users: User[]) {
	if (users.length === 0) {
		return new Error('No Users Selected');
	}
	const headers = new Headers();
	headers.append('Authorization', `Bearer ${getToken()}`);
	const response = await fetch('http://localhost:8080/user', {
		body: JSON.stringify({ users }),
		method: 'PUT',
		headers,
		credentials: 'include'
	});

	if (response.status >= 400) {
		return new Error('Error updating users, please try again later');
	}
}
