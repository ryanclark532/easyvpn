import { handleRedirects } from '$lib/auth';
import type { User } from '../../../types/users';

export async function load({ fetch, cookies, depends }) {
	const headers = new Headers();
	headers.append('JWT', cookies.get('JWT') ?? '');
	const authcheck = await fetch('http://localhost:8080/auth/user', {
		headers,
		credentials: 'include'
	}).then((response) => response.json());	
	handleRedirects(authcheck);

	const users = await fetch('http://localhost:8080/user', {
		headers,
		credentials: 'include'
	}).then((response) => response.json());

	depends('admin:users');
	return {
		users: (users as User[]) ?? []
	};
}
