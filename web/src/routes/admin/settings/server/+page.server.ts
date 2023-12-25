import { handleRedirects } from '$lib/auth';
import type { Settings } from '../../../../types/settings';

export async function load({ fetch, cookies, depends }) {
	const settings = await fetch('http://localhost:8080/settings', {
		headers: {
			JWT: cookies.get('JWT')
		},
		credentials: 'include',
		method: 'GET'
	}).then((response) => response.json());
	depends('admin:settings');
	return {
		settings: settings satisfies Settings
	};
}
