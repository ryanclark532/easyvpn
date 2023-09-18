import {writable} from 'svelte/store';
import type {DataWithStatus} from '../types/data-with-status';
import {typedFetch} from '$lib/fetch';
import type {AuthResponse, CheckTokenResponse} from '../types/auth';

export const loginResponse = writable<DataWithStatus<string | AuthResponse>>({
	data: undefined,
	status: 'initial'
});

export const passwordChangeResponse = writable<DataWithStatus<any>>({
	data: undefined,
	status: 'initial'
});

export function getToken(): string | undefined {
	if (localStorage) {
		return localStorage.getItem('token') ?? undefined;
	}
	console.log('localstorage undefined');
}
export function getID(): string | undefined {
	if (localStorage) {
		return localStorage.getItem('ID') ?? undefined;
	}
	console.log('localstorage undefined');
}

function setToken(token: string) {
	if (localStorage) {
		localStorage.setItem('token', token);
	}
}
function setID(id: string) {
	if (localStorage) {
		localStorage.setItem('ID', id);
	}
}

export async function handleLogin(e: Event) {
	loginResponse.set({
		status: 'loading',
		data: undefined
	});
	e.preventDefault();
	const formData = new FormData(e.target as HTMLFormElement);
	const username = formData.get('username') as string;
	const password = formData.get('password') as string;

	if (!validateLoginAttempt(username, password)) {
		loginResponse.set({
			status: 'error',
			data: 'Invalid username or password'
		});
		return;
	}

	const response = await typedFetch<AuthResponse>('http://localhost:8080/user/sign-in', {
		body: JSON.stringify({ username, password }),
		method: 'POST'
	});
	if (response.status >= 400 || response.data.error) {
		loginResponse.set({
			status: 'error',
			data: response.data.error ?? 'Something Went Wrong, Please Try Again'
		});
		return;
	}

	if (response.data.token && response.data.id) {
		setID(response.data.id);
		setToken(response.data.token);

		window.location.href = response.data.password_expired
			? '/user/reset'
			: response.data.is_admin
			? '/admin/status'
			: '/';
	}
}

export async function changePassword(e: Event) {
	passwordChangeResponse.set({
		status: 'loading',
		data: undefined
	});
	e.preventDefault();
	const formData = new FormData(e.target as HTMLFormElement);
	const password = formData.get('password') as string;
	const passwordConfirmation = formData.get('confirmPassword') as string;

	if (!validatePasswordChangeAttempt(password, passwordConfirmation)) {
		passwordChangeResponse.set({
			status: 'error',
			data: 'Please ensure that passwords match'
		});
		return;
	}
	const token = getToken();
	const headers = new Headers();
	headers.append('Authorization', `Bearer ${token}`);
	const response = await fetch('http://localhost:8080/user/change-password', {
		body: JSON.stringify({
			password,
			ID: getID()
		}),
		method: 'POST',
		headers
	});

	if (response.status >= 400) {
		passwordChangeResponse.set({
			status: 'error',
			data: 'There was an error, please try again later'
		});
		return;
	}

	passwordChangeResponse.set({
		status: 'ready',
		data: 'Password changed successfully'
	});
	setID('');
	setToken('');
	window.location.href = '/login';
}

export async function getTokenValid(
	token: string
): Promise<DataWithStatus<CheckTokenResponse | undefined>> {
	const body = {
		token
	};

	const response = await typedFetch<CheckTokenResponse>('http://localhost:8080/user/check-token', {
		body: JSON.stringify(body),
		method: 'POST'
	});
	if (response.status >= 400) {
		return {
			data: undefined,
			status: 'error'
		};
	}

	return {
		data: response.data,
		status: 'ready'
	};
}

function validatePasswordChangeAttempt(
	password: string | undefined,
	passwordConfirmation: string | undefined
) {
	if (!password || !passwordConfirmation) {
		return false;
	}
	return passwordConfirmation === password;
}

function validateLoginAttempt(username: string | undefined, password: string | undefined) {
	//TODO add some validations here
	return true;
}

export async function checkToken(): Promise<{
	valid: boolean;
	admin: boolean;
	passwordValid: boolean;
}> {
	const token = getToken();
	if (!token) {
		return {
			valid: false,
			admin: false,
			passwordValid: false
		};
	}
	const tokenCheck = await getTokenValid(token);
	if (!tokenCheck || tokenCheck.status === 'error' || !tokenCheck.data) {
		return {
			valid: false,
			admin: false,
			passwordValid: false
		};
	}
	return {
		valid: tokenCheck.data.token_valid,
		admin: tokenCheck.data.is_admin,
		passwordValid: !tokenCheck.data.password_expired
	};
}
