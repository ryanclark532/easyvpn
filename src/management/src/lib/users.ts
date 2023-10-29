import type { User } from '../types/users';
import { writable } from 'svelte/store';
import { getToken } from './auth';
import { invalidate } from '$app/navigation';

export const searchFilter = writable<string>();
export const selectedUsers = writable<User[]>([]);
export const masterCheckbox = writable<boolean>(false);

export async function setTempPw(u: User[]) {
	const body = {
		ID: u.map((u) => u.id)
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
	if (u.length === 0) {
		return new Error('No Users Selected');
	}
	const headers = new Headers();
	headers.append('JWT', getToken() ?? '');

	const response = await fetch('http://localhost:8080/user', {
		body: JSON.stringify(u),
		method: 'DELETE',
		headers,
		credentials: 'include'
	});

	if (response.status >= 400) {
		return new Error('Error retrieving users, please try again later');
	}

	invalidate('admin:users');
}

export async function updateUser(users: User[]) {
	if (users.length === 0) {
		return new Error('No Users Selected');
	}
	const headers = new Headers();
	headers.append('Authorization', `Bearer ${getToken()}`);
	const response = await fetch('http://localhost:8080/user', {
		body: JSON.stringify(users),
		method: 'PUT',
		headers,
		credentials: 'include'
	});

	if (response.status >= 400) {
		return new Error('Error updating users, please try again later');
	}
	invalidate('admin:users');
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

	const headers = new Headers();
	headers.append('JWT', getToken() ?? '');

	const response = await fetch('http://localhost:8080/user', {
		body: JSON.stringify(body),
		method: 'POST',
		headers,
		credentials: 'include'
	});
	if (response.status === 409) {
		return new Error(`User with username ${body.username} already exists`);
	}
	if (response.status >= 400) {
		return new Error(`Something went wrong while creating ${body.username}`);
	}
	invalidate('admin:users');
}
