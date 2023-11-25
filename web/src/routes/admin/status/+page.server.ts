import { handleRedirects } from '$lib/auth';
import type { Connection, ServerStatus } from '../../../types/vpn';

export async function load({ fetch, cookies, depends }) {
	const authcheck = await fetch('http://localhost:8080/auth/user', {
		headers: {
			JWT: cookies.get('JWT')
		},
		credentials: 'include'
	}).then((response) => response.json());
	handleRedirects(authcheck);

	const statusResponse = await fetch('http://localhost:8080/vpn', {
		headers: {
			JWT: cookies.get('JWT')
		},
		credentials: 'include',
		method: 'GET'
	}).then((response) => response.json());
	console.log(statusResponse);
	depends('admin:status');

	const connectionsResponse = await fetch('http://localhost:8080/vpn/connections', {
		headers: {
			JWT: cookies.get('JWT')
		},
		credentials: 'include',
		method: 'GET'
	}).then((response) => response.json());

	return {
		status: statusResponse.status as ServerStatus,
		connections: (connectionsResponse.connections as Connection[]) || []
	};
}
