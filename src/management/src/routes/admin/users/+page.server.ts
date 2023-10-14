import { handleRedirects } from '$lib/auth';
import type { User } from '../../../types/users';

export async function load({ fetch, cookies }) {
	const authcheck = await fetch('http://localhost:8080/auth/user', {
		headers: {
			JWT: cookies.get('JWT')
		},
		credentials: 'include'
	}).then((response) => response.json());
    handleRedirects(authcheck)
	
	const usersResponse = await fetch('http://localhost:8080/user', {
		headers: {
			JWT: cookies.get('JWT')
		},
		credentials: 'include'
	}).then((response) => response.json());

	return {
		users: usersResponse.users as User[]
	};
}
