import { createAuthStore } from '$lib/stores/auth';
import { createUserStore } from '$lib/stores/users';
import { createConnectionsStore, createVpnStore } from '$lib/stores/vpn';

export const prerender = true;

const _authStore = await createAuthStore();
const _userStore = await createUserStore();
const _vpnStore = await createVpnStore();
const _connectionsStore = await createConnectionsStore();
export { _authStore, _userStore, _vpnStore, _connectionsStore };
