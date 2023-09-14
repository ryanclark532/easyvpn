import { checkToken } from '$lib/auth';
import { getUsers } from '$lib/users';

export const ssr = false;

/** @type {import('../../../../.svelte-kit/types/src/routes').PageLoad} */
export async function load() {
	const token = await checkToken();
	const users = await getUsers(token);

	return {
		token,
		users
	};
}
