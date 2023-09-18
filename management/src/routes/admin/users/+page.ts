import { checkToken, getToken } from '$lib/auth';
import { getUsers, users } from '$lib/users';
import { page } from '$app/stores';
import { redirect } from '@sveltejs/kit';

export const ssr = false;

/** @type {import('../../../../.svelte-kit/types/src/routes').PageLoad} */
export async function load() {
	const token = await checkToken();
	if (!token.passwordValid && token.valid) {
		throw redirect(307, '/user/reset');
	}
	if (!token.admin && token.valid) {
		throw redirect(307, '/user');
	}

	if (!token.valid) {
		throw redirect(307, 'login');
	}

	const u = await getUsers(getToken());
	if (!u.data?.users) {
		return;
	}
	users.set(u.data.users);
}
