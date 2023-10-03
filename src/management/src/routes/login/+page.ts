import { createAuthStore } from '$lib/stores/auth';

export const ssr = false;

/** @type {import('../../../../.svelte-kit/types/src/routes').PageLoad} */
export async function load() {
	return { authStore: createAuthStore() };
}
