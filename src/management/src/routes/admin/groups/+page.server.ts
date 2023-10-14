import { handleRedirects } from '$lib/auth';
import type { Group } from '../../../types/groups';

export async function load({ fetch, cookies }) {
	const authcheck = await fetch('http://localhost:8080/auth/user', {
		headers: {
			JWT: cookies.get('JWT')
		},
		credentials: 'include'
	}).then((response) => response.json());
    handleRedirects(authcheck)
	const groupResponse = await fetch('http://localhost:8080/group', {
		headers: {
			JWT: cookies.get('JWT')
		},
		credentials: 'include'
	}).then((response) => response.json());

	return {
		groups: groupResponse.groups as Group[]
	};
}
