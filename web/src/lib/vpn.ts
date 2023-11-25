import { invalidate } from '$app/navigation';
import { writable } from 'svelte/store';
import { getToken } from './auth';

export const activeConnectionsFilter = writable<string>();

export async function vpnOperation(operation: string) {
	const headers = new Headers();
	headers.append('JWT', getToken() ?? '');
	const response = await fetch('http://localhost:8080/vpn', {
		headers,
		method: 'POST',
		body: JSON.stringify({ operation }),
		credentials: 'include'
	}).then((response) => response.json());

	if (response.status >= 400) {
		return new Error('Error While issueing vpn operation');
	}
	invalidate('admin:status');
}
