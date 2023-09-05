import { authResponseMessage } from '../stores/stores';
import { typedFetch } from '$lib/fetch';
import type { AuthResponse, CheckTokenResponse } from '../types/types';

function getToken(): { token: string | undefined; is_admin: boolean | undefined } {
	if (localStorage) {
		return {
			token: localStorage.getItem('token') ?? undefined,
			is_admin: localStorage.getItem('is_admin') === 'true'
		};
	}
	return {
		token: undefined,
		is_admin: undefined
	};
}

function setToken(token: string, is_admin: boolean) {
	if (localStorage) {
		localStorage.setItem('token', token);
		localStorage.setItem('is_admin', `${is_admin}`);
	}
}

export async function handleLogin(e: any): Promise<boolean> {
	const formData = new FormData(e.target);
	const body: { [key: string]: string } = {};

	formData.forEach((value, key) => {
		body[key] = value.toString();
	});

	if (!body['username'] || !body['password']) {
		authResponseMessage.set('Username Or Password Not Provided');
		return false;
	}

	const response = await typedFetch<AuthResponse>('http://localhost:8080/user/sign-in', {
		body: JSON.stringify(body),
		method: 'POST'
	});
	if (response.status >= 500) {
		authResponseMessage.set('Internal Server Error, Please try again');
		return false;
	}
	if (response.status >= 400) {
		authResponseMessage.set('Incorrect Username Or Password');
		return false;
	}

	if (!response.data.token && !response.data.is_admin) {
		return false;
	}

	setToken(response.data.token, response.data.is_admin);
	return true;
}

export async function isAuthed(): Promise<boolean> {
	const { token } = getToken();
	if (!token) {
		return false;
	}
	const body = {
		token
	};

	const response = await typedFetch<CheckTokenResponse>('http://localhost:8080/user/check-token', {
		body: JSON.stringify(body),
		method: 'POST'
	});

	return response.data.token_valid;
}

export async function isAuthedAdmin(): Promise<boolean> {
	const { token } = getToken();
	if (!token) {
		return false;
	}
	const body = {
		token
	};

	const response = await typedFetch<CheckTokenResponse>('http://localhost:8080/user/check-token', {
		body: JSON.stringify(body),
		method: 'POST'
	});

	return response.data.is_admin && response.data.token_valid;
}
