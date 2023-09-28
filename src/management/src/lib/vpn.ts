import type { DataWithStatus } from '../types/data-with-status';
import { typedFetch } from '$lib/fetch';
import type { serverStatus, ServerStatusResponse } from '../types/vpn';
import { getToken } from '$lib/auth';
import { writable } from 'svelte/store';

export const operationResponse = writable<DataWithStatus<string>>();
export const status = writable<serverStatus>();

export async function getVpnStatus(
	token: string | undefined
): Promise<DataWithStatus<ServerStatusResponse | undefined>> {
	const headers = new Headers();
	headers.append('Authorization', `Bearer ${token}`);
	const response = await typedFetch<ServerStatusResponse>('http://localhost:8080/api/vpn', {
		method: 'GET',
		headers
	});

	if (response.status >= 400) {
		return {
			status: 'error',
			data: undefined
		};
	}
	return {
		status: 'ready',
		data: response.data
	};
}

export async function vpnOperation(operation: string) {
	const headers = new Headers();
	const token = getToken();
	if (!token) {
		operationResponse.set({
			status: 'error',
			data: `${operation} failed, please try again later`
		});
		status.set('unknown');
		return;
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

	if (response.status >= 400) {
		operationResponse.set({
			status: 'error',
			data: `${operation} server failed, please try again later`
		});
		status.set('unknown');
		return;
	}

	status.set(response.data.status ?? 'unknown');
	return;
}
