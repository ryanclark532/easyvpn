import { checkToken } from '$lib/auth';

export const ssr = false;

/** @type {import('../../../../.svelte-kit/types/src/routes').PageLoad} */
export async function load() {
	await checkToken();
}
