import { handleRedirects } from '$lib/auth';
import type { ClientSettings } from '../../../../types/settings';

export async function load({ fetch, cookies, depends }) {
	const config = await fetch('http://localhost:8080/settings/file', {
		headers: {
			JWT: cookies.get('JWT')
		},
		credentials: 'include',
		method: 'GET'
	}).then((response) => response.json());
	depends('admin:settings');
	return {
		config: config,
		username: authcheck.name
	};
}
