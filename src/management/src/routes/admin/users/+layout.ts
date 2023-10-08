import { _authStore, _userStore } from '../../+layout';

export async function load() {
	_authStore.checkRoute();
}
