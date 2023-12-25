import { handleRedirects } from '$lib/auth';
import type { Settings } from '../../../../types/settings';
import type { ServerStatus } from '../../../../types/vpn';
import os from 'os';
export async function load({ fetch, cookies, depends }) {
	const statusResponse = await fetch('http://localhost:8080/vpn', {
		headers: {
			JWT: cookies.get('JWT')
		},
		credentials: 'include',
		method: 'GET'
	}).then((response) => response.json());
	depends('admin:status');

	const settings = await fetch('http://localhost:8080/settings', {
		headers: {
			JWT: cookies.get('JWT')
		},
		credentials: 'include',
		method: 'GET'
	}).then((response) => response.json());
	depends('admin:settings');

	const connectionsResponse = await fetch('http://localhost:8080/vpn/connections', {
		headers: {
			JWT: cookies.get('JWT')
		},
		credentials: 'include',
		method: 'GET'
	}).then((response) => response.json());
	return {
		hostname: os.hostname(),
		connections: connectionsResponse.connections ? connectionsResponse.connections.length : 0,
		status: statusResponse.status as ServerStatus,
		settings: settings as Settings
	};
}
