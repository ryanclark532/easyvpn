import { handleRedirects } from '$lib/auth';

export async function load({ fetch, cookies }) {
	const headers = new Headers();
	headers.append('JWT', cookies.get('JWT') ?? '');
	const authcheck = await fetch('http://localhost:8080/auth/user', {
		headers,
		credentials: 'include'
	}).then((response) => response.json());
	console.log(authcheck);
	handleRedirects(authcheck);
}
