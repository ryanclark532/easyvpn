import { typedFetch } from '$lib/fetch';
import { writable } from 'svelte/store';
import type { AuthResponse, CheckTokenResponse } from '../../types/auth';
import { redirect } from '@sveltejs/kit';

export function getToken(): string | undefined {
	try {
		if (localStorage) {
			return localStorage.getItem('token') ?? undefined;
		}
	} catch (e) {
		return undefined;
	}
}
export function getID(): string | undefined {
	try {
		if (localStorage) {
			return localStorage.getItem('ID') ?? undefined;
		}
	} catch (e) {
		return undefined;
	}
}

export function setToken(token: string) {
	if (localStorage) {
		localStorage.setItem('token', token);
	}
}
export function setID(id: string) {
	if (localStorage) {
		localStorage.setItem('ID', id);
	}
}

export async function createAuthStore() {
	const { subscribe, set, update } = writable<{
		passwordValid: boolean;
		admin: boolean;
		valid: boolean;
	}>();

	set(await checkToken());

	return {
		subscribe,
		set,
		update,
		check: checkToken(),
		handleLogin: (e: Event) => handleLogin(e, set),
		changePassword: (e: Event) => handleChangePassword(e),
		getCurrent: () => getCurrent(subscribe),
		checkRoute: () => checkRoute(getCurrent(subscribe))
	};
}

async function handleChangePassword(e: Event) {
	e.preventDefault();
	const formData = new FormData(e.target as HTMLFormElement);
	const password = formData.get('password') as string;
	const passwordConfirmation = formData.get('confirmPassword') as string;
	const ID = getID();
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
	if (!token) return new Error('Token Not Provided');
	const response = await typedFetch<CheckTokenResponse>(
		'http://localhost:8080/api/auth/check-token',
		{
			body: JSON.stringify({ token: token }),
			method: 'POST'
		}
	);

	if (response instanceof Error || response.status >= 400) {
		return new Error('Error while checking token');
	}

	return response.data;
}

async function handleLogin(
	e: Event,
	set: (
		this: void,
		value: {
			valid: boolean;
			admin: boolean;
			passwordValid: boolean;
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

	if (
		response instanceof Error ||
		response.status >= 400 ||
		response.data?.error ||
		!response.data.token ||
		!response.data.id
	) {
		return new Error(
			'data' in response && response.data.error
				? response.data.error
				: 'Error while processing login'
		);
	}
	setID(response.data.id);
	setToken(response.data.token);
	set(await checkToken());
	console.log('?');
	const location = response.data.password_expired
		? 'user/reset'
		: response.data.is_admin
		? 'admin/status'
		: '/';
	console.log(location);
	redirect(307, location);
}

function validateLoginAttempt(username: string | undefined, password: string | undefined) {
	return username && password;
}

function validatePasswordChangeAttempt(password: string | undefined, confirm: string | undefined) {
	return !confirm || !password;
}

export async function checkToken() {
	const token = getToken();
	if (!token) {
		console.log('exits here');
		return {
			valid: false,
			admin: false,
			passwordValid: false
		};
	}
	const tokenCheck = await validateToken(token);
	if (tokenCheck instanceof Error) {
		console.log('Exits here?');
		return {
			valid: false,
			admin: false,
			passwordValid: false
		};
	}
	return {
		valid: tokenCheck?.is_admin,
		admin: tokenCheck?.is_admin,
		passwordValid: !tokenCheck?.password_expired
	};
}

function checkRoute(check: { passwordValid: boolean; admin: boolean; valid: boolean }) {
	const token = getToken();
	if (!token) {
		console.log('?');
		//throw redirect(307, 'login');
	}

	if (!check.passwordValid && check.valid) {
		//throw redirect(307, '/user/reset');
	}
	if (!check.admin && check.valid) {
		//throw redirect(307, '/user');
	}

	if (!check.valid) {
		//throw redirect(307, 'login');
	}
}

function getCurrent(s: any) {
	let c;

	const unsub = s((e: any) => (c = e));
	unsub();
	return c as unknown as { admin: boolean; valid: boolean; passwordValid: boolean };
}
