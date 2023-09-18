import { checkToken, getToken } from '$lib/auth';
import { getUsers, users } from '$lib/users';
import { page } from '$app/stores';
import { redirect } from '@sveltejs/kit';

export const ssr = false;

/** @type {import('../../../../.svelte-kit/types/src/routes').PageLoad} */
export async function load() {
	const token = getToken();
	if (!token) {
		throw redirect(307, '/login');
	}
}
