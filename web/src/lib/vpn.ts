import { invalidate } from '$app/navigation';
import { writable } from 'svelte/store';

export const activeConnectionsFilter = writable<string>();

export async function vpnOperation(operation: string) {
	console.log(document.cookie.split(';'));
	const response = await fetch('http://localhost:8080/vpn/operation', {
		method: 'POST',
		body: JSON.stringify({ operation }),
		credentials: 'include'
	}).then((response) => response.json());

	if (response.status >= 400) {
		return new Error('Error While issueing vpn operation');
	}
	invalidate('http://localhost:8080/vpn');
}
