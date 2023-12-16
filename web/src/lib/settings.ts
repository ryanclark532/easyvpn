import { invalidate } from '$app/navigation';
import type { Settings } from '../types/settings';
import { getToken } from './auth';

export async function setSettings(settings: Settings) {
	const headers = new Headers();
	headers.append('JWT', getToken() ?? '');
	const response = await fetch(`http://localhost:8080/settings`, {
		method: 'PUT',
		headers,
		credentials: 'include',
		body: JSON.stringify(settings)
	});

	if (response.status >= 400) {
		return new Error('Error updating users, please try again later');
	}
	invalidate('admin:setttings');
}
