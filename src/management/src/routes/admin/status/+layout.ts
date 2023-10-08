import { _authStore, _connectionsStore } from '../../+layout';

/** @type {import('../../../../.svelte-kit/types/src/routes').PageLoad} */
export async function load() {
	_authStore.checkRoute();
}
