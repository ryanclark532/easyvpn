import { typedFetch } from '$lib/fetch';
import { writable } from 'svelte/store';
import type {
	ActiveConnectionsResponse,
	Connection,
	serverStatus,
	ServerStatusResponse
} from '../../types/vpn';
import { getToken } from './auth';

export const activeConnectionsFilter = writable<string>();

export async function createVpnStore() {
	const { subscribe, set, update } = writable<serverStatus>('unknown');
	const status = await getStatus();
	if (!(status instanceof Error)) {
		set(status);
	}
	return {
		subscribe,
		set,
		update,
		getStatus: () => getStatus(set),
		operation: (operation: string) => vpnOperation(operation, set)
	};
}

async function getStatus(set?: (this: void, value: serverStatus) => void) {
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
	if (set) {
		set(response.data.status);
	}
	return response.data.status;
}
async function vpnOperation(operation: string, set: (this: void, value: serverStatus) => void) {
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
	set(response.data.status);
	return response.data.status;
}

export async function createConnectionsStore() {
	const { subscribe, set, update } = writable<Connection[]>([]);
	const connections = await getActiveConnections();
	if (!(connections instanceof Error) && connections.length !== 0) {
		set(connections);
	}

	return {
		subscribe,
		set,
		update,
		get: () => getActiveConnections(set)
	};
}

async function getActiveConnections(set?: (this: void, value: Connection[]) => void) {
	const token = getToken();
	if (!token) {
		return new Error('No Token Provided');
	}
	const headers = new Headers();
	headers.append('Authorization', `Bearer ${token}`);
	const response = await typedFetch<ActiveConnectionsResponse>(
		'http://localhost:8080/api/vpn/connections',
		{
			method: 'GET',
			headers
		}
	);

	if (response instanceof Error || response.status >= 400) {
		return new Error('Error Getting Server Status');
	}
	if (set) {
		set(response.data.connections || []);
	}
	return response.data.connections || [];
}
