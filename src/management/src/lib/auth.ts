import { redirect } from '@sveltejs/kit';
import type { AuthError, AuthUser } from '../types/auth';

export function getToken(){
	return localStorage ? localStorage.getItem("JWT") : undefined;
}



export async function handleLogin(e: Event) {
	e.preventDefault();

	const formData = new FormData(e.target as HTMLFormElement);
	const username = formData.get('username');
	const password = formData.get('password');

	const response = await fetch('http://localhost:8080/auth/local/login', {
		body: JSON.stringify({
			username,
			password
		}),
		method: 'POST',
		credentials: 'include'
	});
	if (!response.ok) {
		return new Error('Error while processing login');
	}
	localStorage.setItem('jwt', response.headers.get('jwt') ?? '');
}

export function handleRedirects(input: AuthUser | AuthError) {
	if (isAuthError(input)) {
		throw redirect(307, '/login');
	}
	if(!input.attrs.admin){
		throw redirect(307,'/');
	}
}

function isAuthUser(input: AuthUser | AuthError): input is AuthUser {
	return 'name' in input;
}
function isAuthError(input: AuthUser | AuthError): input is AuthError {
	return 'error' in input;
}
