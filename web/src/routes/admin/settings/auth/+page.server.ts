import { handleRedirects } from '$lib/auth';
import type { AuthSettings, Settings } from '../../../../types/settings';

export async function load({ fetch, cookies, depends }) {
	const authcheck = await fetch('http://localhost:8080/auth/user', {
		headers: {
			JWT: cookies.get('JWT')
		},
		credentials: 'include'
	}).then((response) => response.json());
	handleRedirects(authcheck);

	const settings = await fetch('http://localhost:8080/settings', {
		headers: {
			JWT: cookies.get('JWT')
		},
		credentials: 'include',
		method: 'GET'
	}).then((response) => response.json());
	depends('admin:settings');
	return settings satisfies Settings;
}
