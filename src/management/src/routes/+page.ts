import { _authStore } from './+layout';

export const ssr = false;

/** @type {import('../../../../.svelte-kit/types/src/routes').PageLoad} */
export async function load() {
	_authStore.checkRoute();
}
