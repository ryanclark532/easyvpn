import { _authStore } from '../../+layout';

export const ssr = false;

export async function load() {
	_authStore.checkRoute();
}
