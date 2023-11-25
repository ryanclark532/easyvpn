import { redirect } from '@sveltejs/kit';
import type { AuthError, AuthUser } from '../types/auth';
import { goto } from '$app/navigation';

export function getToken() {
	return localStorage ? localStorage.getItem('jwt') : undefined;
}

export function getID() {
	return localStorage ? localStorage.getItem('id') : undefined;
}
export async function handleLogin(e: Event) {
	e.preventDefault();
	const formData = new FormData(e.target as HTMLFormElement);
	const username = formData.get('username')?.toString();
	const password = formData.get('password')?.toString();

	const headers = new Headers();
	headers.append('Content-Type', 'application/json');
	return await fetch('http://localhost:8080/auth/local/login', {
		body: JSON.stringify({
			user: username,
			passwd: password,
			aud: username
		}),
		headers,
		method: 'POST',
		credentials: 'include'
	}).then(async (response) => {
		if (!response.ok) {
			return new Error('Error while processing login');
		}

		localStorage.setItem('jwt', response.headers.get('jwt') ?? '');
		document.cookie = `JWT=${response.headers.get('jwt') ?? ''}`;

		const json = await response.json();
		localStorage.setItem('id', json.attrs.id);
		const expiry = new Date(json.attrs.password_expiry);
		if (expiry < new Date()) {
			goto('/login/reset');
		}
		goto(json.attrs.admin ? '/admin/status' : '/');
	});
}

export function handleRedirects(input: AuthUser | AuthError) {
	if (isAuthError(input)) {
		throw redirect(307, '/login');
	}
	if (new Date(input.attrs.password_expiry) < new Date()) {
		throw redirect(307, '/login/reset');
	}

	if (!input.attrs.admin) {
		throw redirect(307, '/');
	}
}

function isAuthError(input: AuthUser | AuthError): input is AuthError {
	return 'error' in input;
}
