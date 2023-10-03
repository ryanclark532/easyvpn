import { getID, getToken, setID, setToken } from '$lib/auth';
import { typedFetch } from '$lib/fetch';
import { writable } from 'svelte/store';
import type { AuthResponse, CheckTokenResponse } from '../../types/auth';
import { redirect } from '@sveltejs/kit';

export function createAuthStore() {
	const { subscribe, set, update } = writable<{ token: string; id: string }>();

	const token = getToken();
	const id = getID();
	if (token && id) {
		set({
			token,
			id
		});
	}

	return {
		subscribe,
		set,
		update,
		validate: () => validateToken(token),
		handleLogin: (e: Event) => handleLogin(e, set),
		changePassword: (e: Event) => handleChangePassword(e, id, set)
	};
}

async function handleChangePassword(
	e: Event,
	ID: string | undefined,
	set: (
		this: void,
		value: {
			token: string;
			id: string;
		}
	) => void
) {
	e.preventDefault();
	const formData = new FormData(e.target as HTMLFormElement);
	const password = formData.get('password') as string;
	const passwordConfirmation = formData.get('confirmPassword') as string;

	if (!validatePasswordChangeAttempt(password, passwordConfirmation)) {
		return new Error('Invalid Password Change Attempt');
	}
	const token = getToken();
	const headers = new Headers();
	headers.append('Authorization', `Bearer ${token}`);
	const response = await fetch('http://localhost:8080/api/auth/change-password', {
		body: JSON.stringify({
			password,
			ID
		}),
		method: 'POST',
		headers
	});

	if (response.status >= 400) {
		return new Error('Error Changing password');
	}
	setID('');
	setToken('');

	redirect(307, '/login');
}

async function validateToken(token: string | undefined) {
	if (!token) return;

	const response = await typedFetch<CheckTokenResponse>(
		'http://localhost:8080/api/auth/check-token',
		{
			body: JSON.stringify(token),
			method: 'POST'
		}
	);

	if (response.status >= 400) {
		return new Error('Error while checking token');
	}

	return response.data;
}

async function handleLogin(
	e: Event,
	set: (
		this: void,
		value: {
			token: string;
			id: string;
		}
	) => void
) {
	e.preventDefault();
	const formData = new FormData(e.target as HTMLFormElement);
	const username = formData.get('username') as string | undefined;
	const password = formData.get('password') as string | undefined;
	if (!validateLoginAttempt(username, password)) {
		return new Error('Invalid login attempt');
	}
	const response = await typedFetch<AuthResponse>('http://localhost:8080/api/auth/sign-in', {
		body: JSON.stringify({ username, password }),
		method: 'POST'
	});

	if (response.status >= 400 || response.data.error || !response.data.token || !response.data.id) {
		return new Error(response.data.error ? response.data.error : 'Error while processing login');
	}
	set({
		token: response.data.token,
		id: response.data.id
	});
	redirect(
		307,
		response.data.password_expired ? 'user/reset' : response.data.is_admin ? 'admin/status' : '/'
	);
}

function validateLoginAttempt(username: string | undefined, password: string | undefined) {
	return username && password;
}

function validatePasswordChangeAttempt(password: string | undefined, confirm: string | undefined) {
	return !confirm || !password;
}
