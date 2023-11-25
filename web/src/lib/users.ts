import type { User } from '../types/users';
import { writable } from 'svelte/store';
import { getToken } from './auth';
import { invalidate } from '$app/navigation';

export const userFilter = writable<string>();
export const selectedUsers = writable<User[]>([]);
export const masterCheckbox = writable<boolean>(false);

export async function deleteUser(userId: number) {
	const headers = new Headers();
	headers.append('JWT', getToken() ?? '');
	const response = await fetch(`http://localhost:8080/user/${userId}`, {
		method: 'DELETE',
		headers,
		credentials: 'include'
	});

	if (response.status >= 400) {
		return new Error('Error updating users, please try again later');
	}
	invalidate('admin:users');
	selectedUsers.set([]);
}

export async function updateUser(user: User) {
	const headers = new Headers();
	headers.append('JWT', getToken() ?? '');
	const response = await fetch(`http://localhost:8080/user/${user.id}`, {
		body: JSON.stringify(user),
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
		username: formData.get('username').toString(),
		name: formData.get('name').toString(),
		password: formData.get('password').toString(),
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

export async function changePassword(e: Event, userId: number, temp: bool) {
	e.preventDefault();
	const formData = new FormData(e.target as HTMLFormElement);
	const password = formData.get('password').toString();
	const confirm = formData.get('confirm').toString();
	if (password !== confirm) {
		return new Error('Please enter matching passwords');
	}

	const headers = new Headers();
	headers.append('JWT', getToken() ?? '');

	const response = await fetch(`http://localhost:8080/user/${userId}/set-pw`, {
		body: JSON.stringify({ password, confirm, temp }),
		method: 'POST',
		headers,
		credentials: 'include'
	}).then((res) => res.status);

	if (response >= 400) {
		return new Error('Something went wrong while changing the password. Please try again');
	}
}
