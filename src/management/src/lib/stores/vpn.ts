import { typedFetch } from '$lib/fetch';
import { writable } from 'svelte/store';
import type { serverStatus, ServerStatusResponse } from '../../types/vpn';
import { getToken } from './auth';

export async function createVpnStore() {
	const { subscribe, set, update } = writable<serverStatus>();
	const status = await getStatus();
	if (!(status instanceof Error)) {
		set(status);
	}
	return {
		subscribe,
		set,
		update,
		getStatus: () => getStatus(),
		operation: (operation: string) => vpnOperation(operation,set )
	};
}

async function getStatus() {
	const token = getToken();
	if (!token) {
		return new Error('No Token Provided');
	}
	const headers = new Headers();
	headers.append('Authorization', `Bearer ${token}`);
	const response = await typedFetch<ServerStatusResponse>('http://localhost:8080/api/vpn', {
		method: 'GET',
		headers
	});

	if (response instanceof Error || response.status >= 400) {
		return new Error('Error Getting Server Status');
	}

	return response.data.status;
}

async function vpnOperation(operation: string, set:(this: void, value: serverStatus) => void ) {
	const headers = new Headers();
	const token = getToken();
	if (!token) {
		return new Error('No Token Provided');
	}
	headers.append('Authorization', `Bearer ${token}`);
	const response = await typedFetch<ServerStatusResponse>(
		'http://localhost:8080/api/vpn/operation',
		{
			method: 'POST',
			headers,
			body: JSON.stringify({ operation })
		}
	);

	if (response instanceof Error || response.status >= 400) {
		return new Error('Error While issueing vpn operation');
	}
    set(response.data.status)
	return response.data.status;
}
