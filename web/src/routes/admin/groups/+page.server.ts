import { handleRedirects } from '$lib/auth';
import type { Group } from '../../../types/groups';
import type { User } from '../../../types/users';
export async function load({ fetch, cookies, depends }) {
	const authcheck = await fetch('http://localhost:8080/auth/user', {
		headers: {
			JWT: cookies.get('JWT')
		},
		credentials: 'include'
	}).then((response: { json: () => any }) => response.json());
	//	handleRedirects(authcheck);
	const groupResponse = await fetch('http://localhost:8080/group', {
		headers: {
			JWT: cookies.get('JWT')
		},
		credentials: 'include'
	}).then((response: { json: () => any }) => response.json());
	depends('admin:group');

	const usersResponse = await fetch('http://localhost:8080/user', {
		headers: {
			JWT: cookies.get('JWT')
		},
		credentials: 'include'
	}).then((res: { json: () => any }) => res.json());

	return {
		groups: (groupResponse as Group[]) ?? [],
		users: (usersResponse as User[]) ?? []
	};
}
