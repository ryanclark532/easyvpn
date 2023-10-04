import { createAuthStore } from '$lib/stores/auth';
import { createUserStore } from '$lib/stores/users';
import { createVpnStore } from '$lib/stores/vpn';

export const prerender = false;
export const ssr = false;

const _authStore = await createAuthStore();
const _userStore = await createUserStore();
const _vpnStore = await createVpnStore();
export { _authStore, _userStore, _vpnStore };
