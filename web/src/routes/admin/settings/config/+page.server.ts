import { handleRedirects } from '$lib/auth';
import type { ClientSettings } from '../../../../types/settings';

export async function load({ fetch, cookies, depends }) {
	const authcheck = await fetch('http://localhost:8080/auth/user', {
		headers: {
			JWT: cookies.get('JWT')
		},
		credentials: 'include'
	}).then((response) => response.json());
	handleRedirects(authcheck);

	const config = await fetch('http://localhost:8080/settings/file', {
		headers: {
			JWT: cookies.get('JWT')
		},
		credentials: 'include',
		method: 'GET'
	}).then((response) => response.json());
	depends('admin:settings');
	return {
		config: config
	};
}
