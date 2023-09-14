import { typedFetch } from '$lib/fetch';
import { getToken } from '$lib/auth';
import type { DataWithStatus } from '../types/data-with-status';
import type { GetUsersResponse } from '../types/users';

export async function getUsers(
	token: string
): Promise<DataWithStatus<GetUsersResponse | undefined>> {
	const headers = new Headers();
	headers.append('Authorization', `Bearer ${token}`);
	const response = await typedFetch<GetUsersResponse>('http://localhost:8080/users', {
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
