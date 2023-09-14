import { writable } from 'svelte/store';
import type { DataWithStatus } from '../types/data-with-status';
import { typedFetch } from '$lib/fetch';
import type { AuthResponse, CheckTokenResponse } from '../types/auth';
import { redirect } from '@sveltejs/kit';

export const loginResponse = writable<DataWithStatus<string | AuthResponse>>({
	data: undefined,
	status: 'initial'
});

export function getToken(): string | undefined {
	if (localStorage) {
		return localStorage.getItem('token') ?? undefined;
	}
	console.log('localstorage undefined');
}

function setToken(token: string) {
	if (localStorage) {
		localStorage.setItem('token', token);
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

	if (response.data.token) {
		setToken(response.data.token);
		window.location.href = response.data.is_admin ? '/admin/status' : '/';
	}
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

function validateLoginAttempt(username: string | undefined, password: string | undefined) {
	//TODO add some validations here
	return true;
}

export async function checkToken() {
	const token = getToken();
	if (!token) {
		throw redirect(307, '/login');
	}
	const tokenCheck = await getTokenValid(token);
	if (!tokenCheck || tokenCheck.status === 'error' || !tokenCheck.data?.token_valid) {
		throw redirect(307, '/login');
	}
	if (!tokenCheck.data.is_admin) {
		throw redirect(307, '/');
	}
	return token;
}
