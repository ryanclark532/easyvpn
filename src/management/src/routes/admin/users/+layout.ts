import { _authStore, _userStore } from '../../+layout';

export const ssr = false;

export async function load() {
	_authStore.checkRoute();
}
