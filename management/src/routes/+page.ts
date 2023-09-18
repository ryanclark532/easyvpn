import { checkToken } from '$lib/auth';
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
}
